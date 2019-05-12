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

// GetAo returns the Ao field if it's non-nil, zero value otherwise.
func (l *Layer) GetAo() float64 {
	if l == nil || l.Ao == nil {
		return 0.0
	}
	return *l.Ao
}

// GetBm returns the Bm field if it's non-nil, zero value otherwise.
func (l *Layer) GetBm() float64 {
	if l == nil || l.Bm == nil {
		return 0.0
	}
	return *l.Bm
}

// GetBounds returns the Bounds field.
func (l *Layer) GetBounds() *Bounds {
	if l == nil {
		return nil
	}
	return l.Bounds
}

// GetC returns the C field.
func (l *Layer) GetC() *ValueOrKeyframed {
	if l == nil {
		return nil
	}
	return l.C
}

// GetCl returns the Cl field if it's non-nil, zero value otherwise.
func (l *Layer) GetCl() string {
	if l == nil || l.Cl == nil {
		return ""
	}
	return *l.Cl
}

// GetDdd returns the Ddd field if it's non-nil, zero value otherwise.
func (l *Layer) GetDdd() float64 {
	if l == nil || l.Ddd == nil {
		return 0.0
	}
	return *l.Ddd
}

// GetEn returns the En field if it's non-nil, zero value otherwise.
func (l *Layer) GetEn() int {
	if l == nil || l.En == nil {
		return 0
	}
	return *l.En
}

// GetH returns the H field if it's non-nil, zero value otherwise.
func (l *Layer) GetH() float64 {
	if l == nil || l.H == nil {
		return 0.0
	}
	return *l.H
}

// GetHasMask returns the HasMask field if it's non-nil, zero value otherwise.
func (l *Layer) GetHasMask() bool {
	if l == nil || l.HasMask == nil {
		return false
	}
	return *l.HasMask
}

// GetInd returns the Ind field if it's non-nil, zero value otherwise.
func (l *Layer) GetInd() int {
	if l == nil || l.Ind == nil {
		return 0
	}
	return *l.Ind
}

// GetInPoint returns the InPoint field if it's non-nil, zero value otherwise.
func (l *Layer) GetInPoint() float64 {
	if l == nil || l.InPoint == nil {
		return 0.0
	}
	return *l.InPoint
}

// GetIx returns the Ix field if it's non-nil, zero value otherwise.
func (l *Layer) GetIx() int {
	if l == nil || l.Ix == nil {
		return 0
	}
	return *l.Ix
}

// GetKs returns the Ks field.
func (l *Layer) GetKs() *Transform {
	if l == nil {
		return nil
	}
	return l.Ks
}

// GetLn returns the Ln field if it's non-nil, zero value otherwise.
func (l *Layer) GetLn() string {
	if l == nil || l.Ln == nil {
		return ""
	}
	return *l.Ln
}

// GetMn returns the Mn field if it's non-nil, zero value otherwise.
func (l *Layer) GetMn() string {
	if l == nil || l.Mn == nil {
		return ""
	}
	return *l.Mn
}

// GetNm returns the Nm field if it's non-nil, zero value otherwise.
func (l *Layer) GetNm() string {
	if l == nil || l.Nm == nil {
		return ""
	}
	return *l.Nm
}

// GetOp returns the Op field if it's non-nil, zero value otherwise.
func (l *Layer) GetOp() float64 {
	if l == nil || l.Op == nil {
		return 0.0
	}
	return *l.Op
}

// GetParent returns the Parent field if it's non-nil, zero value otherwise.
func (l *Layer) GetParent() float64 {
	if l == nil || l.Parent == nil {
		return 0.0
	}
	return *l.Parent
}

// GetRefID returns the RefID field if it's non-nil, zero value otherwise.
func (l *Layer) GetRefID() string {
	if l == nil || l.RefID == nil {
		return ""
	}
	return *l.RefID
}

// GetS returns the S field.
func (l *Layer) GetS() *ValueOrKeyframed {
	if l == nil {
		return nil
	}
	return l.S
}

// GetSc returns the Sc field if it's non-nil, zero value otherwise.
func (l *Layer) GetSc() string {
	if l == nil || l.Sc == nil {
		return ""
	}
	return *l.Sc
}

// GetSh returns the Sh field if it's non-nil, zero value otherwise.
func (l *Layer) GetSh() float64 {
	if l == nil || l.Sh == nil {
		return 0.0
	}
	return *l.Sh
}

// GetSr returns the Sr field if it's non-nil, zero value otherwise.
func (l *Layer) GetSr() float64 {
	if l == nil || l.Sr == nil {
		return 0.0
	}
	return *l.Sr
}

// GetSt returns the St field if it's non-nil, zero value otherwise.
func (l *Layer) GetSt() float64 {
	if l == nil || l.St == nil {
		return 0.0
	}
	return *l.St
}

// GetSw returns the Sw field if it's non-nil, zero value otherwise.
func (l *Layer) GetSw() float64 {
	if l == nil || l.Sw == nil {
		return 0.0
	}
	return *l.Sw
}

// GetTd returns the Td field if it's non-nil, zero value otherwise.
func (l *Layer) GetTd() float64 {
	if l == nil || l.Td == nil {
		return 0.0
	}
	return *l.Td
}

// GetTm returns the Tm field.
func (l *Layer) GetTm() *ValueOrKeyframed {
	if l == nil {
		return nil
	}
	return l.Tm
}

// GetTt returns the Tt field if it's non-nil, zero value otherwise.
func (l *Layer) GetTt() float64 {
	if l == nil || l.Tt == nil {
		return 0.0
	}
	return *l.Tt
}

// GetTy returns the Ty field if it's non-nil, zero value otherwise.
func (l *Layer) GetTy() LayerType {
	if l == nil || l.Ty == nil {
		return 0
	}
	return *l.Ty
}

// GetV returns the V field.
func (l *Layer) GetV() *ValueOrKeyframed {
	if l == nil {
		return nil
	}
	return l.V
}

// GetW returns the W field if it's non-nil, zero value otherwise.
func (l *Layer) GetW() float64 {
	if l == nil || l.W == nil {
		return 0.0
	}
	return *l.W
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
func (m *MultiDimensionalOrKeyframed) GetA() float64 {
	if m == nil || m.A == nil {
		return 0.0
	}
	return *m.A
}

// GetIx returns the Ix field if it's non-nil, zero value otherwise.
func (m *MultiDimensionalOrKeyframed) GetIx() int {
	if m == nil || m.Ix == nil {
		return 0
	}
	return *m.Ix
}

// GetX returns the X field if it's non-nil, zero value otherwise.
func (m *MultiDimensionalOrKeyframed) GetX() string {
	if m == nil || m.X == nil {
		return ""
	}
	return *m.X
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

// GetA returns the A field.
func (t *Transform) GetA() *MultiDimensionalOrKeyframed {
	if t == nil {
		return nil
	}
	return t.A
}

// GetO returns the O field.
func (t *Transform) GetO() *ValueOrKeyframed {
	if t == nil {
		return nil
	}
	return t.O
}

// GetP returns the P field.
func (t *Transform) GetP() *MultiDimensionalOrKeyframed {
	if t == nil {
		return nil
	}
	return t.P
}

// GetPx returns the Px field.
func (t *Transform) GetPx() *ValueOrKeyframed {
	if t == nil {
		return nil
	}
	return t.Px
}

// GetPy returns the Py field.
func (t *Transform) GetPy() *ValueOrKeyframed {
	if t == nil {
		return nil
	}
	return t.Py
}

// GetPz returns the Pz field.
func (t *Transform) GetPz() *ValueOrKeyframed {
	if t == nil {
		return nil
	}
	return t.Pz
}

// GetR returns the R field.
func (t *Transform) GetR() *ValueOrKeyframed {
	if t == nil {
		return nil
	}
	return t.R
}

// GetS returns the S field.
func (t *Transform) GetS() *MultiDimensionalOrKeyframed {
	if t == nil {
		return nil
	}
	return t.S
}

// GetSa returns the Sa field.
func (t *Transform) GetSa() *ValueOrKeyframed {
	if t == nil {
		return nil
	}
	return t.Sa
}

// GetSk returns the Sk field.
func (t *Transform) GetSk() *ValueOrKeyframed {
	if t == nil {
		return nil
	}
	return t.Sk
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

// GetA returns the A field if it's non-nil, zero value otherwise.
func (v *ValueOrKeyframed) GetA() float64 {
	if v == nil || v.A == nil {
		return 0.0
	}
	return *v.A
}

// GetIx returns the Ix field if it's non-nil, zero value otherwise.
func (v *ValueOrKeyframed) GetIx() int {
	if v == nil || v.Ix == nil {
		return 0
	}
	return *v.Ix
}

// GetX returns the X field if it's non-nil, zero value otherwise.
func (v *ValueOrKeyframed) GetX() string {
	if v == nil || v.X == nil {
		return ""
	}
	return *v.X
}
