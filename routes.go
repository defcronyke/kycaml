package kycaml

/* Copyright (C) 2021  Jeremy Carter <jeremy@jeremycarter.ca>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/

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
