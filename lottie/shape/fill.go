package shape

// Fill represents a lottie shape/fill.
type Fill interface {
	Type() Type
}

// FillT implements the Fill interface.
type FillT struct {
}

// Type returns the layer type.
func (s *FillT) Type() Type {
	return FillType
}
