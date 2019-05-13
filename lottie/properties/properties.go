//go:generate go run ../gen-accessors.go

// Package properties represents properties in lottie files.
package properties

// In represents a Bezier curve interpolation "in" value.
type In struct {
	// bezier x axis
	X []float64 `json:"x,omitempty"`

	// bezier y axis
	Y []float64 `json:"y,omitempty"`
}

// Out represents a Bezier curve interpolation "out" value.
type Out struct {

	// bezier x axis. Array of numbers.
	X []float64 `json:"x,omitempty"`

	// bezier y axis. Array of numbers.
	Y []float64 `json:"y,omitempty"`
}
