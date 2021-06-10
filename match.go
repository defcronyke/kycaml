package kycaml

import (
	"fmt"
	"os"

	"github.com/antzucaro/matchr"
)

func NewDamerauLevenshtein(s1, s2 string) (distance int) {
	return matchr.DamerauLevenshtein(s1, s2)
}

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

func NewHamming(s1, s2 string) (int, error) {
	res, err := matchr.Hamming(s1, s2)
	if err != nil {
		return -1, err
	}

	return res, nil
}

func NewJaro(s1, s2 string) (distance float64) {
	return matchr.Jaro(s1, s2)
}

func NewJaroWinkler(s1, s2 string, lTol bool) (distance float64) {
	return matchr.JaroWinkler(s1, s2, lTol)
}

func NewLevenshtein(s1, s2 string) (distance int) {
	return matchr.Levenshtein(s1, s2)
}

func NewLongestCommonSubsequence(s1, s2 string) int {
	return matchr.LongestCommonSubsequence(s1, s2)
}

func NewNYSIIS(req ...string) []string {
	res := make([]string, 0, len(req))

	for _, r := range req {
		res1 := matchr.NYSIIS(r)
		res = append(res, res1)
	}

	return res
}

func NewOSA(s1, s2 string) (distance int) {
	return matchr.OSA(s1, s2)
}

func NewPhonex(req ...string) []string {
	res := make([]string, 0, len(req))

	for _, r := range req {
		res1 := matchr.Phonex(r)
		res = append(res, res1)
	}

	return res
}

func NewSmithWaterman(s1, s2 string) float64 {
	return matchr.SmithWaterman(s1, s2)
}
