package ddd

// A CoreLayerSpec represents a stereotyped CORE layer.
type CoreLayerSpec struct {
	name        string
	description string
	api         []StructOrInterface
	factories   []FuncOrStruct
}

// Name of the Layer
func (c *CoreLayerSpec) Name() string {
	return c.name
}

// Description of the layer
func (c *CoreLayerSpec) Description() string {
	return c.description
}

// Stereotype of the layer
func (c *CoreLayerSpec) Stereotype() Stereotype {
	return CORE
}

// Core has never any dependencies to any other layer.
func Core(api []StructOrInterface, factories []FuncOrStruct) *CoreLayerSpec {
	return &CoreLayerSpec{
		name: "core",
		description: "Package core contains all domain specific models for the current bounded context.\n" +
			"It contains an exposed public API to be imported by other layers and an internal package \n" +
			"private implementation accessible by factory functions.",
		api:       api,
		factories: factories,
	}
}
