package shape

// Fill represents a lottie shape/fill.
type Fill interface {
	// GetColor returns an RGBA color (0-1).
	GetColor() [4]float64

	// Type returns FillType.
	Type() Type
}

// FillT implements the Fill interface.
type FillT struct {
}

// Type returns the layer type.
func (s *FillT) Type() Type {
	return FillType
}
