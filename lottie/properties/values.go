package properties

import (
	"encoding/json"
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
	// End value of keyframe segment.
	EndValues []float64 `json:"e,omitempty"`

	// Bezier curve interpolation "in" value.
	In *In `json:"i,omitempty"`

	// Names of keyframe.
	Name []string `json:"n,omitempty"`

	// Bezier curve interpolation "out" value.
	Out *Out `json:"o,omitempty"`

	// Start values of keyframe segment.
	StartValues []float64 `json:"s,omitempty"`

	// Start time of keyframe segment.
	StartTime *float64 `json:"t,omitempty"`
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

// GetValue returns a Value if it exists.
func GetValue(buf json.RawMessage) (*Value, error) {
	v := &Value{}
	if err := json.Unmarshal(buf, v); err != nil {
		return nil, err
	}
	return v, nil
}

// GetValueKeyframed returns a ValueKeyframed if it exists.
func GetValueKeyframed(buf json.RawMessage) (*ValueKeyframed, error) {
	vk := &ValueKeyframed{}
	if err := json.Unmarshal(buf, vk); err != nil {
		return nil, err
	}
	return vk, nil
}

// ValueOrValueKeyframed returns either a Value or a ValueKeyframed.
func ValueOrValueKeyframed(buf json.RawMessage) (*Value, *ValueKeyframed) {
	if v, err := GetValue(buf); err == nil {
		// Success! It's a value.
		return v, nil
	}
	// Not a Value. Try a ValueKeyframed.
	vk, err := GetValueKeyframed(buf)
	if err != nil {
		return nil, nil
	}
	return nil, vk
}
