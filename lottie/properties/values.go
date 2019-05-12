package properties

import (
	"encoding/json"
	"log"
)

// Value represents a property value.
type Value struct {
	// Property Index. Used for expressions.
	Index *int `json:"ix,omitempty"`

	// Property Value
	Value *float64 `json:"k,omitempty"`

	// Property Expression. An AE expression that modifies the value.
	Expression *string `json:"x,omitempty"`
}

// ValueKeyframe represents a keyframe value.
type ValueKeyframe struct {
	// Bezier curve interpolation "in" value.
	In *In `json:"i,omitempty"`

	// Start value of keyframe segment.
	S *float64 `json:"s,omitempty"`

	// Start time of keyframe segment.
	T *float64 `json:"t,omitempty"`
}

// ValueKeyframed represents keyframe values.
type ValueKeyframed struct {
	// Property Index. Used for expressions.
	Index *int `json:"ix,omitempty"`

	// Property Value keyframes
	Values []*ValueKeyframe `json:"k,omitempty"`

	// Property Expression. An AE expression that modifies the value.
	Expression *string `json:"x,omitempty"`
}

// ValueOrValueKeyframed returns either a Value or a ValueKeyframed.
func ValueOrValueKeyframed(buf json.RawMessage) (*Value, *ValueKeyframed) {
	v := &Value{}
	if err := json.Unmarshal(buf, v); err != nil {
		log.Fatalf("unmarshal %s: %v", buf, v)
	}
	vk := &ValueKeyframed{}
	if err := json.Unmarshal(buf, vk); err != nil {
		log.Fatalf("unmarshal %s: %v", buf, v)
	}
	return v, vk
}
