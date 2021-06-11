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

	"github.com/defcronyke/kycaml/model/cons"
	"github.com/defcronyke/kycaml/model/sdn"
)

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

func SetHeader(header, value string, handle http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set(header, value)
		handle.ServeHTTP(w, req)
	}
}
