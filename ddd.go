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

type MimeType string

const (
	MimeTypeJson MimeType = "application/json"
	MimeTypeXml  MimeType = "application/xml"
	MimeTypeText MimeType = "application/text"
)

type StructOrInterface interface {
	Name() string
	structOrInterface()
}

type FuncOrStruct interface {
	Name() string
	funcOrStruct()
}

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
	name string
}

func (m *MethodSpecification) funcOrStruct() {
	panic("implement me")
}

func (m *MethodSpecification) Name() string {
	return m.name
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

func Contracts(repos ...*InterfaceSpec) *InterfaceSpecs {

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


func Filesystem(name TypeName) *ImplementationSpec {

}

type Body struct{}


func DefaultCreate(table string) *Body {}

func DefaultDelete(table string) *Body {}

type MethodImplSpec struct{}

func Implement(method string, body *Body) *MethodImplSpec {

}



func API(specs ...StructOrInterface) []StructOrInterface {
	return specs
}

type MethodSpecs struct {
	specs []*MethodSpecification
}

func Factories(specs ...FuncOrStruct) []FuncOrStruct {
	return specs
}


type TypeSpecs struct{}

type TypeSpecification struct {
	name string
}

func (s *TypeSpecification) funcOrStruct() {
	panic("implement me")
}

func (s *TypeSpecification) structOrInterface() {
	panic("implement me")
}

func (s *TypeSpecification) Name() string {
	return s.name
}

type FieldSpecs struct {
}

type FieldSpec struct {
}

func Fields(fields ...*FieldSpec) *FieldSpecs {

}

func Field(name string, typ TypeName, comment ...string) *FieldSpec {

}

func Type(name string, fields ...*FieldSpec) *TypeSpecification {

}

func Struct(name string, fields ...*FieldSpec) *TypeSpecification {

}

// CopyStruct really issues a copy to ensure that layers are not polluted. E.g. a REST model may be the same
// as the database model or domain model, but actually they must not have much in common and may contain
// undesired internal things, which must not be presented to the outside.
func CopyStruct(from string, name TypeName) *TypeSpecification {

}

func Types(types ...*TypeSpecification) *TypeSpecs {

}

func DataTypes(types ...*TypeSpecification) *TypeSpecs {

}

type PersistenceSpec struct {
}

func ApplicationDomain(name string, domains *DomainsSpecification) *ApplicationBuilder {

}
