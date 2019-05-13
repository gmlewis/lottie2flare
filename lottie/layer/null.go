package layer

import (
	"encoding/json"

	"github.com/gmlewis/lottie2flare/lottie/helpers"
)

// Null represents a lottie null layer.
type Null interface {
	Type() Type
}

// NullT implements the Null and Layer interface.
type NullT struct {
	// Auto-Orient along path AE property.
	AutoOrient *int `json:"ao,omitempty"`

	// Parsed layer name used as html class on SVG/HTML renderer
	Class *string `json:"cl,omitempty"`

	// 3D layer flag
	Ddd *int `json:"ddd,omitempty"`

	// List of Effects
	Effects []json.RawMessage `json:"ef,omitempty"`

	// Layer index in AE. Used for parenting and expressions.
	Ind *int `json:"ind,omitempty"`

	// In Point of layer. Sets the initial frame of the layer.
	InPoint *int `json:"ip,omitempty"`

	// Transform properties
	Transform *helpers.Transform `json:"ks,omitempty"`

	// Parsed layer name used as html id on SVG/HTML renderer
	LayerName *string `json:"ln,omitempty"`

	// After Effects Layer Name. Used for expressions.
	Name *float64 `json:"nm,omitempty"`

	// Out Point of layer. Sets the final frame of the layer.
	OutPoint *int `json:"op,omitempty"`

	// Layer Parent. Uses ind of parent.
	Parent *int `json:"parent,omitempty"`

	// Layer Time Stretching
	TimeStretching *float64 `json:"sr,omitempty"`

	// Start Time of layer. Sets the start time of the layer.
	StartTime *float64 `json:"st,omitempty"`

	// Type of layer: Null.
	Ty Type `json:"ty,omitempty"`
}

// Type returns the layer type.
func (n *NullT) Type() Type {
	return NullType
}

// GetTransform returns the shape's transform.
func (n *NullT) GetTransform() *helpers.Transform {
	return n.Transform
}
