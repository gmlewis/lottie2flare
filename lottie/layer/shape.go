package layer

// Shape represents a lottie shape layer.
type Shape interface {
	Type() Type
}

// ShapeT implements the Shape interface.
type ShapeT struct {
	Ty Type `json:"ty,omitempty"`
}

// Type returns the layer type.
func (s *ShapeT) Type() Type {
	return ShapeType
}

// NewShapeT returns a ShapeT from the JSON map.
func NewShapeT(m map[string]interface{}) *ShapeT {
	return &ShapeT{
		Ty: ShapeType,
	}
}
