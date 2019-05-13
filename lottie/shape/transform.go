package shape

import (
	"encoding/json"
	"log"

	"github.com/gmlewis/lottie2flare/lottie/properties"
)

// Transform represents a lottie shape/transform.
type Transform interface {
	InitialOpacity() float64

	// InitialRotation returns the initial rotation of the transform.
	InitialRotation() float64

	// InitialPosition returns the initial rotation of the transform.
	InitialPosition() []float64

	// InitialScale returns the initial rotation of the transform.
	InitialScale() []float64

	// Type returns TransformType.
	Type() Type
}

// TransformT implements the Transform interface.
type TransformT struct {
	// Shape Transform Anchor Point
	AnchorPoint json.RawMessage `json:"a,omitempty"`

	// After Effect's Name. Used for expressions.
	Name *string `json:"nm,omitempty"`

	// Shape Transform Opacity
	Opacity json.RawMessage `json:"o,omitempty"`

	// Shape Transform Position
	Position json.RawMessage `json:"p,omitempty"`

	// Shape Transform Rotation
	Rotation json.RawMessage `json:"r,omitempty"`

	// Shape Transform Scale
	Scale json.RawMessage `json:"s,omitempty"`

	// Shape Transform Skew Axis
	SkewAxis json.RawMessage `json:"sa,omitempty"`

	// Shape Transform Skew
	Skew json.RawMessage `json:"sk,omitempty"`
}

// Type returns the layer type.
func (t *TransformT) Type() Type {
	return TransformType
}

// InitialOpacity returns the initial opacity of the transform.
func (t *TransformT) InitialOpacity() float64 {
	v, vk := properties.ValueOrValueKeyframed(t.Opacity)
	log.Printf("shape.Transform: v=%#v, vk=%#v", v, vk)
	return 0.0
}

// InitialRotation returns the initial rotation of the transform.
func (t *TransformT) InitialRotation() float64 {
	v, vk := properties.ValueOrValueKeyframed(t.Rotation)
	log.Printf("shape.Transform: v=%#v, vk=%#v", v, vk)
	return 0.0
}

// InitialPosition returns the initial rotation of the transform.
func (t *TransformT) InitialPosition() []float64 {
	md, mdk := properties.MultiDimensionalOrMultiDimensionalKeyframed(t.Position)
	log.Printf("shape.Transform: md=%#v, mdk=%#v", md, mdk)
	return []float64{1, 1}
}

// InitialScale returns the initial rotation of the transform.
func (t *TransformT) InitialScale() []float64 {
	md, mdk := properties.MultiDimensionalOrMultiDimensionalKeyframed(t.Scale)
	log.Printf("shape.Transform: md=%#v, mdk=%#v", md, mdk)
	return []float64{1, 1}
}
