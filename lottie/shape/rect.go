package shape

// Rect represents a lottie shape/rect.
type Rect interface {
	Type() Type
}

// RectT implements the Rect interface.
type RectT struct {
}

// Type returns the layer type.
func (s *RectT) Type() Type {
	return RectType
}
