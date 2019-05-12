package layer

import (
	"encoding/json"

	"github.com/gmlewis/lottie2flare/lottie/helpers"
	ls "github.com/gmlewis/lottie2flare/lottie/shape"
)

// Shape represents a lottie shape layer.
type Shape interface {
	// GetAo returns the Ao field if it's non-nil, zero value otherwise.
	GetAo() int

	// GetBm returns the Bm field if it's non-nil, zero value otherwise.
	GetBm() int

	// GetCl returns the Cl field if it's non-nil, zero value otherwise.
	GetCl() string

	// GetDdd returns the Ddd field if it's non-nil, zero value otherwise.
	GetDdd() int

	// GetHasMask returns the HasMask field if it's non-nil, zero value otherwise.
	GetHasMask() bool

	// GetIndex returns the Index field if it's non-nil, zero value otherwise.
	GetIndex() int

	// GetInPoint returns the InPoint field if it's non-nil, zero value otherwise.
	GetInPoint() int

	// GetLn returns the Ln field if it's non-nil, zero value otherwise.
	GetLn() string

	// GetName returns the Name field if it's non-nil, zero value otherwise.
	GetName() string

	// GetOutPoint returns the OutPoint field if it's non-nil, zero value otherwise.
	GetOutPoint() int

	// GetParent returns the Parent field if it's non-nil, zero value otherwise.
	GetParent() int

	GetShapes() []ls.Shape

	// GetSr returns the Sr field if it's non-nil, zero value otherwise.
	GetSr() float64

	// GetSt returns the St field if it's non-nil, zero value otherwise.
	GetSt() float64

	// GetTransform returns the shape's transform.
	GetTransform() *helpers.Transform

	// Type returns ShapeType.
	Type() Type
}

// ShapeT implements the Shape interface.
type ShapeT struct {

	// Auto-Orient along path AE property.
	Ao *int `json:"ao,omitempty"`

	// Blend Mode
	Bm *int `json:"bm,omitempty"`

	// Parsed layer name used as html class on SVG/HTML renderer
	Cl *string `json:"cl,omitempty"`

	// 3d layer flag
	Ddd *int `json:"ddd,omitempty"`

	// List of Effects
	// Ef []*EffectsItems `json:"ef,omitempty"`

	// Boolean when layer has a mask. Will be deprecated in favor of checking maskProperties.
	HasMask *bool `json:"hasMask,omitempty"`

	// Layer index in AE. Used for parenting and expressions.
	Index *int `json:"ind,omitempty"`

	// In Point of layer. Sets the initial frame of the layer.
	InPoint *int `json:"ip,omitempty"`

	// Shape list of items.
	Shapes []json.RawMessage `json:"it,omitempty"`

	// Transform properties
	Transform *helpers.Transform `json:"ks,omitempty"`

	// Parsed layer name used as html id on SVG/HTML renderer
	Ln *string `json:"ln,omitempty"`

	// List of Masks
	// MaskProperties []*MaskPropertiesItems `json:"maskProperties,omitempty"`

	// After Effects Layer Name. Used for expressions.
	Name *string `json:"nm,omitempty"`

	// Out Point of layer. Sets the final frame of the layer.
	OutPoint *int `json:"op,omitempty"`

	// Layer Parent. Uses ind of parent.
	Parent *int `json:"parent,omitempty"`

	// Layer Time Stretching
	Sr *float64 `json:"sr,omitempty"`

	// Start Time of layer. Sets the start time of the layer.
	St *float64 `json:"st,omitempty"`

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
