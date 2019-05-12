package shape

// Trim represents a lottie shape/trim.
type Trim interface {
	Type() Type
}

// TrimT implements the Trim interface.
type TrimT struct {
}

// Type returns the layer type.
func (s *TrimT) Type() Type {
	return TrimType
}
