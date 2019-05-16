//go:generate go run ../gen-accessors.go

// Package layer represents a lottie layer.
package layer

import (
	"encoding/json"
	"log"

	"github.com/gmlewis/lottie2flare/lottie/helpers"
)

// Type identifies the type of lottie layer.
type Type int

// Layer represents a lottie layer.
type Layer interface {
	// GetTransform returns the shape's transform.
	GetTransform() *helpers.Transform

	// Type returns the layer's type.
	Type() Type
}

// New uses reflection to return a Layer.
func New(buf json.RawMessage) Layer {
	v := &struct{ Ty int }{}
	if err := json.Unmarshal(buf, v); err != nil {
		log.Fatalf("invalid layer: %s", buf)
	}

	var dst Layer
	switch Type(v.Ty) {
	case PreCompType:
		dst = &PreCompT{}
	case ImageType:
		dst = &ImageT{}
	case NullType:
		dst = &NullT{}
	case ShapeType:
		dst = &ShapeT{}
	default:
		log.Fatalf("unsupported layer type %s", buf)
	}

	if err := json.Unmarshal(buf, dst); err != nil {
		log.Fatalf("unmarshal %s: %v", buf, err)
	}
	return dst
}
