//go:generate go run ../gen-accessors.go

// Package layer represents a lottie layer.
package layer

import "log"

// Type identifies the type of lottie layer.
type Type int

const (
	// PreCompType is a lottie preComp layer.
	PreCompType Type = 0
	// ImageType is a lottie image layer.
	ImageType Type = 2
	// NullType is a lottie null layer.
	NullType Type = 3
	// ShapeType is a lottie shape layer.
	ShapeType Type = 4
)

// Layer represents a lottie layer.
type Layer interface {
	Type() Type
}

// New uses reflection to return a Layer.
func New(v interface{}) Layer {
	m, ok := v.(map[string]interface{})
	if !ok || m["ty"] == nil {
		log.Fatalf("invalid layer: %#v", v)
	}
	t := Type(int(m["ty"].(float64)))
	switch t {
	case PreCompType:
		return NewPreCompT(m)
	case ImageType:
		return NewImageT(m)
	case NullType:
		return NewNullT(m)
	case ShapeType:
		return NewShapeT(m)
	default:
		log.Fatalf("unsupported layer type %#v", m)
	}
	return nil
}
