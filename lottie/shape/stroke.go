package shape

const (
	// StrokeType represents a lottie shape/stroke.
	StrokeType Type = "st"
)

// Stroke represents a lottie shape/stroke.
type Stroke interface {
	Type() Type
}

// StrokeT implements the Stroke interface.
type StrokeT struct {
}

// Type returns the layer type.
func (s *StrokeT) Type() Type {
	return StrokeType
}
