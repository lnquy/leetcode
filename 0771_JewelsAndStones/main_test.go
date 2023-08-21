package main

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	tcs := []struct {
		inJ string
		inS string
		out int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},

		{"", "", 0},
		{"a", "", 0},
		{"", "asdzxc", 0},
		{"a", "asdzxc", 1},
		{"a", "asdAzxc", 1},
		{"a", "asdAzxcaefqea", 3},
		{"as", "asdAzxcaefqea", 4},
		{"ase", "asdAzxcaefqea", 6},
		{"asdAzxcefq", "asdAzxcaefqea", 13},
	}

	for idx, tc := range tcs {
		out := numJewelsInStones(tc.inJ, tc.inS)
		if out != tc.out {
			t.Errorf("%d. numJewelsInStones(%q, %q) = %d. Expected: %d", idx, tc.inJ, tc.inS, out, tc.out)
		}
	}
}
