package kycaml

import (
	"fmt"
	"net/http"
)

/** HTTP Route: GET /sdn */
func (k *KycAml) USASdnJSONHandler(w http.ResponseWriter, r *http.Request) {
	res, err := NewJSONSanctionsSA("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", res)
}

/** HTTP Route: GET /cons */
func (k *KycAml) USAConsJSONHandler(w http.ResponseWriter, r *http.Request) {
	res, err := NewJSONSanctionsCA("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", res)
}

/** HTTP Route: GET /score */
func (k *KycAml) ScoreHandler(w http.ResponseWriter, r *http.Request) {
	res := ""

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", res)
}
