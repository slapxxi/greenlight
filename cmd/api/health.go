package main

import (
	"fmt"
	"net/http"
)

func (a *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s", a.config.env)
	fmt.Fprintf(w, "version: %s", version)
}
