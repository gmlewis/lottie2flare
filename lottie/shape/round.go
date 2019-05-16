package shape

const (
	// RoundType represents a lottie shape/round.
	RoundType Type = "rd"
)

// Round represents a lottie shape/round.
type Round interface {
	Type() Type
}

// RoundT implements the Round interface.
type RoundT struct {
}

// Type returns the layer type.
func (s *RoundT) Type() Type {
	return RoundType
}
