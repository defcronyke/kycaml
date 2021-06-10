package kycaml

import (
	"os"

	"github.com/antzucaro/matchr"
)

func NewDamerauLevenshtein(s1, s2 string) (distance int) {
	return matchr.DamerauLevenshtein(s1, s2)
}

func NewDoubleMetaphone(req ...string) [][]string {
	res1 := make([]string, 0, len(os.Args)-1)
	res2 := make([]string, 0, len(os.Args)-1)

	for _, r := range req {
		s1, s2 := matchr.DoubleMetaphone(r)

		res1 = append(res1, s1)
		res2 = append(res2, s2)
	}

	res := [][]string{res1, res2}

	return res
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

func NewSoundex(req ...string) []string {
	res := make([]string, 0, len(req))

	for _, r := range req {
		res1 := matchr.Soundex(r)
		res = append(res, res1)
	}

	return res
}

func NewSmithWaterman(s1, s2 string) float64 {
	return matchr.SmithWaterman(s1, s2)
}
