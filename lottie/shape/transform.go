package shape

// Transform represents a lottie shape/transform.
type Transform interface {
	Type() Type
}

// TransformT implements the Transform interface.
type TransformT struct {
}

// Type returns the layer type.
func (s *TransformT) Type() Type {
	return TransformType
}
