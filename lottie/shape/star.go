package shape

// Star represents a lottie shape/star.
type Star interface {
	Type() Type
}

// StarT implements the Star interface.
type StarT struct {
}

// Type returns the layer type.
func (s *StarT) Type() Type {
	return StarType
}
