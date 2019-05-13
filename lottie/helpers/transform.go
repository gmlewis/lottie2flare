//go:generate go run ../gen-accessors.go

// Package helpers provides helper types for lottie files.
package helpers

import (
	"encoding/json"
	"log"

	"github.com/gmlewis/lottie2flare/lottie/properties"
)

type Transform struct {
	// Transform Anchor Point
	AnchorPoint json.RawMessage `json:"a,omitempty"`

	// Transform Opacity
	Opacity json.RawMessage `json:"o,omitempty"`

	// Transform Position
	Position json.RawMessage `json:"p,omitempty"`

	// Transform Position X
	Px json.RawMessage `json:"px,omitempty"`

	// Transform Position Y
	Py json.RawMessage `json:"py,omitempty"`

	// Transform Position Z
	Pz json.RawMessage `json:"pz,omitempty"`

	// Transform Rotation
	Rotation json.RawMessage `json:"r,omitempty"`

	// Transform Scale
	Scale json.RawMessage `json:"s,omitempty"`

	// Transform Skew Axis
	SkewAxis json.RawMessage `json:"sa,omitempty"`

	// Transform Skew
	Skew json.RawMessage `json:"sk,omitempty"`
}

// InitialOpacity returns the initial opacity (0-100) of the transform.
func (t *Transform) InitialOpacity() float64 {
	v, vk := properties.ValueOrValueKeyframed(t.Opacity)
	if v != nil {
		return v.GetValue()
	}
	if vk == nil || len(vk.Values) == 0 || len(vk.Values[0].StartValues) == 0 {
		log.Fatalf("InitialOpacity failed: %s", t.Opacity)
	}
	return vk.Values[0].StartValues[0]
}

// OpacityKeys returns the opacity keys of the transform.
func (t *Transform) OpacityKeys() *properties.ValueKeyframed {
	vk, err := properties.GetValueKeyframed(t.Opacity)
	if err != nil {
		log.Fatalf("OpacityKeys failed: %s", t.Opacity)
	}
	return vk
}

// InitialRotation returns the initial rotation of the transform.
func (t *Transform) InitialRotation() float64 {
	v, vk := properties.ValueOrValueKeyframed(t.Rotation)
	if v != nil {
		return v.GetValue()
	}
	if vk == nil || len(vk.Values) == 0 || len(vk.Values[0].StartValues) == 0 {
		log.Fatalf("InitialRotation failed: %s", t.Rotation)
	}
	return vk.Values[0].StartValues[0]
}

// InitialPosition returns the initial rotation of the transform.
func (t *Transform) InitialPosition() []float64 {
	md, mdk := properties.MultiDimensionalOrMultiDimensionalKeyframed(t.Position)
	if md != nil {
		return md.Value
	}
	if mdk == nil || len(mdk.Keyframes) == 0 {
		log.Fatalf("InitialPosition failed: %s", t.Position)
	}
	return mdk.Keyframes[0].StartValue
}

// InitialScale returns the initial scale of the transform.
func (t *Transform) InitialScale() []float64 {
	md, mdk := properties.MultiDimensionalOrMultiDimensionalKeyframed(t.Scale)
	if md != nil {
		return md.Value
	}
	if mdk == nil || len(mdk.Keyframes) == 0 {
		log.Fatalf("InitialScale failed: %s", t.Scale)
	}
	return mdk.Keyframes[0].StartValue
}
