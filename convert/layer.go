// Package convert converts lottie to flare animations.
package convert

import (
	"log"

	f "github.com/gmlewis/lottie2flare/flare"
	"github.com/gmlewis/lottie2flare/lottie"
	ll "github.com/gmlewis/lottie2flare/lottie/layer"
	ls "github.com/gmlewis/lottie2flare/lottie/shape"
)

func processLayer(parentIndex int, layer lottie.Layer, ab *f.Artboard) {
	switch layer.Type() {
	case ll.ShapeType:
		v := layer.(ll.Shape)
		processLayerShape(parentIndex, v, ab)
	default:
		log.Fatalf("Lottie layer type %v not yet supported.", layer.Type())
	}
}

func processLayerShape(parentIndex int, layer ll.Shape, ab *f.Artboard) {
	xfrm := layer.GetTransform()
	scale := xfrm.InitialScale()
	n := &f.Node{
		BlendMode:   bmp(f.BlendModeType3),
		Clips:       &[]int{},
		DrawOrder:   ip(layer.GetIndex()),
		IsCollapsed: bp(false),
		IsVisible:   bp(true),
		Name:        sp(layer.GetName()),
		Opacity:     fp(xfrm.InitialOpacity()),
		Rotation:    fp(xfrm.InitialRotation()),
		Scale:       []float64{scale[0] / 100.0, scale[1] / 100.0},
		Translation: xfrm.InitialPosition()[0:2],
		Type:        ntp(f.NodeTypeShape),
	}
	if parentIndex >= 0 {
		n.Parent = ip(parentIndex)
	}
	nodeNum := len(ab.Nodes)
	ab.Nodes = append(ab.Nodes, n)

	for _, shape := range layer.GetShapes() {
		switch shape.Type() {
		case ls.EllipseType:
			v := shape.(ls.Ellipse)
			processShapeEllipse(nodeNum, v, layer, ab)
		case ls.GroupType:
			v := shape.(ls.Group)
			processShapeGroup(nodeNum, v, layer, ab)
		case ls.StrokeType:
			v := shape.(ls.Stroke)
			processShapeStroke(nodeNum, v, layer, ab)
		case ls.TrimType:
			v := shape.(ls.Trim)
			processShapeTrim(nodeNum, v, layer, ab)
		default:
			log.Fatalf("lottie shape type %q not yet supported", shape.Type())
		}
	}

	addKeys(nodeNum, layer, ab)
}

func addKeys(componentNum int, layer lottie.Layer, ab *f.Artboard) {
	lastAnim := ab.Animations[len(ab.Animations)-1]
	keyed := &f.Key{
		Component: ip(componentNum),
	}

	fps := lastAnim.GetFPS()
	tr := layer.GetTransform()
	if tr == nil {
		return
	}
	if o := tr.OpacityKeys(); o != nil {
		var keys []*f.KeyNode
		var lastEnd float64
		for _, v := range o.Values {
			key := &f.KeyNode{
				InterpolatorType: itp(f.InterpolatorType1),
				Time:             fp(v.GetStartTime() / fps),
			}
			if len(v.StartValues) > 0 {
				key.Value = fp(v.StartValues[0] / 100.0)
			} else {
				key.Value = lastEnd
			}
			keys = append(keys, key)
			if len(v.EndValues) > 0 {
				lastEnd = v.EndValues[0] / 100.0
			}
		}
		keyed.Opacity = [][]*f.KeyNode{keys}
	}

	lastAnim.Keyed = append(lastAnim.Keyed, keyed)
}
