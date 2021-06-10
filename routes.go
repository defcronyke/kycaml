package kycaml

import (
	"fmt"
	"net/http"
)

/** HTTP Route: GET /cons */
func USAConsJSONHandler(w http.ResponseWriter, r *http.Request) {
	res, err := NewJSONSanctionsCA("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", res)
}
