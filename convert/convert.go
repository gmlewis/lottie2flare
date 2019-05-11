// Package convert converts lottie to flare animations.
package convert

import (
	"log"

	f "github.com/gmlewis/lottie2flare/flare"
	"github.com/gmlewis/lottie2flare/lottie"
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
						Name:           sp("go"),
					},
				},
			},
		},
	}

	for _, layer := range anim.Layers {
		processLayer(-1, layer, root.Artboards[0])
	}

	return root, nil
}

func processLayer(parentIndex int, layer *lottie.Layer, ab *f.Artboard) {
	switch layer.GetTy() {
	case lottie.LayerShape:
		processLayerShape(parentIndex, layer, ab)
	default:
		log.Fatalf("Lottie shape type %v not yet supported.", layer.GetTy())
	}
}

func processLayerShape(parentIndex int, layer *lottie.Layer, ab *f.Artboard) {
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

	addKeys(nodeNum, layer, ab)
}

func addKeys(componentNum int, layer *lottie.Layer, ab *f.Artboard) {
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

func getKeyE(k interface{}, scaleFactor float64) float64 {
	return getKey(k, "e") * scaleFactor
}

func getKeyS(k interface{}, scaleFactor float64) float64 {
	return getKey(k, "s") * scaleFactor
}

func getKeyT(k interface{}, fps float64) float64 {
	return getKey(k, "t") / fps
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
