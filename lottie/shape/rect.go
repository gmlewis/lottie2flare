package shape

const (
	// RectType represents a lottie shape/rect.
	RectType Type = "rc"
)

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
