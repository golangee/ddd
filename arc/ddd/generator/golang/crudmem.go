package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
)

func renderCrud(file *ast.File, iface *ast.Interface, crud *adl.CRUD) (*ast.Struct, error) {
	entityType := astutil.MakeTypeDecl(crud.EntityType)
	resolvedEntityType := astutil.Resolve(file, entityType.String())
	if resolvedEntityType == nil {
		return nil, fmt.Errorf("crud entity type cannot be resolved: " + entityType.String())
	}

	repo := ast.NewStruct("InMemory" + iface.TypeName)
	repo.SetDefaultRecName("r")
	file.AddTypes(repo)

	switch crud.Persistence {
	case adl.PMemory:
		return repo, renderCrudMem(file, iface, crud, entityType, repo)
	}

	panic("unknown persistence " + crud.Persistence)
}

func renderCrudMem(file *ast.File, iface *ast.Interface, crud *adl.CRUD, entityType ast.TypeDecl, repo *ast.Struct) error {
	repo.SetComment("...implements a hashmap based in-memory implementation for " + ast.Name(entityType.String()).Identifier() + " entities.")
	var keyType ast.TypeDecl
	if crud.IDType != nil {
		keyType = astutil.MakeTypeDecl(crud.IDType)
	} else {
		resolvedEntityType := astutil.Resolve(file, entityType.String())

		id := astutil.FieldByName(resolvedEntityType, "ID")
		if id == nil {
			return fmt.Errorf("crud entity type has neither custom ID type nor an ID field")
		}

		keyType = id.TypeDecl().Clone()
	}

	repo.AddFields(
		ast.NewField("store", ast.NewMapDecl(keyType, entityType)).SetVisibility(ast.Private),
		ast.NewField("mutex", ast.NewSimpleTypeDecl("sync.RWMutex")).SetVisibility(ast.Private),
	)

	if crud.InsertOne {
		repo.AddMethods(
			ast.NewFunc("InsertOne").
				SetComment("...inserts the entity or fails if already exists.").
				AddParams(ast.NewParam("entity", entityType.Clone())).
				AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
				SetRecName(repo.DefaultRecName).
				SetBody(
					ast.NewBlock(ast.NewTpl(`r.mutex.Lock()
						defer r.mutex.Unlock()
					
						if _, ok := r.store[entity.ID]; ok {
							return {{.Use "io/fs.ErrExist"}}
						}
					
						r.store[entity.ID] = entity
					
						return nil
					`)),
				),
		)
	}

	if crud.UpdateOne {
		repo.AddMethods(
			ast.NewFunc("UpdateOne").
				SetComment("...updates the entity or fails if does not exist.").
				AddParams(ast.NewParam("entity", entityType.Clone())).
				AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
				SetRecName(repo.DefaultRecName).
				SetBody(
					ast.NewBlock(ast.NewTpl(`r.mutex.Lock()
						defer r.mutex.Unlock()
					
						if _, ok := r.store[entity.ID]; !ok {
							return {{.Use "io/fs.ErrNotExist"}}
						}
					
						r.store[entity.ID] = entity
					
						return nil
					`)),
				),
		)
	}

	if crud.FindOne {
		repo.AddMethods(
			ast.NewFunc("FindOne").
				SetComment("...finds the entity or fails if does not exist.").
				AddParams(ast.NewParam("id", keyType.Clone())).
				AddResults(ast.NewParam("", entityType.Clone())).
				AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
				SetRecName(repo.DefaultRecName).
				SetBody(
					ast.NewBlock(ast.NewTpl(`r.mutex.RLock()
						defer r.mutex.RUnlock()
					
						v, ok := r.store[id];
						if !ok {
							return v, {{.Use "io/fs.ErrNotExist"}}
						}
					
						return v, nil
					`)),
				),
		)
	}

	if crud.DeleteOne {
		repo.AddMethods(
			ast.NewFunc("DeleteOne").
				SetComment("...deletes the entity with the given id. Is a no-op if no such entity exists.").
				AddParams(ast.NewParam("id", keyType.Clone())).
				AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
				SetRecName(repo.DefaultRecName).
				SetBody(
					ast.NewBlock(ast.NewTpl(`r.mutex.Lock()
						defer r.mutex.Unlock()

						delete(r.store,id)
					
						return nil
					`)),
				),
		)
	}

	return nil
}
