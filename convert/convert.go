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
	ab.Nodes = append(ab.Nodes, n)
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

func ntp(v f.NodeType) *f.NodeType {
	return &v
}

func sp(v string) *string {
	return &v
}
