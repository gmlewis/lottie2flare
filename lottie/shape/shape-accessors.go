// Code generated by gen-accessors; DO NOT EDIT.

package shape

// GetMn returns the Mn field if it's non-nil, zero value otherwise.
func (g *GroupT) GetMn() string {
	if g == nil || g.Mn == nil {
		return ""
	}
	return *g.Mn
}

// GetNm returns the Nm field if it's non-nil, zero value otherwise.
func (g *GroupT) GetNm() string {
	if g == nil || g.Nm == nil {
		return ""
	}
	return *g.Nm
}

// GetNp returns the Np field if it's non-nil, zero value otherwise.
func (g *GroupT) GetNp() float64 {
	if g == nil || g.Np == nil {
		return 0.0
	}
	return *g.Np
}

// GetTy returns the Ty field if it's non-nil, zero value otherwise.
func (g *GroupT) GetTy() string {
	if g == nil || g.Ty == nil {
		return ""
	}
	return *g.Ty
}
