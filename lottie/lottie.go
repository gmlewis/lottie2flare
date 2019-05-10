package lottie

import (
	"encoding/json"
	"io/ioutil"
)

// Asset
type Asset struct {
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
	Layers []*Layer `json:"layers,omitempty"`
}

// CharsItems
type CharsItems struct {
}

type EffectsItems struct {
}

type ItemsItems struct {
}

type Mask struct {
	// Cl ?
	Cl *bool `json:"cl,omitempty"`

	// Inverted Mask flag
	Inv *bool `json:"inv,omitempty"`

	// Mask mode. Not all mask types are supported.
	Mode *string `json:"mode,omitempty"`

	// Mask name. Used for expressions and effects.
	Nm *string `json:"nm,omitempty"`

	// Mask opacity.
	O *ValueOrKeyframed `json:"o,omitempty"`

	// Mask vertices
	Pt *ShapeOrKeyframed `json:"pt,omitempty"`

	// X
	X *ValueOrKeyframed `json:"x,omitempty"`
}

type TextDataItems struct {
}

// I Bezier curve interpolation in value.
type I struct {

	// bezier x axis
	X []float64 `json:"x,omitempty"`

	// bezier y axis
	Y []float64 `json:"y,omitempty"`
}

// Value
type Value struct {

	// Property Index. Used for expressions.
	Ix *string `json:"ix,omitempty"`

	// Property Value
	K *float64 `json:"k,omitempty"`

	// Property Expression. An AE expression that modifies the value.
	X *string `json:"x,omitempty"`
}

// ValueKeyframe
type ValueKeyframe struct {

	// Bezier curve interpolation in value.
	I *I `json:"i,omitempty"`

	// Start value of keyframe segment.
	S []float64 `json:"s,omitempty"`

	// Start time of keyframe segment.
	T float64 `json:"t,omitempty"`
}

// ValueKeyframed
type ValueKeyframed struct {

	// Property Index. Used for expressions.
	Ix *string `json:"ix,omitempty"`

	// Property Value keyframes
	K []*ValueKeyframe `json:"k,omitempty"`

	// Property Expression. An AE expression that modifies the value.
	X *string `json:"x,omitempty"`
}

// Animation represents a Lottie animation JSON blob.
type Animation struct {
	// source items that can be used in multiple places. Comps and Images for now.
	Assets []*Asset `json:"assets"`

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
	Layers []*Layer `json:"layers,omitempty"`

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
