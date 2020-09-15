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
)

// A RestServerSpec to declare various endpoints.
type RestServerSpec struct {
	url         string
	description string
	pos         Pos
}

// Server creates a new endpoint.
func Server(url, description string) *RestServerSpec {
	return &RestServerSpec{
		url:         url,
		description: description,
		pos:         capturePos("Server", 1),
	}
}

// Description returns the comment.
func (s *RestServerSpec) Description() string {
	return s.description
}

// Pos returns the debugging position.
func (s *RestServerSpec) Pos() Pos {
	return s.pos
}

// Url returns the servers url.
func (s *RestServerSpec) Url() string {
	return s.url
}

// Servers returns a slice of RestServerSpec
func Servers(hosts ...*RestServerSpec) []*RestServerSpec {
	return hosts
}

// A RestLayerSpec represents a stereotyped PRESENTATION Layer.
type RestLayerSpec struct {
	name        string
	version     string
	description string
	resources   []*HttpResourceSpec
	hosts       []*RestServerSpec
	pos         Pos
}

// REST is a driven port and directly uses the UseCases API. It does normally not define any further models
// than the use cases it imports. Because of transitivity, it has also access to the CORE layer API models.
//
// Version must be a semantic version string like "v1.2.3".
func REST(version string, hosts []*RestServerSpec, resources []*HttpResourceSpec) *RestLayerSpec {
	major := semver.Major(version)
	name := "Rest" + major
	return &RestLayerSpec{
		name:    name,
		hosts:   hosts,
		version: version,
		description: "Package rest contains the REST specific implementation for the current bounded context.\n" +
			"It depends only from the use cases and transitively on the core API.",
		resources: resources,
		pos:       capturePos("REST", -1),
	}
}

// PrimaryUrl returns always and http|s://... url.
func (r *RestLayerSpec) PrimaryUrl() string {
	if len(r.hosts) == 0 {
		return "http://localhost:8080"
	}

	return r.hosts[0].url
}

// Pos returns the debug position.
func (r *RestLayerSpec) Pos() Pos {
	return r.pos
}

// Prefix returns by default /api/vX/ where x is the major version from the semver version.
func (r *RestLayerSpec) Prefix() string {
	return "/api/" + semver.Major(r.version) + "/"
}

// Resources returns the declared resources.
func (r *RestLayerSpec) Resources() []*HttpResourceSpec {
	return r.resources
}

// Walk inspects all children
func (r *RestLayerSpec) Walk(f func(obj interface{}) error) error {
	if err := f(r); err != nil {
		return err
	}

	for _, obj := range r.hosts {
		if err := f(obj); err != nil {
			return err
		}
	}

	for _, obj := range r.resources {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	return nil
}

// Version returns the semantic version
func (r *RestLayerSpec) Version() string {
	return r.version
}

// MajorVersion returns something like v1
func (r *RestLayerSpec) MajorVersion() string {
	return semver.Major(r.version)
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

// HttpResourceSpec represents a REST resource (e.g. /book/:id)
type HttpResourceSpec struct {
	path        string
	description string
	params      []QueryParamsOrHeaderParamsOrPathParams
	verbs       []*VerbSpec
	pos         Pos
}

// Resource is a factory for a HttpResourceSpec. It defines a path, containing potential path parameters, and multiple
// verb specifiers. Parameters can be defined globally.
func Resource(path, comment string, params []QueryParamsOrHeaderParamsOrPathParams, verbs ...*VerbSpec) *HttpResourceSpec {
	return &HttpResourceSpec{
		path:        path,
		description: comment,
		params:      params,
		verbs:       verbs,
		pos:         capturePos("Resource", 1),
	}
}

// Pos returns the debugging position.
func (r *HttpResourceSpec) Pos() Pos {
	return r.pos
}

// Path returns resource path
func (r *HttpResourceSpec) Path() string {
	return r.path
}

// Description of the resource.
func (r *HttpResourceSpec) Description() string {
	return r.description
}

// Params returns any params.
func (r *HttpResourceSpec) Params() []QueryParamsOrHeaderParamsOrPathParams {
	return r.params
}

// Verbs returns the defined http verbs.
func (r *HttpResourceSpec) Verbs() []*VerbSpec {
	return r.verbs
}

// Walk loops over self and children
func (r *HttpResourceSpec) Walk(f func(obj interface{}) error) error {
	if err := f(r); err != nil {
		return err
	}

	for _, p := range r.params {
		for _, obj := range p.Params() {
			if err := f(obj); err != nil {
				return err
			}
		}

	}

	for _, obj := range r.verbs {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	return nil
}

// Resources is a factory for a slice of HttpResourceSpec
func Resources(resources ...*HttpResourceSpec) []*HttpResourceSpec {
	return resources
}

// VerbSpec models a HTTP Verb like GET, POST, PATCH, PUT, DELETE.
type VerbSpec struct {
	method      string
	description string
	params      []QueryParamsOrHeaderParamsOrPathParams
	body        []*ContentSpec
	responses   []*HttpResponseSpec
	pos         Pos
}

// newVerbSpec creates a new verb specifiction.
func newVerbSpec(method, description string, params []QueryParamsOrHeaderParamsOrPathParams, body []*ContentSpec, responses []*HttpResponseSpec) *VerbSpec {
	verb := &VerbSpec{
		method:      method,
		description: description,
		body:        body,
		responses:   responses,
		params:      params,
		pos:         capturePos(method, 2),
	}

	return verb
}

// Pos returns the debugging position.
func (v *VerbSpec) Pos() Pos {
	return v.pos
}

// Method returns the http method like GET or DELETE.
func (v *VerbSpec) Method() string {
	return v.method
}

// Description returns the textual description of the verb.
func (v *VerbSpec) Description() string {
	return v.description
}

// Params returns the declared http header, path or query parameter.
func (v *VerbSpec) Params() []QueryParamsOrHeaderParamsOrPathParams {
	return v.params
}

// Body returns the optional body specifications.
func (v *VerbSpec) Body() []*ContentSpec {
	return v.body
}

// Responses returns the http response specifications.
func (v *VerbSpec) Responses() []*HttpResponseSpec {
	return v.responses
}

// Walk loops over self and children
func (v *VerbSpec) Walk(f func(obj interface{}) error) error {
	if err := f(v); err != nil {
		return err
	}

	for _, params := range v.params {
		for _, obj := range params.Params() {
			if err := f(obj); err != nil {
				return err
			}
		}

	}

	for _, obj := range v.responses {
		if err := obj.Walk(f); err != nil {
			return err
		}
	}

	for _, obj := range v.body {
		if err := f(obj); err != nil {
			return err
		}
	}

	return nil
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
	Params() []*ParamSpec
}

// QueryParams is a slice of ParamSpec but used to describe query parameters.
type QueryParams []*ParamSpec

// queryParamsOrHeaderParamsOrPathParams identifies the according sum type.
func (q QueryParams) queryParamsOrHeaderParamsOrPathParams() {
	panic("marker interface")
}

// Params returns the actual parameters.
func (q QueryParams) Params() []*ParamSpec {
	return q
}

// Query is a factory for a bunch of ParamSpec to create HeaderParams.
func Query(params ...*ParamSpec) QueryParams {
	for _, p := range params {
		p.nameValidationKind = NVHttpQueryParameter
	}
	return params
}

// HeaderParams is a slice of ParamSpec but used to describe header parameters.
type HeaderParams []*ParamSpec

// queryParamsOrHeaderParamsOrPathParams identifies the according sum type.
func (q HeaderParams) queryParamsOrHeaderParamsOrPathParams() {
	panic("marker interface")
}

// Params returns the actual parameters.
func (q HeaderParams) Params() []*ParamSpec {
	return q
}

// Header is a factory for a bunch of ParamSpec to create HeaderParams.
func Header(params ...*ParamSpec) HeaderParams {
	for _, p := range params {
		p.nameValidationKind = NVHttpHeaderParameter
	}
	return params
}

// PathParams is a slice of ParamSpec but used to describe header parameters.
type PathParams []*ParamSpec

// queryParamsOrHeaderParamsOrPathParams identifies the according sum type.
func (q PathParams) queryParamsOrHeaderParamsOrPathParams() {
	panic("marker interface")
}

// Params returns the actual parameters.
func (q PathParams) Params() []*ParamSpec {
	return q
}

// Path is a factory for a bunch of ParamSpec to create PathParams.
func Path(params ...*ParamSpec) PathParams {
	for _, p := range params {
		p.nameValidationKind = NVHttpPathParameter
	}
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

// Walk loops over self and children
func (r *HttpResponseSpec) Walk(f func(obj interface{}) error) error {
	if err := f(r); err != nil {
		return err
	}

	for _, obj := range r.headers {
		if err := f(obj); err != nil {
			return err
		}
	}

	for _, obj := range r.bodyVariants {
		if err := f(obj); err != nil {
			return err
		}
	}

	return nil
}
