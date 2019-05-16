package shape

const (
	// GFillType represents a lottie shape/gFill.
	GFillType Type = "gf"
)

// GFill represents a lottie shape/gFill.
type GFill interface {
	Type() Type
}

// GFillT implements the GFill interface.
type GFillT struct {
}

// Type returns the layer type.
func (s *GFillT) Type() Type {
	return GFillType
}
