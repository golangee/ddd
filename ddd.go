package ddd

type TypeName string

const (
	Int64  TypeName = "int64"
	String TypeName = "string"
	UUID   TypeName = "uuid"
	Error  TypeName = "error"
	Reader TypeName = "io.Reader"
	Writer TypeName = "io.Writer"
)

func Int32(name string, comment ...string) *ParamSpec {

}

func List(t TypeName) TypeName {
	return "[]" + t
}


type ApplicationBuilder struct {
}

func (a *ApplicationBuilder) Generate() error {

}

type MethodSpecification struct {
}

type DomainsSpecification struct {
}

func BoundedContexts(domains ...*DomainSpec) *DomainsSpecification {

}


func Func(name string, comment string, inSpec *ParamSpecs, outSpec *ParamSpecs) *MethodSpecification {

}

type ParamSpec struct{}

type ParamSpecs struct{}

func In(params ...*ParamSpec) *ParamSpecs {

}

func Out(params ...*ParamSpec) *ParamSpecs {

}

func Var(name, typ TypeName, comment ...string) *ParamSpec {

}

func Return(typ TypeName) *ParamSpec {

}


type InterfaceSpecs struct{}

func Repositories(repos ...*InterfaceSpec) *InterfaceSpecs {

}

func Requires(repos ...*InterfaceSpec) *InterfaceSpecs {

}
func Persistence(repos *InterfaceSpecs, types *TypeSpecs, impls *ImplementationSpecs) *PersistenceSpec {

}

type ImplementationSpec struct {
}

type ImplementationSpecs struct {
}

func Implementations(impls ...*ImplementationSpec) *ImplementationSpecs {

}

func MySQL(name TypeName, migrations *MigrationSpecs, methods ...*MethodImplSpec) *ImplementationSpec {

}

func Filesystem(name TypeName) *ImplementationSpec {

}

type Body struct{}

func Statement(query string) *Body {

}

func DefaultCreate(table string) *Body {}

func DefaultDelete(table string) *Body {}

type MethodImplSpec struct{}

func Implement(method string, body *Body) *MethodImplSpec {

}

func SQL(migrations *MigrationSpecs) *ImplementationSpec {}

type MigrationSpec struct{}

func Schema(migrations ...*MigrationSpec) *MigrationSpecs {

}

func Bla(str string) {

}

// in yyyyMMddHHmmss format
func Migrate(dateTime uint64, sql string) *MigrationSpec {

}

type MigrationSpecs struct{}

type TypeSpecs struct{}

type TypeSpecification struct{}

type FieldSpecs struct {
}

type FieldSpec struct {
}

func Fields(fields ...*FieldSpec) *FieldSpecs {

}

func Field(name string, comment string, typ TypeName) *FieldSpec {

}

func Type(name string, fields ...*FieldSpec) *TypeSpecification {

}

func Types(types ...*TypeSpecification) *TypeSpecs {

}

type PersistenceSpec struct {
}

func ApplicationDomain(name string, domains *DomainsSpecification) *ApplicationBuilder {

}
