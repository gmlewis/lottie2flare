package layer

import (
	"encoding/json"

	"github.com/gmlewis/lottie2flare/lottie/helpers"
	ls "github.com/gmlewis/lottie2flare/lottie/shape"
)

const (
	// ShapeType is a lottie shape layer.
	ShapeType Type = 4
)

// Shape represents a lottie shape layer.
type Shape interface {
	// GetIndex returns the shape index.
	GetIndex() int

	// GetName returns the shape name.
	GetName() string

	// GetShapes returns the list of shape items.
	GetShapes() []ls.Shape

	// GetTransform returns the shape's transform.
	GetTransform() *helpers.Transform

	// Type returns ShapeType.
	Type() Type
}

// ShapeT implements the Shape and Layer interface.
type ShapeT struct {

	// Auto-Orient along path AE property.
	AutoOrient *int `json:"ao,omitempty"`

	// Blend Mode
	BlendMode *int `json:"bm,omitempty"`

	// Parsed layer name used as html class on SVG/HTML renderer
	Class *string `json:"cl,omitempty"`

	// 3d layer flag
	Ddd *int `json:"ddd,omitempty"`

	// List of Effects
	Effects []json.RawMessage `json:"ef,omitempty"`

	// Boolean when layer has a mask. Will be deprecated in favor of checking maskProperties.
	HasMask *bool `json:"hasMask,omitempty"`

	// Layer index in AE. Used for parenting and expressions.
	Index *int `json:"ind,omitempty"`

	// In Point of layer. Sets the initial frame of the layer.
	InPoint *float64 `json:"ip,omitempty"`

	// Shape list of items.
	Shapes []json.RawMessage `json:"shapes,omitempty"`

	// Transform properties
	Transform *helpers.Transform `json:"ks,omitempty"`

	// Parsed layer name used as html id on SVG/HTML renderer
	LayerName *string `json:"ln,omitempty"`

	// List of Masks
	MaskProperties []json.RawMessage `json:"maskProperties,omitempty"`

	// After Effects Layer Name. Used for expressions.
	Name *string `json:"nm,omitempty"`

	// Out Point of layer. Sets the final frame of the layer.
	OutPoint *float64 `json:"op,omitempty"`

	// Layer Parent. Uses ind of parent.
	Parent *int `json:"parent,omitempty"`

	// Layer Time Stretching
	TimeStretching *float64 `json:"sr,omitempty"`

	// Start Time of layer. Sets the start time of the layer.
	StartTime *float64 `json:"st,omitempty"`

	// Type of layer: Shape.
	Ty Type `json:"ty,omitempty"`
}

// Type returns the layer type.
func (s *ShapeT) Type() Type {
	return ShapeType
}

// GetShapes returns the shapes casted to their appropriate types.
func (s *ShapeT) GetShapes() (result []ls.Shape) {
	for _, v := range s.Shapes {
		result = append(result, ls.New(v))
	}
	return result
}

// GetTransform returns the shape's transform.
func (s *ShapeT) GetTransform() *helpers.Transform {
	return s.Transform
}
