package flare

import (
	"encoding/json"
	"io/ioutil"
)

// Animation represents a flare animation.
type Animation struct {
	// AnimationEnd represents the end time of the animation.
	AnimationEnd *float64 `json:"animationEnd,omitempty"`

	// AnimationStart represents the start time of the animation.
	AnimationStart *float64 `json:"animationStart,omitempty"`

	// Duration represents the duration of the animation.
	Duration *float64 `json:"duration,omitempty"`

	// FPS is the frames per second rate of the animation.
	FPS *float64 `json:"fps,omitempty"`

	// IsLooping controls the looping of the animation.
	IsLooping *bool `json:"isLooping,omitempty"`

	// Keyed represents a list of keys for the animation.
	Keyed []*Key `json:"keyed,omitempty"`

	// Name represents the name of the animation.
	Name *string `json:"name,omitempty"`

	// Type is "animation".
	Type *string `json:"type,omitempty"`
}

// Artboard represents a flare artboard.
type Artboard struct {
	// Animations represents a list of flare animations.
	Animations []*Animation `json:"animations,omitempty"`

	// ClipContents controls clipping of contents.
	ClipContents *bool `json:"clipContents,omitempty"`

	// Color represents a flare RGBA color (0-1).
	Color *Color `json:"color,omitempty"`

	// Height represents a height.
	Height *float64 `json:"height,omitempty"`

	// Name represents an artboard name.
	Name *string `json:"name,omitempty"`

	// Nodes represents a list of flare nodes.
	Nodes []*Node `json:"nodes,omitempty"`

	// Origin represents a flare origin.
	Origin []float64 `json:"origin,omitempty"`

	// Translation represents a flare translation.
	Translation []float64 `json:"translation,omitempty"`

	// Type is "artboard".
	Type *string `json:"type,omitempty"`

	// Width represents a width.
	Width *float64 `json:"width,omitempty"`
}

// Color represents a flare RGBA color (0-1).
type Color [4]float64

// BlendModeType controls the blend mode of the node.
type BlendModeType int

const (
	BlendModeType3 BlendModeType = 3 // TODO: figure these out.
)

// Bone represents a flare bone.
type Bone struct {
}

// InterpolatorType controls the type of interpolation used by a key.
type InterpolatorType int

const (
	InterpolatorType1 InterpolatorType = 1 // TODO: figure these out.
)

// Key represents an animation key in flare.
type Key struct {
	// Component is the component that this key affects.
	Component *int `json:"component,omitempty"`

	// Opacity is a 2D array of opacity keys.
	Opacity [][]*KeyNode `json:"opacity,omitempty"`

	// PathVertices is a 2D array of pathVertices keys.
	PathVertices [][]*KeyNode `json:"pathVertices,omitempty"`

	// PosX is a 2D array of position keys.
	PosX [][]*KeyNode `json:"posX,omitempty"`

	// PosY is a 2D array of position keys.
	PosY [][]*KeyNode `json:"posY,omitempty"`

	// Rotation is a 2D array of rotation keys.
	Rotation [][]*KeyNode `json:"rotation,omitempty"`

	// ScaleX is a 2D array of scale keys.
	ScaleX [][]*KeyNode `json:"scaleX,omitempty"`

	// ScaleY is a 2D array of scale keys.
	ScaleY [][]*KeyNode `json:"scaleY,omitempty"`
}

// KeyNode is a single key node.
type KeyNode struct {
	// CubicX1 controls the scale key.
	CubicX1 *float64 `json:"cubicX1,omitempty"`

	// CubicX2 controls the scale key.
	CubicX2 *float64 `json:"cubicX2,omitempty"`

	// CubicY1 controls the scale key.
	CubicY1 *float64 `json:"cubicY1,omitempty"`

	// CubicY2 controls the scale key.
	CubicY2 *float64 `json:"cubicY2,omitempty"`

	// InterpolatorType controls the interpolation type of the key.
	InterpolatorType *InterpolatorType `json:"interpolatorType,omitempty"`

	// Time is the key time.
	Time *float64 `json:"time,omitempty"`

	// Value is the key value or array of values.
	Value interface{} `json:"value,omitempty"`
}

// NodeType represents the node type.
type NodeType string

const (
	// NodeTypeColorFill is a node type.
	NodeTypeColorFill NodeType = "colorFill"
	// NodeTypeColorStroke is a node type.
	NodeTypeColorStroke NodeType = "colorStroke"
	// NodeTypeEllipse is a node type.
	NodeTypeEllipse NodeType = "ellipse"
	// NodeTypeNode is a node type.
	NodeTypeNode NodeType = "node"
	// NodeTypePath is a node type.
	NodeTypePath NodeType = "path"
	// NodeTypeRadialGradientFill is a node type.
	NodeTypeRadialGradientFill NodeType = "radialGradientFill"
	// NodeTypeShape is a node type.
	NodeTypeShape NodeType = "shape"
)

// Node represents a flare node.
type Node struct {
	// BlendMode is a node.
	BlendMode *BlendModeType `json:"blendMode,omitempty"`

	// Bones is a list of bones.
	// Note that the pointer to the slice is needed in this
	// case to preserve the case where an empty list is
	// in the original flare file.
	Bones *[]*Bone `json:"bones,omitempty"`

	// Cap is a cap value.
	Cap *float64 `json:"cap,omitempty"`

	// Clips is a list of clips.
	// Note that the pointer to the slice is needed in this
	// case to preserve the case where an empty list is
	// in the original flare file.
	Clips *[]int `json:"clips,omitempty"`

	// Color is an RGBA color.
	Color *Color `json:"color,omitempty"`

	// ColorStops is a list of colorStops.
	ColorStops []float64 `json:"colorStops,omitempty"`

	// DrawOrder is the node's draw order.
	DrawOrder *int `json:"drawOrder,omitempty"`

	// End is an end value.
	End []float64 `json:"end,omitempty"`

	// FillRule is a fill rule.
	FillRule *int `json:"fillRule,omitempty"`

	// Height is a height value.
	Height *float64 `json:"height,omitempty"`

	// IsClosed controls the closed state of the node.
	IsClosed *bool `json:"isClosed,omitempty"`

	// IsCollapsed controls the collapsed state of the node.
	IsCollapsed *bool `json:"isCollapsed,omitempty"`

	// IsVisible contols the visibility of the node.
	IsVisible *bool `json:"isVisible,omitempty"`

	// Join is a join value.
	Join *float64 `json:"join,omitempty"`

	// Name is the node name.
	Name *string `json:"name,omitempty"`

	// NumColorStops is the number of color stops.
	NumColorStops *int `json:"numColorStops,omitempty"`

	// Opacity is an opacity.
	Opacity *float64 `json:"opacity,omitempty"`

	// Parent is the parent node number.
	Parent *int `json:"parent,omitempty"`

	// Points is a list of points for the node.
	Points []*Point `json:"points,omitempty"`

	// Rotation is a rotation.
	Rotation *float64 `json:"rotation,omitempty"`

	// Scale is a 2D scale.
	Scale []float64 `json:"scale,omitempty"`

	// SecondaryRadiusScale is a secondary radius scale.
	SecondaryRadiusScale *float64 `json:"secondaryRadiusScale,omitempty"`

	// Start is a start value.
	Start []float64 `json:"start,omitempty"`

	// Translation is a 2D translation.
	Translation []float64 `json:"translation,omitempty"`

	// Trim is a trim value.
	Trim *float64 `json:"trim,omitempty"`

	// Type is the node type.
	Type *NodeType `json:"type,omitempty"`

	// Width is a width value.
	Width *float64 `json:"width,omitempty"`
}

// PointType represents the point type.
type PointType int

// Point represents a flare point
type Point struct {
	// In is a list of in values.
	In []float64 `json:"in,omitempty"`

	// Out is a list of out values.
	Out []float64 `json:"out,omitempty"`

	// PointType is the point type.
	PointType *PointType `json:"pointType,omitempty"`

	// Radius is the radius of the point.
	Radius *float64 `json:"radius,omitempty"`

	// Translation is the translation of the point.
	Translation []float64 `json:"translation,omitempty"`
}

// Root represents a flare top-level JSON blob.
type Root struct {
	// Artboards represents a list of flare artboards.
	Artboards []*Artboard `json:"artboards,omitempty"`

	// Version represents a flare version.
	Version *int `json:"version,omitempty"`
}

// NewFile parses the provided filename and returns a new Root.
func NewFile(filename string) (*Root, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(buf)
}

// New returns a new Root from the provided JSON blob.
func New(buf []byte) (*Root, error) {
	result := &Root{}
	if err := json.Unmarshal(buf, result); err != nil {
		return nil, err
	}
	return result, nil
}

// WriteFile writes an Root as a JSON blob to the named file.
func (a *Root) WriteFile(filename string) error {
	buf, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf, 0644)
}
