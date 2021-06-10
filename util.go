package kycaml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	newFile, err := os.Open(fmt.Sprintf("%v/%v", os.Getenv("PWD"), path))
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

func NewSanctionsCA() (*SanctionsCA, error) {
	newFile, err := NewFile("./static/cons.xml")
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
