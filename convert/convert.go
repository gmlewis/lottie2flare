// Package convert converts lottie to flare animations.
package convert

import (
	f "github.com/gmlewis/lottie2flare/flare"
	"github.com/gmlewis/lottie2flare/lottie"
)

// Lottie2Flare converts a lottie animation to a flare root.
func Lottie2Flare(anim *lottie.Animation) (*f.Root, error) {
	root := &f.Root{
		Artboards: []*f.Artboard{
			{
				Animations: []*f.Animation{
					{
						AnimationEnd: fp(0.4),
					},
				},
			},
		},
	}
	return root, nil
}

func fp(v float64) *float64 {
	return &v
}
