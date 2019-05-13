package layer

import (
	"encoding/json"

	"github.com/gmlewis/lottie2flare/lottie/helpers"
)

// PreComp represents a lottie preComp layer.
type PreComp interface {
	// GetTransform returns the shape's transform.
	GetTransform() *helpers.Transform

	// Type returns PreCompType.
	Type() Type
}

// PreCompT implements the PreComp and Layer interface.
type PreCompT struct {
	// Auto-Orient along path AE property.
	AutoOrient *int `json:"ao,omitempty"`

	// Blend Mode
	BlendMode *int `json:"bm,omitempty"`

	// Parsed layer name used as html class on SVG/HTML renderer
	Class *string `json:"cl,omitempty"`

	// 3D layer flag
	Ddd *int `json:"ddd,omitempty"`

	// List of Effects
	Effects []json.RawMessage `json:"ef,omitempty"`

	// Boolean when layer has a mask. Will be deprecated in favor of checking maskProperties.
	HasMask *bool `json:"hasMask,omitempty"`

	// Layer index in AE. Used for parenting and expressions.
	Index *int `json:"ind,omitempty"`

	// In Point of layer. Sets the initial frame of the layer.
	InPoint *int `json:"ip,omitempty"`

	// Transform properties
	Transform *helpers.Transform `json:"ks,omitempty"`

	// Parsed layer name used as html id on SVG/HTML renderer
	LayerName *string `json:"ln,omitempty"`

	// List of Masks
	MaskProperties []json.RawMessage `json:"maskProperties,omitempty"`

	// After Effects Layer Name. Used for expressions.
	Name *string `json:"nm,omitempty"`

	// Out Point of layer. Sets the final frame of the layer.
	OutPoint *int `json:"op,omitempty"`

	// Layer Parent. Uses ind of parent.
	Parent *int `json:"parent,omitempty"`

	// RefID pointing to the source composition defined on 'assets' object
	RefID *string `json:"refId,omitempty"`

	// Layer Time Stretching
	TimeStretching *float64 `json:"sr,omitempty"`

	// Start Time of layer. Sets the start time of the layer.
	StartTime *float64 `json:"st,omitempty"`

	// Comp's Time remapping
	TimeMapping *float64 `json:"tm,omitempty"`

	// Type of layer: Precomp.
	Ty Type `json:"ty,omitempty"`
}

// Type returns the layer type.
func (p *PreCompT) Type() Type {
	return PreCompType
}

// GetTransform returns the shape's transform.
func (p *PreCompT) GetTransform() *helpers.Transform {
	return p.Transform
}
