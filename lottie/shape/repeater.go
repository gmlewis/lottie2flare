package shape

// Repeater represents a lottie shape/repeater.
type Repeater interface {
	Type() Type
}

// RepeaterT implements the Repeater interface.
type RepeaterT struct {
}

// Type returns the layer type.
func (s *RepeaterT) Type() Type {
	return RepeaterType
}
