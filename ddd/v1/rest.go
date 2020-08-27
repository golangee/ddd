// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ddd

import (
	"golang.org/x/mod/semver"
	"reflect"
)

// A RestLayerSpec represents a stereotyped PRESENTATION Layer.
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
//
// Version must be a semantic version string like "v1.2.3".
func REST(version string, resources []*HttpResourceSpec) *RestLayerSpec {
	major := semver.Major(version)
	name := "rest" + major
	return &RestLayerSpec{
		name:    name,
		version: version,
		description: "Package " + name + " contains the REST specific implementation for the current bounded context.\n" +
			"It depends only from the use cases and transitively on the core API.",
		resources: resources,
	}
}

// HttpResourceSpec represents a REST resource (e.g. /book/:id)
type HttpResourceSpec struct {
	path        string
	description string
	params      []QueryParamsOrHeaderParamsOrPathParams
	verbs       []*VerbSpec
}

// Resources is a factory for a slice of HttpResourceSpec
func Resources(resources ...*HttpResourceSpec) []*HttpResourceSpec {
	return resources
}

// VerbSpec models a HTTP Verb like GET, POST, PATCH, PUT, DELETE.
type VerbSpec struct {
	method       string
	description  string
	pathParams   PathParams
	headerParams HeaderParams
	queryParams  QueryParams
	body         []*ContentSpec
	responses    []*HttpResponseSpec
}

// newVerbSpec creates a new verb specifiction.
func newVerbSpec(method, description string, params []QueryParamsOrHeaderParamsOrPathParams, body []*ContentSpec, responses []*HttpResponseSpec) *VerbSpec {
	verb := &VerbSpec{
		method:       method,
		description:  description,
		pathParams:   nil,
		headerParams: nil,
		queryParams:  nil,
		body:         body,
		responses:    responses,
	}
	for _, param := range params {
		switch t := param.(type) {
		case PathParams:
			verb.pathParams = t
		case HeaderParams:
			verb.headerParams = t
		case QueryParams:
			verb.queryParams = t
		default:
			panic("not implemented: " + reflect.TypeOf(t).String())
		}
	}
	return &VerbSpec{}
}

// GET declares the according verb and is not allowed to have a request body, as defined per RFC 7231.
func GET(description string, params []QueryParamsOrHeaderParamsOrPathParams, responses []*HttpResponseSpec) *VerbSpec {
	return newVerbSpec("GET", description, params, nil, responses)
}

// POST declares the according verb.
func POST(description string, params []QueryParamsOrHeaderParamsOrPathParams, body []*ContentSpec, responses []*HttpResponseSpec) *VerbSpec {
	return newVerbSpec("POST", description, params, body, responses)
}

// DELETE is not allowed to have a request body, as defined per RFC 7231.
func DELETE(description string, params []QueryParamsOrHeaderParamsOrPathParams, responses []*HttpResponseSpec) *VerbSpec {
	return newVerbSpec("DELETE", description, params, nil, responses)
}

// POST declares the according verb.
func PUT(description string, params []QueryParamsOrHeaderParamsOrPathParams, body []*ContentSpec, responses []*HttpResponseSpec) *VerbSpec {
	return newVerbSpec("PUT", description, params, body, responses)
}

// POST declares the according verb.
func PATCH(description string, params []QueryParamsOrHeaderParamsOrPathParams, body []*ContentSpec, responses []*HttpResponseSpec) *VerbSpec {
	return newVerbSpec("PATCH", description, params, body, responses)
}

// HEAD is not allowed to have a request body, as defined per RFC 7231.
func HEAD(description string, params []QueryParamsOrHeaderParamsOrPathParams, responses []*HttpResponseSpec) *VerbSpec {
	return newVerbSpec("HEAD", description, params, nil, responses)
}

// MimeType is a string which represents a media type (mime type) as defined by RFC 6838.
type MimeType string

const (
	Json MimeType = "application/json"
	Xml  MimeType = "application/xml"
	Any  MimeType = "application/octet-stream"
)

// RequestBody is a factory for a slice of ContentSpec.
func RequestBody(content ...*ContentSpec) []*ContentSpec {
	return content
}

// QueryParamsOrHeaderParamsOrPathParams is a marker interface for the sum type of parameters for query, header or
// path.
type QueryParamsOrHeaderParamsOrPathParams interface {
	queryParamsOrHeaderParamsOrPathParams()
}

// QueryParams is a slice of ParamSpec but used to describe query parameters.
type QueryParams []*ParamSpec

// queryParamsOrHeaderParamsOrPathParams identifies the according sum type.
func (q QueryParams) queryParamsOrHeaderParamsOrPathParams() {
	panic("marker interface")
}

// Query is a factory for a bunch of ParamSpec to create HeaderParams.
func Query(params ...*ParamSpec) QueryParams {
	return params
}

// HeaderParams is a slice of ParamSpec but used to describe header parameters.
type HeaderParams []*ParamSpec

// queryParamsOrHeaderParamsOrPathParams identifies the according sum type.
func (q HeaderParams) queryParamsOrHeaderParamsOrPathParams() {
	panic("marker interface")
}

// Header is a factory for a bunch of ParamSpec to create HeaderParams.
func Header(params ...*ParamSpec) HeaderParams {
	return params
}

// PathParams is a slice of ParamSpec but used to describe header parameters.
type PathParams []*ParamSpec

// queryParamsOrHeaderParamsOrPathParams identifies the according sum type.
func (q PathParams) queryParamsOrHeaderParamsOrPathParams() {
	panic("marker interface")
}

// Path is a factory for a bunch of ParamSpec to create PathParams.
func Path(params ...*ParamSpec) PathParams {
	return params
}

// Parameters is a factory to create a slice of the Parameters triple sum type.
func Parameters(params ...QueryParamsOrHeaderParamsOrPathParams) []QueryParamsOrHeaderParamsOrPathParams {
	return params
}

// ContentSpec aggregates a mime type which may also indicate the serialization format
// for a specific type.
type ContentSpec struct {
	mimeType MimeType
	typeName TypeName
}

// Content is a factory method for a ContentSpec. It is a confusing mix of passive
// information and/or active serialization strategy. The validator should bail out
// for any invalid or at least unsupported combinations. For example you cannot
// combine an integer with json or a struct type as a jpeg. However an io.Reader
// or a slice of bytes can be combined with any mime type.
func Content(mime MimeType, typeName TypeName) *ContentSpec {
	return &ContentSpec{
		mimeType: mime,
		typeName: typeName,
	}
}

// Resource is a factory for a HttpResourceSpec. It defines a path, containing potential path parameters, and multiple
// verb specifiers. Parameters can be defined globally.
func Resource(path, comment string, params []QueryParamsOrHeaderParamsOrPathParams, verbs ...*VerbSpec) *HttpResourceSpec {
	return &HttpResourceSpec{
		path:        path,
		description: comment,
		params:      params,
		verbs:       verbs,
	}
}

// HttpResponseSpec represents an https status code and various body mime type variants.
type HttpResponseSpec struct {
	comment      string
	statusCode   int
	headers      HeaderParams
	bodyVariants []*ContentSpec
}

// Responses is a factory for a slice of HttpResponseSpec.
func Responses(responses ...*HttpResponseSpec) []*HttpResponseSpec {
	return responses
}

// Response is a factory for a HttpResponseSpec and defines a http response with a status code
// and multiple body variants, which are usually selected by the incoming body.
func Response(status int, comment string, headers HeaderParams, content ...*ContentSpec) *HttpResponseSpec {
	return &HttpResponseSpec{
		comment:      comment,
		statusCode:   status,
		headers:      headers,
		bodyVariants: content,
	}
}
