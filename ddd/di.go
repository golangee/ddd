package ddd

type ConversionSpec struct {
}

func ConvertStruct(fromLayer string, fromName TypeName, toLayer string, toName TypeName) *ConversionSpec {
	return nil
}

func DependencyInjection(conversions ...*ConversionSpec) *LayerSpec {
	return nil
}
