package shape

import (
	"encoding/json"
	"log"

	"github.com/gmlewis/lottie2flare/lottie/properties"
)

// Ellipse represents a lottie shape/ellipse.
type Ellipse interface {
	// GetName gets the shape's name.
	GetName() string

	// GetPosition gets the shape's position.
	GetPosition() []float64

	// GetSize gets the shape's size.
	GetSize() []float64

	// Type returns EllipseType.
	Type() Type
}

// EllipseT implements the Ellipse interface.
type EllipseT struct {
	// After Effect's Direction. Direction how the shape is drawn. Used for trim path for example.
	Direction *int `json:"d,omitempty"`

	// After Effect's Match Name. Used for expressions.
	MatchName *string `json:"mn,omitempty"`

	// After Effect's Name. Used for expressions.
	Name *string `json:"nm,omitempty"`

	// Ellipse's position
	Position json.RawMessage `json:"p,omitempty"`

	// Ellipse's size
	Size json.RawMessage `json:"s,omitempty"`

	// Shape content type.
	Ty *string `json:"ty,omitempty"`
}

// Type returns the layer type.
func (e *EllipseT) Type() Type {
	return EllipseType
}

// GetPosition returns the initial rotation of the transform.
func (e *EllipseT) GetPosition() []float64 {
	md, mdk := properties.MultiDimensionalOrMultiDimensionalKeyframed(e.Position)
	log.Printf("md=%#v, mdk=%#v", md, mdk)
	return []float64{1, 1}
}

// GetSize returns the initial rotation of the transform.
func (e *EllipseT) GetSize() []float64 {
	md, mdk := properties.MultiDimensionalOrMultiDimensionalKeyframed(e.Size)
	log.Printf("md=%#v, mdk=%#v", md, mdk)
	return []float64{1, 1}
}
