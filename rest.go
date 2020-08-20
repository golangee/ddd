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

func GET(comment string, in *ParamSpecs, out *ParamSpecs) *VerbSpec {
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

type VerbSpec struct {
}
