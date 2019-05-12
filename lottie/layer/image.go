package layer

// Image represents a lottie image layer.
type Image interface {
	Type() Type
}

// ImageT implements the Image interface.
type ImageT struct {
	Ty Type `json:"ty,omitempty"`
}

// Type returns the layer type.
func (s *ImageT) Type() Type {
	return ImageType
}

// NewImageT returns a ImageT from the JSON map.
func NewImageT(m map[string]interface{}) *ImageT {
	return &ImageT{
		Ty: ImageType,
	}
}
