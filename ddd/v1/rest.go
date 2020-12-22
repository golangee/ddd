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

// MimeType is a string which represents a media type (mime type) as defined by RFC 6838.
type MimeType string

const (
	Json MimeType = "application/json"
	Xml  MimeType = "application/xml"
	Any  MimeType = "application/octet-stream"
)

// A RestServerSpec to declare various endpoints.
type RestServerSpec struct {
	url         string
	description string
	pos         Pos
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

// Query is a factory for a bunch of ParamSpec to create HeaderParams.
func Query(params ...*ParamSpec) QueryParams {
	for _, p := range params {
		p.nameValidationKind = NVHttpQueryParameter
	}
	return params
}

// Header is a factory for a bunch of ParamSpec to create HeaderParams.
func Header(params ...*ParamSpec) HeaderParams {
	for _, p := range params {
		p.nameValidationKind = NVHttpHeaderParameter
	}
	return params
}


// Path is a factory for a bunch of ParamSpec to create PathParams.
func Path(params ...*ParamSpec) PathParams {
	for _, p := range params {
		p.nameValidationKind = NVHttpPathParameter
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

// Parameters is a factory to create a slice of the Parameters triple sum type.
func Parse(params ...QueryParamsOrHeaderParamsOrPathParams) []QueryParamsOrHeaderParamsOrPathParams {
	return params
}

type ProvidedParamSpecs []*ParamSpec

func Provide(params ...*ParamSpec) ProvidedParamSpecs {
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
