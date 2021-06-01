package adl

import (
	"github.com/golangee/architecture/arc/token"
	"runtime"
)

// A Project has multiple modules, like libraries, servers or clients, frontends or backends.
type Project struct {
	Name    token.String
	Modules []*Module
}

func NewProject(name string) *Project {
	return &Project{
		Name: traceStr(name),
	}
}

func (w *Project) AddProjects(p ...*Module) *Project {
	w.Modules = append(w.Modules, p...)
	return w
}

// A Module is e.g. a server application, a frontend or a shared library.
type Module struct {
	Name      token.String // the name of the project
	Generator *Generator
	Domain    *Domain
}

func NewModule(name string) *Module {
	return &Module{
		Name: traceStr(name),
	}
}

func (p *Module) SetGenerator(g *Generator) *Module {
	p.Generator = g
	return p
}

func (p *Module) SetDomain(d *Domain) *Module {
	p.Domain = d
	return p
}

// A Generator describes how this project should be generated.
type Generator struct {
	Go     *Golang
	OutDir token.String // the target directory to (re) write the module
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) SetGo(d *Golang) *Generator {
	g.Go = d
	return g
}

func (g *Generator) SetOutDir(dir string) *Generator {
	g.OutDir = traceStr(dir)
	return g
}

// Golang describes how a Go project (or module) must be created or updated.
type Golang struct {
	Module token.String // the name of the go module, e.g. github.com/worldiety/supportiety
}

func NewGolang() *Golang {
	return &Golang{}
}

func (g *Golang) SetModName(name string) *Golang {
	g.Module = traceStr(name)
	return g
}

// A BoundedContext is a cross-cutting thing, which is referenced from various places and contains its own
// ubiquitous language (== glossary).
type BoundedContext struct {
	Name token.String
}

type Domain struct {
	Name token.String
	Core        *Package
	Usecase     *Package
}

func NewDomain(name string) *Domain {
	return &Domain{
		Name: traceStr(name),
	}
}

func (d *Domain) SetCore(l *Package) *Domain {
	d.Core = l
	return d
}

func (d *Domain) SetUsecase(l *Package) *Domain {
	d.Usecase = l
	return d
}

type Package struct {
	Repositories []*Interface
}

func NewPackage() *Package {
	return &Package{}
}

func (p *Package) AddRepositories(r ...*Interface) *Package {
	p.Repositories = append(p.Repositories, r...)
	return p
}

type Interface struct {
	Comment token.String
	Name    token.String
	Methods []*Method
}

func NewInterface(name, comment string) *Interface {
	return &Interface{
		Comment: traceStr(comment),
		Name:    traceStr(name),
	}
}

func (i *Interface) AddMethods(m ...*Method) *Interface {
	i.Methods = append(i.Methods, m...)
	return i
}

type DTO struct {
	Fields []Field
}

type Field struct {
	Name token.String
}

type Method struct {
	Comment token.String
	Name    token.String
	In      []*Param
	Out     []*Param
}

func NewMethod(name, comment string) *Method {
	return &Method{
		Comment: traceStr(comment),
		Name:    traceStr(name),
	}
}

func (m *Method) AddIn(name, comment, typeName string) *Method {
	m.In = append(m.In, &Param{
		Name:     traceStr(name),
		Comment:  traceStr(comment),
		TypeName: traceStr(typeName),
	})

	return m
}

func (m *Method) AddOut(name, comment, typeName string) *Method {
	m.Out = append(m.Out, &Param{
		Name:     traceStr(name),
		Comment:  traceStr(comment),
		TypeName: traceStr(typeName),
	})

	return m
}

func (m *Method) OutParams(p ...*Param) *Method {
	m.In = p
	return m
}

type Param struct {
	Comment  token.String
	Name     token.String
	TypeName token.String
}

func traceStr(v string) token.String {
	_, file, line, _ := runtime.Caller(2)
	str := token.NewString(v)
	str.BeginPos.File = file
	str.BeginPos.Line = line
	str.BeginPos.Col = 1
	str.BeginPos.Offset = -1

	str.EndPos = str.BeginPos
	str.EndPos.Col += len(v)

	return str
}
