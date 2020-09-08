package ddd

type ExternalParamSpec ParamSpec

type ExternalParams []*ParamSpec

type ConfigSpec struct {
}

type RequireInterfaceSpec struct {
	interfaceName string
}

func Configure(interfaceName string, require []*RequireInterfaceSpec, params ExternalParams) *ConfigSpec {
	return &ConfigSpec{}
}

func Require(interfaceNames ...string) []*RequireInterfaceSpec {
	var r []*RequireInterfaceSpec
	for _, name := range interfaceNames {
		r = append(r, &RequireInterfaceSpec{interfaceName: name})
	}
	return r
}

func External(params ...*ExternalParamSpec) ExternalParams {
	var r []*ParamSpec
	for _, param := range params {
		r = append(r, (*ParamSpec)(param))
	}

	return r
}

func Env(name string,typeName TypeName,defaultValue string,comment string)*ExternalParamSpec{
	return &ExternalParamSpec{}
}