package ddd

// REST is a driven port and directly uses the UseCases API
func REST(version string, resources []*HttpResourceSpec, types *TypeSpecs) *LayerSpec {
	return nil
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

func GET(comment string, in *ParamSpecs, responses []*HttpResponseSpec) *VerbSpec {
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
