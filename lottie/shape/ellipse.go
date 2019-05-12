package shape

// Ellipse represents a lottie shape/ellipse.
type Ellipse interface {
	Type() Type
}

// EllipseT implements the Ellipse interface.
type EllipseT struct {
}

// Type returns the layer type.
func (s *EllipseT) Type() Type {
	return EllipseType
}
