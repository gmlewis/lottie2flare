package shape

// Merge represents a lottie shape/merge.
type Merge interface {
	Type() Type
}

// MergeT implements the Merge interface.
type MergeT struct {
}

// Type returns the layer type.
func (s *MergeT) Type() Type {
	return MergeType
}
