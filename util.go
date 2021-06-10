package kycaml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func NewSanctionsCA(path string) (*SanctionsCA, error) {
	pathDefault := "./static/cons.xml"

	if path == "" {
		path = pathDefault
	}

	newFile, err := NewFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var sanctionsCA SanctionsCA

	err = xml.Unmarshal(newFile, &sanctionsCA)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &sanctionsCA, nil
}

func NewJSONSanctionsCA(path string) ([]byte, error) {
	sanctionsCA, err := NewSanctionsCA(path)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	resBytes, err := json.MarshalIndent(sanctionsCA, "", "  ")
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return resBytes, nil
}
