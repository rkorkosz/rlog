package rlog

import (
	"encoding/json"
	"io"
)

func RenderJSON(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
