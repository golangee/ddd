package ddd

type TypeName string

const (
	Int64  TypeName = "int64"
	String TypeName = "string"
	UUID TypeName = "uuid"
	Error TypeName = "error"
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

func SQL()*ImplementationSpec{}

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
