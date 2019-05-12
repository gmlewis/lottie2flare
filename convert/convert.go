// Package convert converts lottie to flare animations.
package convert

import (
	"log"

	"github.com/gmlewis/lottie2flare/flare"
	f "github.com/gmlewis/lottie2flare/flare"
	"github.com/gmlewis/lottie2flare/lottie"
	ll "github.com/gmlewis/lottie2flare/lottie/layer"
	ls "github.com/gmlewis/lottie2flare/lottie/shape"
)

// Lottie2Flare converts a lottie animation to a flare root.
func Lottie2Flare(anim *lottie.Animation) (*f.Root, error) {
	root := &f.Root{
		Version: ip(21),
		Artboards: []*f.Artboard{
			{
				ClipContents: bp(true),
				Color:        &f.Color{0, 0, 0, 1},
				Height:       fp(anim.GetHeight()),
				Name:         sp("Artboard"),
				Origin:       []float64{0, 0},
				Translation:  []float64{0, 0},
				Type:         sp("artboard"),
				Width:        fp(anim.GetWidth()),
				Animations: []*f.Animation{
					{
						Type:           sp("animation"),
						AnimationStart: fp(anim.GetInPoint() / anim.GetFrameRate()),
						AnimationEnd:   fp(anim.GetOutPoint() / anim.GetFrameRate()),
						FPS:            fp(anim.GetFrameRate()),
						Duration:       fp(anim.GetOutPoint() / anim.GetFrameRate()),
						IsLooping:      bp(true),
						Name:           sp("idle"),
					},
				},
			},
		},
	}

	for _, layer := range anim.GetLayers() {
		processLayer(-1, layer, root.Artboards[0])
	}

	return root, nil
}

func processLayer(parentIndex int, layer lottie.Layer, ab *f.Artboard) {
	switch layer.Type() {
	case ll.ShapeType:
		v := layer.(ll.Shape)
		processLayerShape(parentIndex, v, ab)
	default:
		log.Fatalf("Lottie shape type %v not yet supported.", layer.Type())
	}
}

func processLayerShape(parentIndex int, layer ll.Shape, ab *f.Artboard) {
	n := &f.Node{
		BlendMode:   bmp(f.BlendModeType3),
		Clips:       &[]int{},
		DrawOrder:   ip(layer.GetInd()),
		IsCollapsed: bp(false),
		IsVisible:   bp(true),
		Name:        sp(layer.GetNm()),
		Opacity:     layer.InitialOpacity(),
		Rotation:    layer.InitialRotation(),
		Scale:       layer.InitialScale(),
		Translation: layer.InitialPosition(),
		Type:        ntp(f.NodeTypeShape),
	}
	if parentIndex >= 0 {
		n.Parent = ip(parentIndex)
	}
	nodeNum := len(ab.Nodes)
	ab.Nodes = append(ab.Nodes, n)

	for _, shape := range layer.GetShapes() {
		switch shape.Type() {
		case ls.GroupType:
			v := shape.(ls.Group)
			processShapeGroup(nodeNum, v, layer, ab)
		default:
			log.Fatalf("lottie shape type %q not yet supported", shape.GetTy())
		}
	}

	addKeys(nodeNum, layer, ab)
}

func processShapeGroup(parentIndex int, shape ls.Group, layer lottie.Layer, ab *f.Artboard) {
	items := map[lottie.ShapeType][]*lottie.Shape{}
	for _, item := range shape.It {
		ty := item.GetTy()
		items[ty] = append(items[ty], item)
		if ty != lottie.ShapeFill && ty != lottie.ShapeTransform && ty != lottie.ShapeEllipse {
			log.Printf("WARNING Shape type %q not yet supported. Ignoring.", ty)
		}
	}

	// Process fills
	for _, fill := range items[lottie.ShapeFill] {
		n := &f.Node{
			Color:    getColor(fill.GetC().K),
			FillRule: ip(1),
			Name:     sp(fill.GetNm()),
			Opacity:  fp(getScalar(fill.GetO().K) / 100.0),
			Parent:   ip(parentIndex),
			Type:     ntp(f.NodeTypeColorFill),
		}
		ab.Nodes = append(ab.Nodes, n)
	}

	tr := items[lottie.ShapeTransform]
	// Process ellipses
	for _, ellipse := range items[lottie.ShapeEllipse] {
		s := getVector(ellipse.GetS().K)
		p := getVector(ellipse.GetP().K)
		if len(s) < 2 {
			log.Fatalf("Bad ellipse S: %#v", s)
		}
		n := &f.Node{
			Clips:       &[]int{},
			Height:      fp(s[1]),
			IsCollapsed: bp(false),
			Name:        sp(ellipse.GetNm()),
			Parent:      ip(parentIndex),
			Translation: p,
			Type:        ntp(f.NodeTypeEllipse),
			Width:       fp(s[0]),
		}
		if len(tr) > 0 {
			n.Opacity = fp(getScalar(tr[0].GetO().K) / 100.0)
			n.Rotation = fp(getKey(tr[0].R, "k"))
			sf := getVector(tr[0].GetS().K)
			n.Scale = []float64{sf[0] / 100.0, sf[1] / 100.0}
		}
		ab.Nodes = append(ab.Nodes, n)
	}
}

func getColor(k interface{}) *flare.Color {
	v := getVector(k)
	if len(v) < 4 {
		log.Fatalf("getColor: bad K vector: %#v", v)
	}
	return &flare.Color{v[0], v[1], v[2], v[3]}
}

func addKeys(componentNum int, layer lottie.Layer, ab *f.Artboard) {
	lastAnim := ab.Animations[len(ab.Animations)-1]
	keyed := &f.Key{
		Component: ip(componentNum),
	}

	if o := layer.Ks.GetO(); o != nil {
		switch v := o.K.(type) {
		case []interface{}:
			keyed.Opacity = [][]*f.KeyNode{getKeys(v, *lastAnim.FPS, 1.0/100.0)}
		default:
			log.Fatalf("Unknown layer.Ks.O.K type %T: %v", v, v)
		}
	}

	lastAnim.Keyed = append(lastAnim.Keyed, keyed)
}

func getKeys(ks []interface{}, fps, scaleFactor float64) (result []*f.KeyNode) {
	lastEnd := 0.0
	for _, k := range ks {
		t := getKeyT(k, fps)
		if hasKey(k, "s") && hasKey(k, "e") {
			s := getKeyS(k, scaleFactor)
			lastEnd = getKeyE(k, scaleFactor)
			result = append(result, &f.KeyNode{
				InterpolatorType: itp(f.InterpolatorType1),
				Time:             fp(t),
				Value:            fp(s),
			})
		} else { // final key
			result = append(result, &f.KeyNode{
				InterpolatorType: itp(f.InterpolatorType1),
				Time:             fp(t),
				Value:            fp(lastEnd),
			})
		}
	}
	return result
}

func getKey(k interface{}, key string) float64 {
	switch m := k.(type) {
	case map[string]interface{}:
		v, ok := m[key]
		if !ok {
			log.Fatalf("getKey(%q) missing key", key)
		}
		switch val := v.(type) {
		case []interface{}:
			return val[0].(float64)
		case interface{}:
			return val.(float64)
		default:
			log.Fatalf("Unknown getKey(%q) val type %T: %v", key, val, val)
		}
	default:
		log.Fatalf("Unknown getKey(%q) type %T: %v", key, m, m)
	}
	return 0.0
}

func getKeyVector(k interface{}, key string) []float64 {
	switch m := k.(type) {
	case map[string]interface{}:
		v, ok := m[key]
		if !ok {
			log.Fatalf("getKeyVector(%q) missing key", key)
		}
		switch val := v.(type) {
		case []interface{}:
			var result []float64
			for _, x := range val {
				result = append(result, x.(float64))
			}
			return result
		default:
			log.Fatalf("Unknown getKeyVector(%q) val type %T: %v", key, val, val)
		}
	default:
		log.Fatalf("Unknown getKeyVector(%q) type %T: %v", key, m, m)
	}
	return nil
}

func getKeyE(k interface{}, scaleFactor float64) float64 {
	return getKey(k, "e") * scaleFactor
}

func getKeyS(k interface{}, scaleFactor float64) float64 {
	return getKey(k, "s") * scaleFactor
}

func getKeyT(k interface{}, fps float64) float64 {
	return getKey(k, "t") / fps
}

func getScalar(k interface{}) float64 {
	switch v := k.(type) {
	case []interface{}:
		return v[0].(float64)
	case interface{}:
		return v.(float64)
	default:
		log.Fatalf("getScalar unknown type %T: %v", v, v)
	}
	return 0.0
}

func getVector(k interface{}) (result []float64) {
	switch v := k.(type) {
	case []interface{}:
		for _, f := range v {
			result = append(result, f.(float64))
		}
	default:
		log.Fatalf("getVector unknown type %T: %v", v, v)
	}
	return result
}

func hasKey(k interface{}, key string) bool {
	switch v := k.(type) {
	case map[string]interface{}:
		_, ok := v[key]
		return ok
	default:
		log.Fatalf("Unknown hasKey(%q) type %T: %v", key, v, v)
	}
	return false
}

func bmp(v f.BlendModeType) *f.BlendModeType {
	return &v
}

func bp(v bool) *bool {
	return &v
}

func fp(v float64) *float64 {
	return &v
}

func ip(v int) *int {
	return &v
}

func itp(v f.InterpolatorType) *f.InterpolatorType {
	return &v
}

func ntp(v f.NodeType) *f.NodeType {
	return &v
}

func sp(v string) *string {
	return &v
}
