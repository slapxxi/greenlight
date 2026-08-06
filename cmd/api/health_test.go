package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	app := &application{
		config: config{
			env: "test",
		},
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	app.healthHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code: %d", w.Code)
	}

	var data envelope
	err := json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if data["status"] != "available" {
		t.Errorf("unexpected status: %v", data["status"])
	}

	if data["environment"] != "test" {
		t.Errorf("unexpected environment: %v", data["environment"])
	}

	if data["version"] != version {
		t.Errorf("unexpected version: %v", data["version"])
	}
}

func TestHealthError(t *testing.T) {
	app := &application{}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	app.healthHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code: %d", w.Code)
	}

	var data envelope
	err := json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if data["environment"] != "" {
		t.Errorf("unexpected environment: %v", data["environment"])
	}
}

func TestHealthFormatJSON(t *testing.T) {
	app := &application{
		config: config{
			env: "test",
		},
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	app.healthHandler(w, r)

	b, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	output := string(b)
	data := []byte(`{"environment":"test","status":"available","version":"0.0.1"}`)
	buffer := bytes.NewBuffer(nil)
	json.Indent(buffer, data, "", "  ")
	expected := buffer.String()
	if output != expected+"\n" {
		t.Errorf("unexpected output: %s expected: %s", output, expected)
	}
}
