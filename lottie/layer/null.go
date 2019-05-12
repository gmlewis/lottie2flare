package layer

// Null represents a lottie null layer.
type Null interface {
	Type() Type
}

// NullT implements the Null interface.
type NullT struct {
	Ty Type `json:"ty,omitempty"`
}

// Type returns the layer type.
func (s *NullT) Type() Type {
	return NullType
}

// NewNullT returns a NullT from the JSON map.
func NewNullT(m map[string]interface{}) *NullT {
	return &NullT{
		Ty: NullType,
	}
}
