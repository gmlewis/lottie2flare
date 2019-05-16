package convert

import (
	"log"

	f "github.com/gmlewis/lottie2flare/flare"
	"github.com/gmlewis/lottie2flare/lottie"
	ls "github.com/gmlewis/lottie2flare/lottie/shape"
	goon "github.com/shurcooL/go-goon"
)

func processShapeEllipse(parentIndex int, ellipse ls.Ellipse, layer lottie.Layer, ab *f.Artboard) {
	size := ellipse.GetSize()
	n := &f.Node{
		Clips:       &[]int{},
		Height:      fp(size[1]),
		IsCollapsed: bp(false),
		Name:        sp(ellipse.GetName()),
		Parent:      ip(parentIndex),
		Translation: ellipse.GetPosition()[0:2],
		Type:        ntp(f.NodeTypeEllipse),
		Width:       fp(size[0]),
	}
	ab.Nodes = append(ab.Nodes, n)
}

func processShapeGroup(parentIndex int, shape ls.Group, layer lottie.Layer, ab *f.Artboard) {
	items := map[ls.Type][]ls.Shape{}
	for _, item := range shape.GetShapes() {
		ty := item.Type()
		items[ty] = append(items[ty], item)
		if ty != ls.FillType && ty != ls.TransformType && ty != ls.EllipseType {
			log.Printf("WARNING Shape type %q not yet supported. Ignoring.", ty)
		}
	}

	// Process fills
	for _, item := range items[ls.FillType] {
		fill := item.(ls.Fill)
		n := &f.Node{
			Color:    cp(fill.GetColor()),
			FillRule: ip(1),
			Name:     sp(fill.GetName()),
			Opacity:  fp(fill.GetOpacity() / 100.0),
			Parent:   ip(parentIndex),
			Type:     ntp(f.NodeTypeColorFill),
		}
		ab.Nodes = append(ab.Nodes, n)
	}

	tr := items[ls.TransformType]
	// Process ellipses
	for _, item := range items[ls.EllipseType] {
		ellipse := item.(ls.Ellipse)
		size := ellipse.GetSize()
		n := &f.Node{
			Clips:       &[]int{},
			Height:      fp(size[1]),
			IsCollapsed: bp(false),
			Name:        sp(ellipse.GetName()),
			Parent:      ip(parentIndex),
			Translation: ellipse.GetPosition()[0:2],
			Type:        ntp(f.NodeTypeEllipse),
			Width:       fp(size[0]),
		}
		if len(tr) > 0 {
			xfrm := tr[0].(ls.Transform)
			n.Opacity = fp(xfrm.InitialOpacity() / 100.0)
			n.Rotation = fp(xfrm.InitialRotation())
			scale := xfrm.InitialScale()
			n.Scale = []float64{scale[0] / 100.0, scale[1] / 100.0}
		}
		ab.Nodes = append(ab.Nodes, n)
	}
}

func processShapeStroke(parentIndex int, stroke ls.Stroke, layer lottie.Layer, ab *f.Artboard) {
	goon.Dump(stroke)
}

func processShapeTrim(parentIndex int, trim ls.Trim, layer lottie.Layer, ab *f.Artboard) {
	goon.Dump(trim)
}
