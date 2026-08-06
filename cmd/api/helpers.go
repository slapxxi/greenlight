package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

func (a *application) readIdParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, fmt.Errorf("invalid id: %w", err)
	}
	return id, nil
}

func (a *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	output = append(output, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(output)

	return nil
}
