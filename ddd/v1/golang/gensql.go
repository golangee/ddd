package golang

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"path/filepath"
	"strings"
)

func createSQLLayer(ctx *genctx, rslv *resolver, bc *ddd.BoundedContextSpec, sql ddd.SQLLayer) error {
	bcPath := filepath.Join("internal", text.Safename(bc.Name()))
	layerPath := filepath.Join(bcPath, sql.Name())

	file := ctx.newFile(layerPath, "sql", "").
		SetPackageDoc(sql.Description())

	for _, repo := range sql.Repositories() {
		iface := bc.SPIServiceByName(repo.InterfaceName())
		if iface == nil {
			panic("illegal state: validate the model first")
		}

		repoSpec := ctx.repoSpecByName(repo.InterfaceName())
		if repoSpec == nil {
			panic("illegal state: validate the model first and define core layer before sql layer")
		}

		impl := src.Implement(repoSpec.iface, true)
		impl.SetName(sql.Name() + repo.InterfaceName())
		impl.SetDoc("...is an implementation of the " + pkgNameCore + "." + repo.InterfaceName() + " defined as SPI/driven port in the domain/core layer.\nThe queries are specific for the " + strings.ToLower(sql.Name()) + " dialect.")
		file.AddTypes(impl)

		for _, method := range impl.Methods() {
			sqlSpec := repo.ImplementationByName(method.Name())
			if sqlSpec == nil {
				panic("illegal state: undefined method: validate the model first")
			}
			body := src.NewBlock()
			method.AddBody(body)
			body.AddLine("const s = \"", string(sqlSpec.RawStatement())+"\"")
		}
	}

	return nil
}
