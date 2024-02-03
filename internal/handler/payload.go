package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Payload(r *http.Request) (map[string]any, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while executing io.ReadAll: %w", err)
	}

	var p map[string]any

	if err = json.Unmarshal(b, &p); err != nil {
		return nil, fmt.Errorf("error while executing json.Unmarshal: %w", err)
	}

	return p, nil
}
