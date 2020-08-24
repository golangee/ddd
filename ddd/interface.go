package ddd

type CRUDOptions int

const (
	CreateOne CRUDOptions = 1 << iota
	ReadOne
	ReadAll
	UpdateOne
	DeleteOne
	DeleteAll
	CountAll
)

const CRUD = ReadOne | CreateOne | DeleteOne | UpdateOne | DeleteAll | CountAll | ReadAll

type InterfaceSpec struct {
	name  string
	funcs []*MethodSpecification
}

func (s *InterfaceSpec) structOrInterface() {
	panic("implement me")
}

func (s *InterfaceSpec) Name() string {
	return s.name
}

func Interface(name string, methods ...*MethodSpecification) StructOrInterface {
	return &InterfaceSpec{name: name, funcs: methods}
}

// CRUDRepository declares a repository which will be automatically implemented based on the given
// table name and entity type name (a struct defined in the current layer). The opts indicate which
// kind of methods will be created. If the name collides with another generated type, the implementation is
// mixed into the already defined type.
//
// The generated signatures are as follows
//  * ReadOne: ReadOne<TypeName> (ctx context.Context, <primaryKey> <primaryKey-Type>) (<TypeName>,error)
//    -> ReadOneBook (ctx context.Context, id int64) (Book, error)
//  * ReadAll: ReadAll<T>s (ctx context.Context) (<T>,error)
//    -> ReadAllBooks (ctx context.Context) ([]Book, error)
func CRUDRepository(name, comment, table string, entity TypeName, opts CRUDOptions) *GenSpec {
	return &GenSpec{}
}
