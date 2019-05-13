//go:generate go run ../gen-accessors.go

// Package shape represents a lottie shape.
package shape

import (
	"encoding/json"
	"log"
)

// Type represents a shape type.
type Type string

const (
	// EllipseType represents a lottie shape/ellipse.
	EllipseType Type = "el"
	// FillType represents a lottie shape/fill.
	FillType Type = "fl"
	// GFillType represents a lottie shape/gFill.
	GFillType Type = "gf"
	// GroupType represents a lottie shape/group.
	GroupType Type = "gr"
	// GStrokeType represents a lottie shape/gStroke.
	GStrokeType Type = "gs"
	// MergeType represents a lottie shape/merge.
	MergeType Type = "mm"
	// RectType represents a lottie shape/rect.
	RectType Type = "rc"
	// RepeaterType represents a lottie shape/repeater.
	RepeaterType Type = "rp"
	// RoundType represents a lottie shape/round.
	RoundType Type = "rd"
	// ShapeType represents a lottie shape/shape.
	ShapeType Type = "sh"
	// StarType represents a lottie shape/star.
	StarType Type = "sr"
	// StrokeType represents a lottie shape/stroke.
	StrokeType Type = "st"
	// TrimType represents a lottie shape/trim.
	TrimType Type = "tm"
	// TransformType represents a lottie transform.
	TransformType Type = "tr"
)

// Shape represents a lottie shape/shape.
type Shape interface {
	Type() Type
}

// ShapeT implements the Shape interface.
type ShapeT struct {
}

// Type returns the layer type.
func (s *ShapeT) Type() Type {
	return ShapeType
}

// New uses reflection to return a Shape.
func New(buf json.RawMessage) Shape {
	v := &struct{ Ty string }{}
	if err := json.Unmarshal(buf, v); err != nil {
		log.Fatalf("invalid shape: %s", buf)
	}

	var dst Shape
	switch Type(v.Ty) {
	case EllipseType:
		dst = &EllipseT{}
	case FillType:
		dst = &FillT{}
	case GFillType:
		dst = &GFillT{}
	case GroupType:
		dst = &GroupT{}
	case GStrokeType:
		dst = &GStrokeT{}
	case MergeType:
		dst = &MergeT{}
	case RectType:
		dst = &RectT{}
	case RepeaterType:
		dst = &RepeaterT{}
	case RoundType:
		dst = &RoundT{}
	case ShapeType:
		dst = &ShapeT{}
	case StarType:
		dst = &StarT{}
	case StrokeType:
		dst = &StrokeT{}
	case TrimType:
		dst = &TrimT{}
	case TransformType:
		dst = &TransformT{}
	default:
		log.Fatalf("unsupported shape type %s", buf)
	}

	if err := json.Unmarshal(buf, dst); err != nil {
		log.Fatalf("unmarshal %s: %v", buf, err)
	}
	return dst
}
