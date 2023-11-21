package utils

import (
	"encoding/json"
	"io"
)

// Converts an interface to JSON string
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// Converts an interface to Golang struct from JSON string
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
