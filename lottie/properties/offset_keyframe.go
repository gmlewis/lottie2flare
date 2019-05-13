package properties

// OffsetKeyframe represents a lottie offset keyframe.
type OffsetKeyframe struct {

	// Bezier curve interpolation in value.
	In *In `json:"i,omitempty"`

	// Bezier curve interpolation out value.
	Out *Out `json:"o,omitempty"`

	// Start value of keyframe segment.
	StartValue []float64 `json:"s,omitempty"`

	// Start time of keyframe segment.
	StartTime *float64 `json:"t,omitempty"`
}
