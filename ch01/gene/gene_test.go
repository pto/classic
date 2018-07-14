package gene

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	cases := []struct {
		input   string
		isError bool
	}{
		{"", false},
		{"A", false},
		{"C", false},
		{"G", false},
		{"T", false},
		{"Q", true},
		{"AAAAATTTTTGGGGGCCCCC", false},
		{"AAAAAAAAAAAAAAAAAX", true},
		{"ATGAATGCC", false},
	}
	var cg CompressedGene
	for _, c := range cases {
		err := cg.Compress(c.input)
		gotError := err != nil
		if gotError != c.isError {
			t.Errorf("%q is error %t, want %t", c.input, gotError, c.isError)
		} else if err == nil && cg.String() != c.input {
			t.Errorf("%q uncompresses as %q", c.input, cg.String())
		}
	}
}

func Example() {
	var cg CompressedGene
	err := cg.Compress("ACGT")
	fmt.Println(err, cg)
	err = cg.Compress("AAAAX")
	fmt.Println(err)
	// Output:
	// <nil> ACGT
	// unexpected character 'X' at position 4
}
