// Package convert converts lottie to flare animations.
package convert

import (
	"github.com/gmlewis/lottie2flare/flare"
	"github.com/gmlewis/lottie2flare/lottie"
)

// Lottie2Flare converts a lottie animation to a flare root.
func Lottie2Flare(anim *lottie.Animation) (*flare.Root, error) {
	root := &flare.Root{}
	return root, nil
}
