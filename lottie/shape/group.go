package shape

import "encoding/json"

const (
	// GroupType represents a lottie shape/group.
	GroupType Type = "gr"
)

// Group represents a lottie shape/group.
type Group interface {
	// GetShapes returns the shapes in the group.
	GetShapes() []Shape

	// Type returns GroupType.
	Type() Type
}

// GroupT implements the Group interface.
type GroupT struct {
	// Group list of shapes.
	Shapes []json.RawMessage `json:"it,omitempty"`

	// After Effect's Match Name. Used for expressions.
	MatchName *string `json:"mn,omitempty"`

	// After Effect's Name. Used for expressions.
	Name *string `json:"nm,omitempty"`

	// Group number of properties. Used for expressions.
	NumProperties *int `json:"np,omitempty"`

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
