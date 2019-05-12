package layer

// Shape represents a lottie shape layer.
type Shape interface {
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
	Ind *int `json:"ind,omitempty"`

	// In Point of layer. Sets the initial frame of the layer.
	Ip *int `json:"ip,omitempty"`

	// Shape list of items
	// It []*ItemsItems `json:"it,omitempty"`

	// Transform properties
	// Ks *Transform `json:"ks,omitempty"`

	// Parsed layer name used as html id on SVG/HTML renderer
	Ln *string `json:"ln,omitempty"`

	// List of Masks
	// MaskProperties []*MaskPropertiesItems `json:"maskProperties,omitempty"`

	// After Effects Layer Name. Used for expressions.
	Nm *string `json:"nm,omitempty"`

	// Out Point of layer. Sets the final frame of the layer.
	Op *int `json:"op,omitempty"`

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
