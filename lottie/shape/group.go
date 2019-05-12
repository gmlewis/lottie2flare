package shape

// Group represents a lottie shape/group.
type Group interface {
	Type() Type
}

// GroupT implements the Group interface.
type GroupT struct {
}

// Type returns the layer type.
func (s *GroupT) Type() Type {
	return GroupType
}
