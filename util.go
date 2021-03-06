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
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/defcronyke/kycaml/model/cons"
	"github.com/defcronyke/kycaml/model/sdn"
)

func SetHeader(header, value string, handle http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set(header, value)
		handle.ServeHTTP(w, req)
	}
}

func NewFile(path string) ([]byte, error) {
	newFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer newFile.Close()

	byteValue, err := ioutil.ReadAll(newFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return byteValue, nil
}

/** SDN List */
func NewSanctionsSA(path string) (*sdn.Sanctions, error) {
	pathDefault := "./static/sdn.xml"

	if path == "" {
		path = pathDefault
	}

	newFile, err := NewFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var sanctions sdn.Sanctions

	err = xml.Unmarshal(newFile, &sanctions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &sanctions, nil
}

func NewJSONSanctionsSA(path string) ([]byte, error) {
	sanctions, err := NewSanctionsSA(path)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	resBytes, err := json.MarshalIndent(sanctions, "", "  ")
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return resBytes, nil
}

/** Cons List */
func NewSanctionsCA(path string) (*cons.Sanctions, error) {
	pathDefault := "./static/cons.xml"

	if path == "" {
		path = pathDefault
	}

	newFile, err := NewFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var sanctions cons.Sanctions

	err = xml.Unmarshal(newFile, &sanctions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &sanctions, nil
}

func NewJSONSanctionsCA(path string) ([]byte, error) {
	sanctions, err := NewSanctionsCA(path)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	resBytes, err := json.MarshalIndent(sanctions, "", "  ")
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return resBytes, nil
}

func GetNamesSA(s *sdn.Sanctions) [][]string {
	names := [][]string{}

	for _, party := range s.DistinctParties.DistinctParty {
		for _, profile := range party.Profile {
			for _, identity := range profile.Identity {
				for _, alias := range identity.Alias {
					for _, documentedName := range alias.DocumentedName {
						nameParts := []string{
							documentedName.FixedRef,
						}

						for _, documentedNamePart := range documentedName.DocumentedNamePart {
							for _, namePartValue := range documentedNamePart.NamePartValue {
								nameParts = append(nameParts, namePartValue.Value)
							}
						}

						names = append(names, nameParts)
					}
				}
			}
		}
	}

	return names
}

func GetNamesSAJSON(s *sdn.Sanctions) ([]byte, error) {
	names := GetNamesSA(s)

	res, err := json.MarshalIndent(names, "", "  ")
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return res, err
}

func GetNamesCA(s *cons.Sanctions) [][]string {
	names := [][]string{}

	for _, party := range s.DistinctParties.DistinctParty {
		for _, profile := range party.Profile {
			for _, identity := range profile.Identity {
				for _, alias := range identity.Alias {
					for _, documentedName := range alias.DocumentedName {
						nameParts := []string{
							documentedName.FixedRef,
						}

						for _, documentedNamePart := range documentedName.DocumentedNamePart {
							for _, namePartValue := range documentedNamePart.NamePartValue {
								nameParts = append(nameParts, namePartValue.Value)
							}
						}

						names = append(names, nameParts)
					}
				}
			}
		}
	}

	return names
}

func GetNamesCAJSON(s *cons.Sanctions) ([]byte, error) {
	names := GetNamesCA(s)

	res, err := json.MarshalIndent(names, "", "  ")
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return res, err
}

func GetNames(pathSA, pathCA string) ([][]string, error) {
	sanctionsSA, err := NewSanctionsSA(pathSA)
	if err != nil {
		return nil, err
	}

	sanctionsCA, err := NewSanctionsCA(pathCA)
	if err != nil {
		return nil, err
	}

	names := append(
		GetNamesSA(sanctionsSA),
		GetNamesCA(sanctionsCA)...,
	)

	return names, nil
}

func GetNamesJSON(pathSA, pathCA string) ([]byte, error) {
	names, err := GetNames(pathSA, pathCA)
	if err != nil {
		return nil, err
	}

	res, err := json.MarshalIndent(names, "", "  ")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetNamesDM(pathSA, pathCA string) (map[string][]string, error) {
	res := make(map[string][]string)

	names, err := GetNames(pathSA, pathCA)
	if err != nil {
		return nil, err
	}

	for _, name := range names {
		for idx, namePart := range name {
			// First idx is the DocumentedName's ID value.
			if idx == 0 {
				continue
			}

			namesIn := strings.Split(namePart, " ")

			namesDM := NewDoubleMetaphone(namesIn...)

			for _, nameDM := range namesDM {
				nameDMStr := strings.Join(nameDM, " ")

				if nameDMStr == "" {
					continue
				}

				res[nameDMStr] = name
			}
		}
	}

	return res, nil
}

func GetNamesDMJSON(pathSA, pathCA string) ([]byte, error) {
	names, err := GetNamesDM(pathSA, pathCA)
	if err != nil {
		return nil, err
	}

	res, err := json.MarshalIndent(names, "", "  ")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetNamesDB(pathSA, pathCA string) (map[string][]string, error) {
	res := make(map[string][]string)

	names, err := GetNames(pathSA, pathCA)
	if err != nil {
		return nil, err
	}

	for _, name := range names {
		// NOTE: First idx is the DocumentedName's ID value.
		for _, namePart := range name {
			/* Remove this weird "linebreak" entry. Maybe we shouldn't
			remove it but it looks wrong. */
			if namePart == "\r\n                " {
				continue
			}

			res[namePart] = name
		}
	}

	namesDM, err := GetNamesDM(pathSA, pathCA)
	if err != nil {
		return nil, err
	}

	for key, val := range namesDM {
		res[key] = val
	}

	return res, nil
}

func GetNamesDBJSON(pathSA, pathCA string) ([]byte, error) {
	names, err := GetNamesDB(pathSA, pathCA)
	if err != nil {
		return nil, err
	}

	res, err := json.MarshalIndent(names, "", "  ")
	if err != nil {
		return nil, err
	}

	return res, nil
}
