package shape

import (
	"encoding/json"
	"log"

	"github.com/gmlewis/lottie2flare/lottie/properties"
)

// Fill represents a lottie shape/fill.
type Fill interface {
	// GetColor returns an RGBA color (0-1).
	GetColor() [4]float64

	// GetName returns the name of the fill.
	GetName() string

	// GetOpacity returns the opacity of the fill.
	GetOpacity() float64

	// Type returns FillType.
	Type() Type
}

// FillT implements the Fill interface.
type FillT struct {
	// Fill Color
	Color json.RawMessage `json:"c,omitempty"`

	// After Effect's Match Name. Used for expressions.
	MatchName *string `json:"mn,omitempty"`

	// After Effect's Name. Used for expressions.
	Name *string `json:"nm,omitempty"`

	// Fill Opacity
	Opacity json.RawMessage `json:"o,omitempty"`

	// Shape content type.
	Ty *string `json:"ty,omitempty"`
}

// Type returns the layer type.
func (f *FillT) Type() Type {
	return FillType
}

// GetColor returns an RGBA color (0-1).
func (f *FillT) GetColor() (color [4]float64) {
	md, mdk := properties.MultiDimensionalOrMultiDimensionalKeyframed(f.Color)
	log.Printf("FillT.GetColor: md=%#v, mdk=%#v", md, mdk)
	return color
}

// GetOpacity returns the initial opacity of the transform.
func (f *FillT) GetOpacity() float64 {
	v, vk := properties.ValueOrValueKeyframed(f.Opacity)
	log.Printf("v=%#v, vk=%#v", v, vk)
	return 0.0
}
