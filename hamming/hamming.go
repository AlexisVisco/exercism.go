package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var diff = 0
	if len(a) != len(b) {
		return -1, errors.New("not sames length")
	}
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
