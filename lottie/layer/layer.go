//go:generate go run ../gen-accessors.go

// Package layer represents a lottie layer.
package layer

// Type identifies the type of lottie layer.
type Type int

const (
	// PreCompT is a lottie preComp layer.
	PreCompT Type = 0
	// ImageT is a lottie image layer.
	ImageT Type = 2
	// NullT is a lottie null layer.
	NullT Type = 3
	// ShapeT is a lottie shape layer.
	ShapeT Type = 4
)

// Layer represents a lottie layer.
type Layer interface {
	Type() Type
}
