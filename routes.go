package kycaml

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func USAConsHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	sanctionsCA, err := NewSanctionsCA()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error: %v", err)
		fmt.Fprintf(w, "error: %v\n", err)
		return
	}

	resBytes, err := json.MarshalIndent(sanctionsCA, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error: %v", err)
		fmt.Fprintf(w, "error: %v\n", err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", resBytes)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "200 OK\n")
}
