package ddd

// A Layer can be of very different kind. Its dependencies depends on the Stereotype and its concrete payload
// on the actual kind of the layer. Different IMPLEMENTATION types provides distinct types.
type Layer interface {
	// Name of the layer
	Name() string

	// Description of the layer
	Description() string

	// Stereotype of the layer
	Stereotype() Stereotype
}
