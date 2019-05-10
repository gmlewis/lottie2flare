package lottie

import (
	"encoding/json"
	"io/ioutil"
)

// AssetsItems
type AssetsItems struct {
	// Common to sources/image and sources/precomp:

	// ID is the image or precomp ID.
	ID *string `json:"id,omitempty"`

	// From sources/image:

	// Height is the image height.
	Height *float64 `json:"h,omitempty"`

	// P is the image name.
	P *string `json:"p,omitempty"`

	// URL is the image path.
	URL *string `json:"u,omitempty"`

	// Width is the image width.
	Width *float64 `json:"w,omitempty"`

	// From sources/precomp:

	// List of Precomp Layers
	Layers []*LayersItems `json:"layers,omitempty"`
}

// CharsItems
type CharsItems struct {
}

type EffectsItems struct {
}

type ItemsItems struct {
}

type MaskPropertiesItems struct {
}

type TextDataItems struct {
}

type Transform struct {
}

type LayerType int

// LayersItems
type LayersItems struct {
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
	HasMask *float64 `json:"hasMask,omitempty"`

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
	Nm *float64 `json:"nm,omitempty"`

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
	It []*ItemsItems `json:"it,omitempty"`

	// From layers/solid:

	// Color of the solid in hex
	Sc *string `json:"sc,omitempty"`

	// Height of the solid.
	Sh *float64 `json:"sh,omitempty"`

	// Width of the solid.
	Sw *float64 `json:"sw,omitempty"`

	// From layers/precomp:

	// Comp's Time remapping
	Tm *float64 `json:"tm,omitempty"`

	// From layers/image:

	// From layers/null:

	// From layers/text:

	// Text Data
	T []*TextDataItems `json:"t,omitempty"`
}

// Animation represents a Lottie animation JSON blob.
type Animation struct {
	// source items that can be used in multiple places. Comps and Images for now.
	Assets []*AssetsItems `json:"assets,omitempty"`

	// source chars for text layers
	Chars []*CharsItems `json:"chars,omitempty"`

	// Composition has 3-D layers (0 or 1).
	Ddd *int `json:"ddd,omitempty"`

	// Frame Rate
	FrameRate *float64 `json:"fr,omitempty"`

	// Composition Height
	Height *float64 `json:"h,omitempty"`

	// In Point of the Time Ruler. Sets the initial Frame of the animation.
	InPoint *float64 `json:"ip,omitempty"`

	// List of Composition Layers
	Layers []*LayersItems `json:"layers,omitempty"`

	// Composition name
	Name *string `json:"nm,omitempty"`

	// Out Point of the Time Ruler. Sets the final Frame of the animation
	OutPoint *float64 `json:"op,omitempty"`

	// Bodymovin Version
	Version *string `json:"v,omitempty"`

	// Composition Width
	Width *float64 `json:"w,omitempty"`
}

// NewFile parses the provided filename and returns a new Animation.
func NewFile(filename string) (*Animation, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(buf)
}

// New returns a new Animation from the provided JSON blob.
func New(buf []byte) (*Animation, error) {
	result := &Animation{}
	if err := json.Unmarshal(buf, result); err != nil {
		return nil, err
	}
	return result, nil
}

// WriteFile writes an Animation as a JSON blob to the named file.
func (a *Animation) WriteFile(filename string) error {
	buf, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf, 0644)
}
