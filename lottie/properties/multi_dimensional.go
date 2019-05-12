package properties

import (
	"encoding/json"
	"log"
)

// MultiDimensional represents a multi-dimensional property value.
type MultiDimensional struct {
	// Property Index. Used for expressions.
	Index *int `json:"ix,omitempty"`

	// Property Value
	Value []json.RawMessage `json:"k,omitempty"`

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

// MultiDimensionalOrMultiDimensionalKeyframed returns either a
// MultiDimensional or a MultiDimensionalKeyframed.
func MultiDimensionalOrMultiDimensionalKeyframed(buf json.RawMessage) (*MultiDimensional, *MultiDimensionalKeyframed) {
	v := &MultiDimensional{}
	if err := json.Unmarshal(buf, v); err != nil {
		log.Fatalf("unmarshal %s: %v", buf, v)
	}
	vk := &MultiDimensionalKeyframed{}
	if err := json.Unmarshal(buf, vk); err != nil {
		log.Fatalf("unmarshal %s: %v", buf, v)
	}
	return v, vk
}
