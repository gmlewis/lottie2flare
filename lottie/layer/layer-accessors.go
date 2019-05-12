// Code generated by gen-accessors; DO NOT EDIT.

package layer

// GetAo returns the Ao field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetAo() int {
	if s == nil || s.Ao == nil {
		return 0
	}
	return *s.Ao
}

// GetBm returns the Bm field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetBm() int {
	if s == nil || s.Bm == nil {
		return 0
	}
	return *s.Bm
}

// GetCl returns the Cl field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetCl() string {
	if s == nil || s.Cl == nil {
		return ""
	}
	return *s.Cl
}

// GetDdd returns the Ddd field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetDdd() int {
	if s == nil || s.Ddd == nil {
		return 0
	}
	return *s.Ddd
}

// GetHasMask returns the HasMask field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetHasMask() bool {
	if s == nil || s.HasMask == nil {
		return false
	}
	return *s.HasMask
}

// GetIndex returns the Index field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetIndex() int {
	if s == nil || s.Index == nil {
		return 0
	}
	return *s.Index
}

// GetInPoint returns the InPoint field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetInPoint() int {
	if s == nil || s.InPoint == nil {
		return 0
	}
	return *s.InPoint
}

// GetLn returns the Ln field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetLn() string {
	if s == nil || s.Ln == nil {
		return ""
	}
	return *s.Ln
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

// GetOutPoint returns the OutPoint field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetOutPoint() int {
	if s == nil || s.OutPoint == nil {
		return 0
	}
	return *s.OutPoint
}

// GetParent returns the Parent field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetParent() int {
	if s == nil || s.Parent == nil {
		return 0
	}
	return *s.Parent
}

// GetSr returns the Sr field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetSr() float64 {
	if s == nil || s.Sr == nil {
		return 0.0
	}
	return *s.Sr
}

// GetSt returns the St field if it's non-nil, zero value otherwise.
func (s *ShapeT) GetSt() float64 {
	if s == nil || s.St == nil {
		return 0.0
	}
	return *s.St
}
