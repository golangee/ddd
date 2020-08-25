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
	method      string
	path        string
	description string
}

// Resources is a factory for a slice of HttpResourceSpec
func Resources(resources ...*HttpResourceSpec) []*HttpResourceSpec {
	return resources
}

// GET declares the according verb and is not allowed to have a request body, as defined per RFC 7231.
func GET(comment string, in []*ParamSpec, responses []*HttpResponseSpec) *VerbSpec {
	return nil
}

func POST(description string, body *HttpRequestBodySpec) *VerbSpec {
	return nil
}

type HttpRequestBodySpec struct {
}

type HttpContentSpec struct {
}

func RequestBody(required bool, content ...*ContentSpec) *HttpRequestBodySpec {
	return nil
}

type ContentSpec struct {
	mimeType MimeType
	typeName TypeName
}

func Content(mime MimeType, typeName TypeName) *ContentSpec {
	return &ContentSpec{
		mimeType: mime,
		typeName: typeName,
	}
}

type HttpRequestSpec struct{}

// DELETE is not allowed to have a request body, as defined per RFC 7231.
func DELETE(comment string) *VerbSpec {
	return nil
}

func PUT(comment string) *VerbSpec {
	return nil
}

// HEAD is not allowed to have a request body, as defined per RFC 7231.
func HEAD(comment string) *VerbSpec {
	return nil
}

type MimeType string

const (
	MimeTypeJson MimeType = "application/json"
	MimeTypeXml  MimeType = "application/xml"
	MimeTypeText MimeType = "application/text"
)

func Resource(path, comment string, verbs ...*VerbSpec) *HttpResourceSpec {
	return &HttpResourceSpec{path: path, description: comment}
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
