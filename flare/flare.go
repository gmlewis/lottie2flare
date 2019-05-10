package flare

import (
	"encoding/json"
	"io/ioutil"
)

// Animation represents a flare animation JSON blob.
type Animation struct {
}

// NewFile parses the provided filename and returns a new Animation.
func NewFile(filename string) (*Animation, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(buf)
}

// New returns a new Animation from the provided JSON blob.
func New(buf []byte) (*Animation, error) {
	result := &Animation{}
	if err := json.Unmarshal(buf, result); err != nil {
		return nil, err
	}
	return result, nil
}

// WriteFile writes an Animation as a JSON blob to the named file.
func (a *Animation) WriteFile(filename string) error {
	buf, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf, 0644)
}
