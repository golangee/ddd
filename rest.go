package ddd

func REST(version string, resources *HttpResourceSpecs,types *TypeSpecs) *PresentationSpec {
	return &PresentationSpec{version: version}
}

type HttpResourceSpec struct {
	comment string
	method  string
	path    string
}

type HttpResourceSpecs struct {
	resources []*HttpResourceSpec
}

func Resources(resources ...*HttpResourceSpec) *HttpResourceSpecs {
	return &HttpResourceSpecs{resources: resources}
}

func GET(path, comment string, in *ParamSpecs, out *ParamSpecs) *HttpResourceSpec {
	return Resource("GET", path, comment)
}

func DELETE(path, comment string) *HttpResourceSpec {
	return Resource("DELETE", path, comment)
}

func Resource(method, path, comment string) *HttpResourceSpec {
	return &HttpResourceSpec{path: path, comment: comment, method: method}
}
