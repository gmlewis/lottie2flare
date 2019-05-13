package shape

import "testing"

func TestShape_Interface(t *testing.T) {
	// This test will not compile if any of the interfaces are not satisfied.
	var v Shape = &EllipseT{}
	if v == nil {
		t.Errorf("EllipseT does not implement the Shape interface")
	}

	v = &FillT{}
	if v == nil {
		t.Errorf("FillT does not implement the Shape interface")
	}

	v = &GFillT{}
	if v == nil {
		t.Errorf("GFillT does not implement the Shape interface")
	}

	v = &GStrokeT{}
	if v == nil {
		t.Errorf("GStrokeT does not implement the Shape interface")
	}

	v = &GroupT{}
	if v == nil {
		t.Errorf("GroupT does not implement the Shape interface")
	}

	v = &MergeT{}
	if v == nil {
		t.Errorf("MergeT does not implement the Shape interface")
	}

	v = &RectT{}
	if v == nil {
		t.Errorf("RectT does not implement the Shape interface")
	}

	v = &RepeaterT{}
	if v == nil {
		t.Errorf("RepeaterT does not implement the Shape interface")
	}

	v = &RoundT{}
	if v == nil {
		t.Errorf("RoundT does not implement the Shape interface")
	}

	v = &ShapeT{}
	if v == nil {
		t.Errorf("ShapeT does not implement the Shape interface")
	}

	v = &StarT{}
	if v == nil {
		t.Errorf("StarT does not implement the Shape interface")
	}

	v = &StrokeT{}
	if v == nil {
		t.Errorf("StrokeT does not implement the Shape interface")
	}

	v = &TransformT{}
	if v == nil {
		t.Errorf("TransformT does not implement the Shape interface")
	}

	v = &TrimT{}
	if v == nil {
		t.Errorf("TrimT does not implement the Shape interface")
	}
}
