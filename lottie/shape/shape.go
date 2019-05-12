// Package shape represents a lottie shape.
package shape

// Type represents a shape type.
type Type string

const (
	// EllipseT represents a lottie shape/ellipse.
	EllipseT Type = "el"
	// FillT represents a lottie shape/fill.
	FillT Type = "fl"
	// GFillT represents a lottie shape/gFill.
	GFillT Type = "gf"
	// GroupT represents a lottie shape/group.
	GroupT Type = "gr"
	// GStrokeT represents a lottie shape/gStroke.
	GStrokeT Type = "gs"
	// MergeT represents a lottie shape/merge.
	MergeT Type = "mm"
	// RectT represents a lottie shape/rect.
	RectT Type = "rc"
	// RepeaterT represents a lottie shape/repeater.
	RepeaterT Type = "rp"
	// RoundT represents a lottie shape/round.
	RoundT Type = "rd"
	// ShapeT represents a lottie shape/shape.
	ShapeT Type = "sh"
	// StarT represents a lottie shape/star.
	StarT Type = "sr"
	// StrokeT represents a lottie shape/stroke.
	StrokeT Type = "st"
	// TrimT represents a lottie shape/trim.
	TrimT Type = "tm"
	// TransformT represents a lottie transform.
	TransformT Type = "tr"
)

// Shape represents a lottie shape.
type Shape interface {
	Type() Type
}
