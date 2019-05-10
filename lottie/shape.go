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

// Scale Shape Transform Scale
type Scale struct {
}

// Skew Shape Transform Skew
type Skew struct {
}

// SkewAxis Shape Transform Skew Axis
type SkewAxis struct {
}

// Start Trim Start.
type Start struct {
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

// ShapeProp represents a lottie properties/shapeProp.
type ShapeProp struct {
	// Closed property of shape
	C *bool `json:"c,omitempty"`

	// Bezier curve In points. Array of 2 dimensional arrays.
	I [][]float64 `json:"i,omitempty"`

	// Bezier curve Out points. Array of 2 dimensional arrays.
	O [][]float64 `json:"o,omitempty"`

	// Bezier curve Vertices. Array of 2 dimensional arrays.
	V [][]float64 `json:"v,omitempty"`
}

// ShapePropOrKeyframe represents a lottie
// properties/shapeProp or
// properties/shapePropKeyframe.
type ShapePropOrKeyframe struct {
	ShapeProp

	// Bezier curve interpolation in value.
	// I *I `json:"i,omitempty"`

	// Bezier curve interpolation out value.
	// O *O `json:"o,omitempty"`

	// Start value of keyframe segment.
	S []*ShapeProp `json:"s,omitempty"`

	// Start time of keyframe segment.
	T float64 `json:"t,omitempty"`
}

// ShapeOrKeyframed represents a lottie
// properties/shape or
// properties/shapeKeyframed.
type ShapeOrKeyframed struct {
	// Defines if property is animated
	A *float64 `json:"a,omitempty"`

	// Property Index. Used for expressions.
	Ix *string `json:"ix,omitempty"`

	// Property Value or keyframes.
	K interface{} `json:"k,omitempty"`

	// Property Expression. An AE expression that modifies the value.
	X *string `json:"x,omitempty"`

	// In Spatial Tangent. Only for spatial properties. Array of numbers.
	Ti []float64 `json:"ti,omitempty"`

	// Out Spatial Tangent. Only for spatial properties. Array of numbers.
	To []float64 `json:"to,omitempty"`
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

	// Closed property of shape
	Closed *bool `json:"closed,omitempty"`

	// Fill Color
	C *MultiDimensionalOrKeyframed `json:"c,omitempty"`

	// After Effect's Match Name. Used for expressions.
	Mn *string `json:"mn,omitempty"`

	// After Effect's Name. Used for expressions.
	Nm *string `json:"nm,omitempty"`

	// Fill Opacity
	O *ValueOrKeyframed `json:"o,omitempty"`

	// Shape content type.
	Ty *ShapeType `json:"ty,omitempty"`

	// Highlight Angle. Only if type is Radial
	A *MultiDimensionalOrKeyframed `json:"a,omitempty"`

	// Gradient or trim End Point
	// CONFLICT: E *EndPoint `json:"e,omitempty"`

	// Gradient Colors
	G *GradientColors `json:"g,omitempty"`

	// Gradient Highlight Length. Only if type is Radial
	H *HighlightLength `json:"h,omitempty"`

	// Layer index in AE. Used for parenting and expressions.
	Ind *float64 `json:"ind,omitempty"`

	// Gradient Type
	T *Type `json:"t,omitempty"`

	// After Effect's Direction. Direction how the shape is drawn. Used for trim path for example.
	D *float64 `json:"d,omitempty"`

	// Shape's position
	P *MultiDimensionalOrKeyframed `json:"p,omitempty"`

	// Gradient or Stroke Line Cap
	Lc *float64 `json:"lc,omitempty"`

	// Gradient or Stroke Line Join
	Lj *float64 `json:"lj,omitempty"`

	// Gradient or Stroke Miter Limit. Only if Line Join is set to Miter.
	Ml *float64 `json:"ml,omitempty"`

	// Gradient or Stroke Width
	W *ValueOrKeyframed `json:"w,omitempty"`

	// Shape's size
	S *MultiDimensionalOrKeyframed `json:"s,omitempty"`

	// From shapes/ellipse:

	// From shapes/fill:

	// FillEnabled is whether the fill is enabled.
	FillEnabled *bool `json:"fillEnabled,omitempty"`

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
	Ks *ShapeOrKeyframed `json:"ks,omitempty"`

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
	R *ValueOrKeyframed `json:"r,omitempty"`

	// Shape Transform Scale
	// CONFLICT: S *Scale `json:"s,omitempty"`

	// Shape Transform Skew Axis
	Sa *SkewAxis `json:"sa,omitempty"`

	// Shape Transform Skew
	Sk *Skew `json:"sk,omitempty"`

	// From shapes/trim:

	// Trim End.
	E *ValueOrKeyframed `json:"e,omitempty"`

	// Trim Offset.
	// CONFLICT: O *Offset `json:"o,omitempty"`

	// Trim Start.
	// CONFLICT: S *Start `json:"s,omitempty"`
}
