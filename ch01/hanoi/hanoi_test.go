package hanoi

import "testing"

func TestSolve(t *testing.T) {
	cases := []Tower{
		Tower{},
		Tower{1},
		Tower{1, 2},
		Tower{1, 2, 3},
		Tower{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
	for _, c := range cases {
		from := Tower{}
		from = append(from, c...)
		if from.String() != c.String() {
			t.Errorf("copy failed from %v to %v", c, from)
		}
		to := Tower{}
		temp := Tower{}
		Solve(&from, &to, &temp, len(c))
		if from.String() != "[]" || temp.String() != "[]" ||
			to.String() != c.String() {
			t.Errorf("invalid result for %v: %v %v %v", c, from, to, temp)
		}
	}
}
