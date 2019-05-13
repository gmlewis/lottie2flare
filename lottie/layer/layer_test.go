package layer

import "testing"

func TestLayer_Interfaces(t *testing.T) {
	// This test will not compile if any of the interfaces are not satisfied.
	var v Layer = &ImageT{}
	if v == nil {
		t.Errorf("ImageT does not implement the Layer interface")
	}

	v = &NullT{}
	if v == nil {
		t.Errorf("NullT does not implement the Layer interface")
	}

	v = &PreCompT{}
	if v == nil {
		t.Errorf("PreCompT does not implement the Layer interface")
	}

	v = &ShapeT{}
	if v == nil {
		t.Errorf("ShapeT does not implement the Layer interface")
	}
}
