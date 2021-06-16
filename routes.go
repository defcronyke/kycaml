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
	"os"
)

/** HTTP Route: GET /sdn
Get the USA's Specially Designated Nationals
And Blocked Persons list data, in a custom
unofficial JSON format.
Built from the official XML list found here:
https://home.treasury.gov/policy-issues/financial-sanctions/specially-designated-nationals-list-data-formats-data-schemas
-> SDN_ADVANCED.XML
(USA Sanctions-based) */
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

/** HTTP Route: GET /cons
Get the USA's Consolidated Advanced Sanctions
list data, in a custom unofficial JSON format.
Built from the official XML list found here:
https://home.treasury.gov/policy-issues/financial-sanctions/consolidated-sanctions-list-non-sdn-lists
-> CONS_ADVANCED.XML
(USA Sanctions-based) */
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

/** HTTP Route: GET /names
Get the names of sanctioned entities.
(USA Sanctions-based) */
func (k *KycAml) NamesHandler(w http.ResponseWriter, r *http.Request) {
	res, err := GetNamesJSON(
		fmt.Sprintf("%v/static/%v", os.Getenv("PWD"), "sdn.xml"),
		fmt.Sprintf("%v/static/%v", os.Getenv("PWD"), "cons.xml"),
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", res)
}

/** HTTP Route: GET /names-dm
Get the names of sanctioned entities, with a
double-metaphone encoded lookup key.
(USA Sanctions-based) */
func (k *KycAml) NamesDMHandler(w http.ResponseWriter, r *http.Request) {
	res, err := GetNamesDMJSON(
		fmt.Sprintf("%v/static/%v", os.Getenv("PWD"), "sdn.xml"),
		fmt.Sprintf("%v/static/%v", os.Getenv("PWD"), "cons.xml"),
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", res)
}

/** HTTP Route: GET /names-db
Get the names of sanctioned entities, with a
combination of all supported encodings of lookup
keys. (USA Sanctions-based) */
func (k *KycAml) NamesDBHandler(w http.ResponseWriter, r *http.Request) {
	res, err := GetNamesDBJSON(
		fmt.Sprintf("%v/static/%v", os.Getenv("PWD"), "sdn.xml"),
		fmt.Sprintf("%v/static/%v", os.Getenv("PWD"), "cons.xml"),
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", res)
}
