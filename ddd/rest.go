package ddd

import (
	"golang.org/x/mod/semver"
)

// A RestLayerSpec represents a stereotyped PRESENTATION layer.
type RestLayerSpec struct {
	name        string
	version     string
	description string
	resources   []*HttpResourceSpec
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

// REST is a driven port and directly uses the UseCases API. It does normally not define any further models
// than the use cases it imports. Because of transitivity, it has also access to the CORE layer API models.
// Version must be a semantic version string like "v1.2.3".
func REST(version string, resources []*HttpResourceSpec) *RestLayerSpec {
	major := semver.Major(version)
	name := "rest" + major
	return &RestLayerSpec{
		name:    name,
		version: version,
		description: "Package " + name + " contains the REST specific implementation for the current bounded context.\n" +
			"It depends only from the usecases and transitivley on the core API.",
		resources: resources,
	}
}

type HttpResourceSpec struct {
	comment string
	method  string
	path    string
}

type HttpResourceSpecs struct {
	resources []*HttpResourceSpec
}

func Resources(resources ...*HttpResourceSpec) []*HttpResourceSpec {
	return resources
}

func GET(comment string, in []*ParamSpec, responses []*HttpResponseSpec) *VerbSpec {
	return nil
}

func DELETE(comment string) *VerbSpec {
	return nil
}

func POST(comment string) *VerbSpec {
	return nil
}

func PUT(comment string) *VerbSpec {
	return nil
}

func Resource(path, comment string, verbs ...*VerbSpec) *HttpResourceSpec {
	return &HttpResourceSpec{path: path, comment: comment}
}

func Responses(responses ...*HttpResponseSpec) []*HttpResponseSpec {
	return responses
}

func Response(status int, comment string, headers []*HeaderSpec, mimes []*ResponseFormatSpec) *HttpResponseSpec {
	return &HttpResponseSpec{
		comment:    comment,
		statusCode: status,
		mimeTypes:  mimes,
	}
}

type ResponseFormatSpec struct {
	typeName TypeName
	mimeType MimeType
}

func JSON(typeName TypeName) *ResponseFormatSpec {
	return ForMimeType(MimeTypeJson, typeName)
}

func Text(typeName TypeName) *ResponseFormatSpec {
	return ForMimeType(MimeTypeText, typeName)
}

func XML(typeName TypeName) *ResponseFormatSpec {
	return ForMimeType(MimeTypeXml, typeName)
}

func BinaryStream(typeName TypeName) *ResponseFormatSpec {
	return ForMimeType("application/octetstream", typeName)
}

func JPEG(typeName TypeName) *ResponseFormatSpec {
	return ForMimeType("image/jpeg", typeName)
}

func ForMimeType(mimeType MimeType, name TypeName) *ResponseFormatSpec {
	return &ResponseFormatSpec{
		typeName: name,
		mimeType: mimeType,
	}
}

func ContentTypes(mimes ...*ResponseFormatSpec) []*ResponseFormatSpec {
	return mimes
}

type VerbSpec struct {
}

type HttpResponseSpec struct {
	comment    string
	statusCode int
	mimeTypes  []*ResponseFormatSpec
}

type HeaderSpec struct {
	key      string
	typeName TypeName
	comment  []string
}

func Header(key string, typeName TypeName, comment ...string) *HeaderSpec {
	return &HeaderSpec{
		key:      key,
		typeName: typeName,
		comment:  comment,
	}
}

func Headers(headers ...*HeaderSpec) []*HeaderSpec {
	return headers
}
