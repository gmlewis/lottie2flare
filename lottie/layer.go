package lottie

type LayerType int

const (
	LayerPreComp LayerType = 0
	LayerImage   LayerType = 2
	LayerNull    LayerType = 3
	LayerShape   LayerType = 4
)

// Layer represents a lottie layer.
type Layer struct {
	// Common fields:

	// Auto-Orient along path AE property.
	Ao *float64 `json:"ao,omitempty"`

	// Blend Mode
	Bm *float64 `json:"bm,omitempty"`

	// Parsed layer name used as html class on SVG/HTML renderer
	Cl *string `json:"cl,omitempty"`

	// 3d layer flag
	Ddd *float64 `json:"ddd,omitempty"`

	// List of Effects
	Ef []*EffectsItems `json:"ef,omitempty"`

	// Boolean when layer has a mask. Will be deprecated in favor of checking maskProperties.
	HasMask *bool `json:"hasMask,omitempty"`

	// Layer index in AE. Used for parenting and expressions.
	Ind *float64 `json:"ind,omitempty"`

	// In Point of layer. Sets the initial frame of the layer.
	Ip *float64 `json:"ip,omitempty"`

	// Transform properties
	Ks *Transform `json:"ks,omitempty"`

	// Parsed layer name used as html id on SVG/HTML renderer
	Ln *string `json:"ln,omitempty"`

	// List of Masks
	MaskProperties []*MaskPropertiesItems `json:"maskProperties,omitempty"`

	// After Effects Layer Name. Used for expressions.
	Nm *string `json:"nm,omitempty"`

	// Out Point of layer. Sets the final frame of the layer.
	Op *float64 `json:"op,omitempty"`

	// Layer Parent. Uses ind of parent.
	Parent *float64 `json:"parent,omitempty"`

	// id pointing to the source composition or image defined on 'assets' object
	RefId *string `json:"refId,omitempty"`

	// Layer Time Stretching
	Sr *float64 `json:"sr,omitempty"`

	// Start Time of layer. Sets the start time of the layer.
	St *float64 `json:"st,omitempty"`

	// Type of layer.
	Ty *LayerType `json:"ty,omitempty"`

	// From layers/shape:

	// Shape list of items
	Shapes []*ShapeItems `json:"shapes,omitempty"`

	// From layers/solid:

	// Color of the solid in hex
	Sc *string `json:"sc,omitempty"`

	// Height of the solid.
	Sh *float64 `json:"sh,omitempty"`

	// Width of the solid.
	Sw *float64 `json:"sw,omitempty"`

	// From layers/precomp:

	// Comp's Time remapping
	Tm *ValueKeyframed `json:"tm,omitempty"`

	// From layers/image:

	// From layers/null:

	// From layers/text:

	// Text Data
	T []*TextDataItems `json:"t,omitempty"`
}
