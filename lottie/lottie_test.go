package lottie

import (
	"reflect"
	"testing"

	ll "github.com/gmlewis/lottie2flare/lottie/layer"
)

func TestAnimation_GetLayers(t *testing.T) {
	tests := []struct {
		name string
		json string
		want []Layer
	}{
		{
			name: "No layers",
			json: `{"layers":[]}`,
			want: nil,
		},
		{
			name: "One layer.PreCompT",
			json: `{"layers":[{"ty":0}]}`,
			want: []Layer{&ll.PreCompT{Ty: ll.PreCompType}},
		},
		{
			name: "One layer.ImageT",
			json: `{"layers":[{"ty":2}]}`,
			want: []Layer{&ll.ImageT{Ty: ll.ImageType}},
		},
		{
			name: "One layer.NullT",
			json: `{"layers":[{"ty":3}]}`,
			want: []Layer{&ll.NullT{Ty: ll.NullType}},
		},
		{
			name: "One layer.ShapeT",
			json: `{"layers":[{"ty":4}]}`,
			want: []Layer{&ll.ShapeT{Ty: ll.ShapeType}},
		},
		{
			name: "One of each",
			json: `{"layers":[{"ty":0},{"ty":2},{"ty":3},{"ty":4}]}`,
			want: []Layer{
				&ll.PreCompT{Ty: ll.PreCompType},
				&ll.ImageT{Ty: ll.ImageType},
				&ll.NullT{Ty: ll.NullType},
				&ll.ShapeT{Ty: ll.ShapeType},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			anim, err := New([]byte(tt.json))
			if err != nil {
				t.Fatalf("New: %v", err)
			}

			got := anim.GetLayers()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLayers: got %#v, want %#v", got, tt.want)
			}
		})
	}
}
