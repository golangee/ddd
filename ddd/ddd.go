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

func List(t TypeName) TypeName {
	return "[]" + t
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




func Func(name string, comment string, inSpec []*ParamSpec, outSpec []*ParamSpec) *MethodSpecification {
	panic("implement me")
}

type ParamSpec struct{}

func In(params ...*ParamSpec) []*ParamSpec {
	return params
}

func Out(params ...*ParamSpec) []*ParamSpec {
	return params
}

func Var(name, typ TypeName, comment ...string) *ParamSpec {
	panic("implement me")
}

func Return(typ TypeName) *ParamSpec {
	panic("implement me")
}

type InterfaceSpecs struct{}

type ImplementationSpec struct {
}

type ImplementationSpecs struct {
}

type Body struct{}

type MethodImplSpec struct{}

func API(specs ...StructOrInterface) []StructOrInterface {
	return specs
}

type MethodSpecs struct {
	specs []*MethodSpecification
}

func Factories(specs ...FuncOrStruct) []FuncOrStruct {
	return specs
}

type StructSpec struct {
	name string
}

func (s *StructSpec) funcOrStruct() {
	panic("implement me")
}

func (s *StructSpec) structOrInterface() {
	panic("implement me")
}

func (s *StructSpec) Name() string {
	return s.name
}

type FieldSpec struct {
}

func Fields(fields ...*FieldSpec) []*FieldSpec {
	return fields
}

// TODO a good idea?
func RemoveField(name string) *FieldSpec {
	return nil
}

func Field(name string, typ TypeName, comment ...string) *FieldSpec {
	panic("implement me")
}

func Struct(name, comment string, fields ...*FieldSpec) *StructSpec {
	panic("implement me")
}

// CopyStruct really issues a copy to ensure that layers are not polluted. E.g. a REST model may be the same
// as the database model or domain model, but actually they must not have much in common and may contain
// undesired internal things, which must not be presented to the outside.
func CopyStruct(layerName string, structName TypeName, newFields []*FieldSpec) *StructSpec {
	panic("implement me")
}

func Types(types ...*StructSpec) []*StructSpec {
	return types
}

type PersistenceSpec struct {
}

