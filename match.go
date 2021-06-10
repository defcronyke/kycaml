package kycaml

import (
	"fmt"
	"os"

	"github.com/antzucaro/matchr"
)

func NewDoubleMetaphone(req ...string) [][]*matchr.String {
	res1 := make([]*matchr.String, 0, len(os.Args)-1)
	res2 := make([]*matchr.String, 0, len(os.Args)-1)

	for _, r := range req {
		s1, s2 := matchr.DoubleMetaphone(r)

		r1 := matchr.NewString(s1)
		r2 := matchr.NewString(s2)

		res1 = append(res1, r1)
		res2 = append(res2, r2)
	}

	res := [][]*matchr.String{res1, res2}

	return res
}

func DoubleMetaphoneToString(req [][]*matchr.String) ([][]string, error) {
	if len(req) < 2 {
		return nil, fmt.Errorf("req must have at least 2 elements: num elements: %v", len(req))
	}

	res1 := make([]string, 0, len(req[0]))
	res2 := make([]string, 0, len(req[1]))

	for _, r := range req[0] {
		res1 = append(res1, r.String())
	}

	for _, r := range req[1] {
		res2 = append(res2, r.String())
	}

	res := [][]string{res1, res2}

	return res, nil
}
