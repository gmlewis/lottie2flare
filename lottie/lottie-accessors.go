// Code generated by gen-accessors; DO NOT EDIT.

package lottie

// GetDdd returns the Ddd field if it's non-nil, zero value otherwise.
func (a *Animation) GetDdd() int {
	if a == nil || a.Ddd == nil {
		return 0
	}
	return *a.Ddd
}

// GetFrameRate returns the FrameRate field if it's non-nil, zero value otherwise.
func (a *Animation) GetFrameRate() float64 {
	if a == nil || a.FrameRate == nil {
		return 0.0
	}
	return *a.FrameRate
}

// GetHeight returns the Height field if it's non-nil, zero value otherwise.
func (a *Animation) GetHeight() float64 {
	if a == nil || a.Height == nil {
		return 0.0
	}
	return *a.Height
}

// GetInPoint returns the InPoint field if it's non-nil, zero value otherwise.
func (a *Animation) GetInPoint() float64 {
	if a == nil || a.InPoint == nil {
		return 0.0
	}
	return *a.InPoint
}

// GetMarkers returns the Markers field if it's non-nil, zero value otherwise.
func (a *Animation) GetMarkers() []int {
	if a == nil || a.Markers == nil {
		return nil
	}
	return *a.Markers
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (a *Animation) GetName() string {
	if a == nil || a.Name == nil {
		return ""
	}
	return *a.Name
}

// GetOutPoint returns the OutPoint field if it's non-nil, zero value otherwise.
func (a *Animation) GetOutPoint() float64 {
	if a == nil || a.OutPoint == nil {
		return 0.0
	}
	return *a.OutPoint
}

// GetVersion returns the Version field if it's non-nil, zero value otherwise.
func (a *Animation) GetVersion() string {
	if a == nil || a.Version == nil {
		return ""
	}
	return *a.Version
}

// GetWidth returns the Width field if it's non-nil, zero value otherwise.
func (a *Animation) GetWidth() float64 {
	if a == nil || a.Width == nil {
		return 0.0
	}
	return *a.Width
}

// GetHeight returns the Height field if it's non-nil, zero value otherwise.
func (a *Asset) GetHeight() float64 {
	if a == nil || a.Height == nil {
		return 0.0
	}
	return *a.Height
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (a *Asset) GetID() string {
	if a == nil || a.ID == nil {
		return ""
	}
	return *a.ID
}

// GetP returns the P field if it's non-nil, zero value otherwise.
func (a *Asset) GetP() string {
	if a == nil || a.P == nil {
		return ""
	}
	return *a.P
}

// GetURL returns the URL field if it's non-nil, zero value otherwise.
func (a *Asset) GetURL() string {
	if a == nil || a.URL == nil {
		return ""
	}
	return *a.URL
}

// GetWidth returns the Width field if it's non-nil, zero value otherwise.
func (a *Asset) GetWidth() float64 {
	if a == nil || a.Width == nil {
		return 0.0
	}
	return *a.Width
}

// GetCl returns the Cl field if it's non-nil, zero value otherwise.
func (m *Mask) GetCl() bool {
	if m == nil || m.Cl == nil {
		return false
	}
	return *m.Cl
}

// GetInv returns the Inv field if it's non-nil, zero value otherwise.
func (m *Mask) GetInv() bool {
	if m == nil || m.Inv == nil {
		return false
	}
	return *m.Inv
}

// GetMode returns the Mode field if it's non-nil, zero value otherwise.
func (m *Mask) GetMode() string {
	if m == nil || m.Mode == nil {
		return ""
	}
	return *m.Mode
}

// GetNm returns the Nm field if it's non-nil, zero value otherwise.
func (m *Mask) GetNm() string {
	if m == nil || m.Nm == nil {
		return ""
	}
	return *m.Nm
}

// GetO returns the O field.
func (m *Mask) GetO() *ValueOrKeyframed {
	if m == nil {
		return nil
	}
	return m.O
}

// GetPt returns the Pt field.
func (m *Mask) GetPt() *ShapeOrKeyframed {
	if m == nil {
		return nil
	}
	return m.Pt
}

// GetX returns the X field.
func (m *Mask) GetX() *ValueOrKeyframed {
	if m == nil {
		return nil
	}
	return m.X
}

// GetA returns the A field if it's non-nil, zero value otherwise.
func (s *ShapeOrKeyframed) GetA() float64 {
	if s == nil || s.A == nil {
		return 0.0
	}
	return *s.A
}

// GetIx returns the Ix field if it's non-nil, zero value otherwise.
func (s *ShapeOrKeyframed) GetIx() int {
	if s == nil || s.Ix == nil {
		return 0
	}
	return *s.Ix
}

// GetX returns the X field if it's non-nil, zero value otherwise.
func (s *ShapeOrKeyframed) GetX() string {
	if s == nil || s.X == nil {
		return ""
	}
	return *s.X
}

// GetIx returns the Ix field if it's non-nil, zero value otherwise.
func (v *Value) GetIx() string {
	if v == nil || v.Ix == nil {
		return ""
	}
	return *v.Ix
}

// GetK returns the K field if it's non-nil, zero value otherwise.
func (v *Value) GetK() float64 {
	if v == nil || v.K == nil {
		return 0.0
	}
	return *v.K
}

// GetX returns the X field if it's non-nil, zero value otherwise.
func (v *Value) GetX() string {
	if v == nil || v.X == nil {
		return ""
	}
	return *v.X
}

// GetI returns the I field.
func (v *ValueKeyframe) GetI() *I {
	if v == nil {
		return nil
	}
	return v.I
}

// GetIx returns the Ix field if it's non-nil, zero value otherwise.
func (v *ValueKeyframed) GetIx() string {
	if v == nil || v.Ix == nil {
		return ""
	}
	return *v.Ix
}

// GetX returns the X field if it's non-nil, zero value otherwise.
func (v *ValueKeyframed) GetX() string {
	if v == nil || v.X == nil {
		return ""
	}
	return *v.X
}
