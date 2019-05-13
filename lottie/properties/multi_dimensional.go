package properties

import (
	"encoding/json"
)

// MultiDimensional represents a multi-dimensional property value.
type MultiDimensional struct {
	// Property Index. Used for expressions.
	Index *int `json:"ix,omitempty"`

	// Property Value
	Value []float64 `json:"k,omitempty"`

	// Property Expression. An AE expression that modifies the value.
	Expression *string `json:"x,omitempty"`
}

// MultiDimensionalKeyframed represents keyframed multi-dimensional property
// values.
type MultiDimensionalKeyframed struct {
	// Property Index. Used for expressions.
	Index *int `json:"ix,omitempty"`

	// Property Value keyframes
	Keyframes []*OffsetKeyframe `json:"k,omitempty"`

	// In Spatial Tangent. Only for spatial properties. Array of numbers.
	Ti []interface{} `json:"ti,omitempty"`

	// Out Spatial Tangent. Only for spatial properties. Array of numbers.
	To []interface{} `json:"to,omitempty"`

	// Property Expression. An AE expression that modifies the value.
	Expression *string `json:"x,omitempty"`
}

// GetMultiDimensional returns a MultiDimensional if it exists.
func GetMultiDimensional(buf json.RawMessage) (*MultiDimensional, error) {
	v := &MultiDimensional{}
	if err := json.Unmarshal(buf, v); err != nil {
		return nil, err
	}
	return v, nil
}

// GetMultiDimensionalKeyframed returns a MultiDimensionalKeyframed if it exists.
func GetMultiDimensionalKeyframed(buf json.RawMessage) (*MultiDimensionalKeyframed, error) {
	vk := &MultiDimensionalKeyframed{}
	if err := json.Unmarshal(buf, vk); err != nil {
		return nil, err
	}
	return vk, nil
}

// MultiDimensionalOrMultiDimensionalKeyframed returns either a MultiDimensional or a MultiDimensionalKeyframed.
func MultiDimensionalOrMultiDimensionalKeyframed(buf json.RawMessage) (*MultiDimensional, *MultiDimensionalKeyframed) {
	if v, err := GetMultiDimensional(buf); err == nil {
		// Success! It's a MultiDimensional.
		return v, nil
	}
	// Not a MultiDimensional. Try a MultiDimensionalKeyframed.
	vk, err := GetMultiDimensionalKeyframed(buf)
	if err != nil {
		return nil, nil
	}
	return nil, vk
}
