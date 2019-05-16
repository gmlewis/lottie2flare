// Package convert converts lottie to flare animations.
package convert

import (
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

/*
func getColor(k interface{}) *flare.Color {
	v := getVector(k)
	if len(v) < 4 {
		log.Fatalf("getColor: bad K vector: %#v", v)
	}
	return &flare.Color{v[0], v[1], v[2], v[3]}
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
*/

func bmp(v f.BlendModeType) *f.BlendModeType {
	return &v
}

func bp(v bool) *bool {
	return &v
}

func cp(v [4]float64) *f.Color {
	return &f.Color{v[0], v[1], v[2], v[3]}
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
