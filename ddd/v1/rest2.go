package ddd

import "golang.org/x/mod/semver"

// A RestLayerSpec represents a stereotyped PRESENTATION Layer.
type RestLayerSpec struct {
	name        string
	description string
	versions    []*VersionSpec
	pos         Pos
}

// REST is a driven port and directly uses the UseCases API. It does normally not define any further models
// than the use cases it imports. Because of transitivity, it has also access to the CORE layer API models.
//
// Version must be a semantic version string like "v1.2.3".
func REST(versions ...*VersionSpec) *RestLayerSpec {
	return &RestLayerSpec{
		name:     "Rest",
		versions: versions,
		description: "Package rest contains the REST specific implementation for the current bounded context.\n" +
			"It depends only from the use cases and transitively on the core API.",
		pos: capturePos("REST", 1),
	}
}

// Versions returns the supported api versions, only one per major version.
func (r *RestLayerSpec) Versions() []*VersionSpec {
	return r.versions
}

// Walk inspects all children
func (r *RestLayerSpec) Walk(f func(obj interface{}) error) error {
	if err := f(r); err != nil {
		return err
	}

	for _, obj := range r.versions {
		if err := f(obj); err != nil {
			return err
		}
	}

	return nil
}

// Pos returns the debug position.
func (r *RestLayerSpec) Pos() Pos {
	return r.pos
}

// Name of the Layer
func (r *RestLayerSpec) Name() string {
	return r.name
}

// Description of the layer
func (r *RestLayerSpec) Description() string {
	return r.description
}

// Stereotype of the layer
func (r *RestLayerSpec) Stereotype() Stereotype {
	return PRESENTATION
}

type VersionSpec struct {
	name        string
	version     string
	middlewares []*MiddlewareSpec
	resources   []*HttpResourceSpec
}

func Version(version, description string, middlewares []*MiddlewareSpec, resources ...*HttpResourceSpec) *VersionSpec {
	return &VersionSpec{
		name:    "Rest" + semver.Major(version),
		version: version,
	}
}

// Walk inspects all children
func (v *VersionSpec) Walk(f func(obj interface{}) error) error {
	if err := f(v); err != nil {
		return err
	}

	for _, obj := range v.middlewares {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	for _, obj := range v.resources {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	return nil
}

// MajorVersion returns the major version like v1 or v2.
func (v *VersionSpec) MajorVersion() string {
	return semver.Major(v.version)
}

// Version returns the actual semantic version like v1.2.3.
func (v *VersionSpec) Version() string {
	return v.version
}

// HttpResourceSpec represents a REST resource (e.g. /book/:id)
type HttpResourceSpec struct {
	path        string
	description string
	params      []QueryParamsOrHeaderParamsOrPathParams
	pos         Pos
}

// Walk inspects all children.
func (r *HttpResourceSpec) Walk(f func(obj interface{}) error) error {
	if err := f(r); err != nil {
		return err
	}

	for _, obj := range r.params {
		for _, p := range obj.Params() {
			if err := f(p); err != nil {
				return err
			}
		}
	}

	return nil
}

// Pos returns the debug position.
func (r *HttpResourceSpec) Pos() Pos {
	return r.pos
}

// Path returns the resource path.
func (r *HttpResourceSpec) Path() string {
	return r.path
}

// Resource2 is a factory for a HttpResourceSpec. It defines a path, containing potential path parameters, and multiple
// verb specifiers. Parameters can be defined globally.
func Resource(path string, spec *VerbAndContentSpec, specs ...*VerbAndContentSpec) *HttpResourceSpec {
	return &HttpResourceSpec{
		path: path,
		pos:  capturePos("Resource", 1),
	}
}

type MiddlewareSpec struct {
	name    string
	comment string
	params  []QueryParamsOrHeaderParamsOrPathParams
	out     ProvidedParamSpecs
}

// A Middleware defines a func specification to be used as an http.HandlerFunc. It returns a unique struct,
// containing the parameters and a potential error.
func Middleware(name, comment string, params []QueryParamsOrHeaderParamsOrPathParams, out ProvidedParamSpecs) *MiddlewareSpec {
	return &MiddlewareSpec{
		name:    name,
		comment: comment,
		params:  params,
		out:     out,
	}
}

// Walk inspects all children
func (v *MiddlewareSpec) Walk(f func(obj interface{}) error) error {
	if err := f(v); err != nil {
		return err
	}

	for _, obj := range v.params {
		for _, p := range obj.Params() {
			if err := f(p); err != nil {
				return err
			}
		}
	}

	return nil
}

func Middlewares(middlewares ...*MiddlewareSpec) []*MiddlewareSpec {
	return middlewares
}

type VerbAndContentSpec struct {
}

type ParamParser struct {
}

func GET() *VerbAndContentSpec {
	return nil
}

func POST(after []string, call *CallSpec, resp *ResponseSpec) *VerbAndContentSpec {
	return nil
}

type ResponseSpec struct {
}

func Response(parsers ...*ParamParser) *ResponseSpec {
	return nil
}

func HeaderVar(name, funcParamName string) *ParamParser {
	return nil
}

func BodyVar(contentType MimeType, funcParamName string) *ParamParser {
	return nil
}

func After(middlewareNames ...string) []string {
	return nil
}

type CallSpec struct {
}

func Invoke(epicName, funcName string, parsers ...*ParamParser) *CallSpec {
	return nil
}
