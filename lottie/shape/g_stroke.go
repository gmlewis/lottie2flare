package shape

// GStroke represents a lottie shape/gStroke.
type GStroke interface {
	Type() Type
}

// GStrokeT implements the GStroke interface.
type GStrokeT struct {
}

// Type returns the layer type.
func (s *GStrokeT) Type() Type {
	return GStrokeType
}
