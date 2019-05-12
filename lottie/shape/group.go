package shape

import "encoding/json"

// Group represents a lottie shape/group.
type Group interface {
	Type() Type
}

// GroupT implements the Group interface.
type GroupT struct {
	// Group list of shapes.
	Shapes []json.RawMessage `json:"it,omitempty"`

	// After Effect's Match Name. Used for expressions.
	Mn *string `json:"mn,omitempty"`

	// After Effect's Name. Used for expressions.
	Nm *string `json:"nm,omitempty"`

	// Group number of properties. Used for expressions.
	Np *float64 `json:"np,omitempty"`

	// Shape content type.
	Ty *string `json:"ty,omitempty"`
}

// Type returns the layer type.
func (s *GroupT) Type() Type {
	return GroupType
}

// GetShapes returns the shapes casted to their appropriate types.
func (s *GroupT) GetShapes() (result []Shape) {
	for _, v := range s.Shapes {
		result = append(result, New(v))
	}
	return result
}
