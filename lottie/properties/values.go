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
	// Bezier curve interpolation "in" value.
	In *In `json:"i,omitempty"`

	// Start value of keyframe segment.
	StartValue *float64 `json:"s,omitempty"`

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
	if v, err := GetValue(buf); err != nil {
		return v, nil
	}
	vk, err := GetValueKeyframed(buf)
	if err != nil {
		return nil, nil
	}
	return nil, vk
}
