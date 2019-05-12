package layer

// PreComp represents a lottie preComp layer.
type PreComp interface {
	Type() Type
}

// PreCompT implements the PreComp interface.
type PreCompT struct {
	Ty Type `json:"ty,omitempty"`
}

// Type returns the layer type.
func (s *PreCompT) Type() Type {
	return PreCompType
}
