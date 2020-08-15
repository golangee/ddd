package ddd

type TypeName string

const (
	Int64  TypeName = "int64"
	String TypeName = "string"
	UUID TypeName = "uuid"
	Error TypeName = "error"
	Reader TypeName = "io.Reader"
	Writer TypeName = "io.Writer"
)

func List(t TypeName)TypeName{
	return "[]"+t
}

type RepositorySpecification struct {
}

type ApplicationBuilder struct {
}

func (a *ApplicationBuilder) Generate() error {

}

type MethodSpecification struct {
}

type DomainSpecification struct {
}

type DomainsSpecification struct {
}

func Domains(domains ...*DomainSpecification) *DomainsSpecification {

}

func Domain(name string, comment string,builder *PersistenceSpecification) *DomainSpecification {

}

func Method(name string, comment string,inSpec *ParamSpecs,outSpec *ParamSpecs) *MethodSpecification {

}

type ParamSpec struct{}

type ParamSpecs struct{}

func In(params...*ParamSpec)*ParamSpecs{

}

func Out(params...*ParamSpec)*ParamSpecs{

}

func Param(name,comment string, typ TypeName)*ParamSpec{

}

func Return(typ TypeName)*ParamSpec{

}



func Interface(name string, methods ...*MethodSpecification) *RepositorySpecification {

}

type RepositoriesSpecification struct{}

func Repositories(repos ...*RepositorySpecification) *RepositoriesSpecification {

}
func Persistence(repos *RepositoriesSpecification, types *TypesSpecification,impls*ImplementationSpecs) *PersistenceSpecification {

}

type ImplementationSpec struct{

}

type ImplementationSpecs struct{

}

func Implementations(impls...*ImplementationSpec)*ImplementationSpecs{

}

func MySQL(name TypeName,migrations *MigrationSpecs,methods...*MethodImplSpec)*ImplementationSpec{

}

func Filesystem(name TypeName)*ImplementationSpec{

}


type Body struct{}

func Statement(query string)*Body{

}

func DefaultCreate(table string)*Body{}

func DefaultDelete(table string)*Body{}

type MethodImplSpec struct {}

func Implement(method string, body *Body)*MethodImplSpec{

}

func SQL(migrations *MigrationSpecs)*ImplementationSpec{}

type MigrationSpec struct{}

func Schema(migrations...*MigrationSpec)*MigrationSpecs{

}

// in yyyyMMddHHmmss format
func Migrate(dateTime uint64,sql string)*MigrationSpec{

}

type MigrationSpecs struct{}

type TypesSpecification struct{}

type TypeSpecification struct{}

type FieldSpecs struct {
}

type FieldSpec struct {
}

func Fields(fields ...*FieldSpec) *FieldSpecs {

}

func Field(name string, comment string, typ TypeName) *FieldSpec {

}

func Type(name string, fields *FieldSpecs) *TypeSpecification {

}

func Types(types ...*TypeSpecification) *TypesSpecification {

}

type PersistenceSpecification struct {
}

func Application(name string, domains *DomainsSpecification) *ApplicationBuilder {

}
