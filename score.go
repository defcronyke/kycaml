package kycaml

import "fmt"

type Score struct {
}

func NewScore() *Score {
	s := Score{}

	return &s
}

func (s *Score) Init() {
	fmt.Printf("%+v", s)
}
