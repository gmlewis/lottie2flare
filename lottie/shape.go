package lottie

// AnchorPoint Shape Transform Anchor Point
type AnchorPoint struct {
}

// Color Fill Color
type Color struct {
}

// Copies Number of Copies
type Copies struct {
}

// Offset Offset of Copies
type Offset struct {
}

// Opacity Fill Opacity
type Opacity struct {
}

// Position Ellipse's position
type Position struct {
}

// Scale Shape Transform Scale
type Scale struct {
}

// Skew Shape Transform Skew
type Skew struct {
}

// SkewAxis Shape Transform Skew Axis
type SkewAxis struct {
}

// Size Ellipse's size
type Size struct {
}

// Start Trim Start.
type Start struct {
}

// EndPoint Gradient End Point
type EndPoint struct {
}

// GradientColors Gradient Colors
type GradientColors struct {
}

// HighlightAngle Highlight Angle. Only if type is Radial
type HighlightAngle struct {
}

// HighlightLength Gradient Highlight Length. Only if type is Radial
type HighlightLength struct {
}

// Radius Rounded Corner Radius
type Radius struct {
}

// RoundedCorners Rect's rounded corners
type RoundedCorners struct {
}

// Rotation Star's rotation.
type Rotation struct {
}

// StarType Star's type. Polygon or Star.
type StarType struct {
}

// StartPoint Gradient Start Point
type StartPoint struct {
}

// StrokeWidth Gradient Stroke Width
type StrokeWidth struct {
}

// Type Gradient Type
type Type struct {
}

// Vertices Shape's vertices
type Vertices struct {
}

// Width Stroke Width
type Width struct {
}

// InnerRadius Star's inner radius. (Star only)
type InnerRadius struct {
}

// InnerRoundness Star's inner roundness. (Star only)
type InnerRoundness struct {
}

// OuterRadius Star's outer radius.
type OuterRadius struct {
}

// OuterRoundness Star's outer roundness.
type OuterRoundness struct {
}

// Points Star's number of points.
type Points struct {
}

// ShapeType represents the type of the shape.
type ShapeType string

const (
	ShapeEllipse  ShapeType = "el"
	ShapeFill     ShapeType = "fl"
	ShapeGFill    ShapeType = "gf"
	ShapeGroup    ShapeType = "gr"
	ShapeGStroke  ShapeType = "gs"
	ShapeMerge    ShapeType = "mm"
	ShapeRect     ShapeType = "rc"
	ShapeRepeater ShapeType = "rp"
	ShapeRound    ShapeType = "rd"
	ShapeShape    ShapeType = "sh"
	ShapeStar     ShapeType = "sr"
	ShapeStroke   ShapeType = "st"
	ShapeTrim     ShapeType = "tm"
)

// Shape represents a lottie shape.
type Shape struct {
	// Common to all shapes:

	// After Effect's Match Name. Used for expressions.
	Mn *string `json:"mn,omitempty"`

	// After Effect's Name. Used for expressions.
	Nm *string `json:"nm,omitempty"`

	// Fill Opacity
	// CONFLICT: O *Opacity `json:"o,omitempty"`

	// Shape content type.
	Ty *ShapeType `json:"ty,omitempty"`

	// Highlight Angle. Only if type is Radial
	// CONFLICT: A *HighlightAngle `json:"a,omitempty"`

	// Gradient or trim End Point
	E *EndPoint `json:"e,omitempty"`

	// Gradient Colors
	G *GradientColors `json:"g,omitempty"`

	// Gradient Highlight Length. Only if type is Radial
	H *HighlightLength `json:"h,omitempty"`

	// Gradient Type
	T *Type `json:"t,omitempty"`

	// After Effect's Direction. Direction how the shape is drawn. Used for trim path for example.
	D *float64 `json:"d,omitempty"`

	// Shape's position
	P *Position `json:"p,omitempty"`

	// Gradient or Stroke Line Cap
	Lc *float64 `json:"lc,omitempty"`

	// Gradient or Stroke Line Join
	Lj *float64 `json:"lj,omitempty"`

	// Gradient or Stroke Miter Limit. Only if Line Join is set to Miter.
	Ml *float64 `json:"ml,omitempty"`

	// Gradient or Stroke Width
	W *StrokeWidth `json:"w,omitempty"`

	// From shapes/ellipse:

	// Ellipse's size
	// CONFLICT: S *Size `json:"s,omitempty"`

	// From shapes/fill:

	// Fill Color
	// CONFLICT: C *Color `json:"c,omitempty"`

	// From shapes/gFill:

	// Gradient Start Point
	// CONFLICT: S *StartPoint `json:"s,omitempty"`

	// From shapes/group:

	// Group list of items
	It []*Shape `json:"it,omitempty"`

	// Group number of properties. Used for expressions.
	Np *float64 `json:"np,omitempty"`

	// From shapes/gStroke:

	// Gradient Start Point
	// CONFLICT: S *StartPoint `json:"s,omitempty"`

	// From shapes/merge:

	// Merge Mode
	Mm *float64 `json:"mm,omitempty"`

	// From shapes/rect:

	// Rect's rounded corners
	// CONFLICT: R *RoundedCorners `json:"r,omitempty"`

	// Rect's size
	// CONFLICT: S *Size `json:"s,omitempty"`

	// From shapes/repeater:

	// Number of Copies
	// CONFLICT: C *Copies `json:"c,omitempty"`

	// Composite of copies
	M *float64 `json:"m,omitempty"`

	// Offset of Copies
	// CONFLICT: O *Offset `json:"o,omitempty"`

	// Transform values for each repeater copy
	Tr *Transform `json:"tr,omitempty"`

	// From shapes/round:

	// Rounded Corner Radius
	// CONFLICT: R *Radius `json:"r,omitempty"`

	// From shapes/shape:

	// Shape's vertices
	Ks *Vertices `json:"ks,omitempty"`

	// From shapes/star:

	// Star's inner radius. (Star only)
	Ir *InnerRadius `json:"ir,omitempty"`

	// Star's inner roundness. (Star only)
	Is *InnerRoundness `json:"is,omitempty"`

	// Star's outer radius.
	Or *OuterRadius `json:"or,omitempty"`

	// Star's outer roundness.
	Os *OuterRoundness `json:"os,omitempty"`

	// Star's number of points.
	Pt *Points `json:"pt,omitempty"`

	// Star's rotation.
	// CONFLICT: R *Rotation `json:"r,omitempty"`

	// Star's type. Polygon or Star.
	Sy *StarType `json:"sy,omitempty"`

	// From shapes/stroke:

	// Stroke Color
	// CONFLICT: C *Color `json:"c,omitempty"`

	// Stroke Opacity
	// CONFLICT: O *Opacity `json:"o,omitempty"`

	// From shapes/transform:

	// Shape Transform Anchor Point
	// CONFLICT: A *AnchorPoint `json:"a,omitempty"`

	// Shape Transform Opacity
	// CONFLICT: O *Opacity `json:"o,omitempty"`

	// Shape Transform Rotation
	// CONFLICT: R *Rotation `json:"r,omitempty"`

	// Shape Transform Scale
	// CONFLICT: S *Scale `json:"s,omitempty"`

	// Shape Transform Skew Axis
	Sa *SkewAxis `json:"sa,omitempty"`

	// Shape Transform Skew
	Sk *Skew `json:"sk,omitempty"`

	// From shapes/trim:

	// Trim End.
	// CONFLICT? E *End `json:"e,omitempty"`

	// Trim Offset.
	// CONFLICT: O *Offset `json:"o,omitempty"`

	// Trim Start.
	// CONFLICT: S *Start `json:"s,omitempty"`
}
