// Package gene provides compressed storage for a string of nucleotides.
package gene

import (
	"fmt"
	"strings"

	"gopl.io/ch6/intset"
)

// CompressedGene represents a string of nucleotides.
type CompressedGene struct {
	length int
	set    intset.IntSet
}

// Compress initializes a CompressedGene from a string of nucleotides.
func (cg *CompressedGene) Compress(gene string) error {
	cg.length = len(gene)
	cg.set = intset.IntSet{}
	for index, nucleotide := range strings.ToUpper(gene) {
		pos := index * 2
		switch nucleotide {
		case 'A':
			// 00
			break
		case 'C':
			// 01
			cg.set.Add(pos + 1)
		case 'G':
			// 10
			cg.set.Add(pos)
		case 'T':
			// 11
			cg.set.Add(pos)
			cg.set.Add(pos + 1)
		default:
			return fmt.Errorf("unexpected character %q at position %d",
				nucleotide, index)
		}
	}
	return nil
}

// btoi is the boolean to int function.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// String returns a CompressedGene as a string of nucleotides.
func (cg CompressedGene) String() (result string) {
	for i := 0; i < cg.length; i++ {
		pos := i * 2
		value := btoi(cg.set.Has(pos))*2 + btoi(cg.set.Has(pos+1))
		switch value {
		case 0: // A = 00
			result += "A"
		case 1: // C = 01
			result += "C"
		case 2: // G = 10
			result += "G"
		case 3: // T = 11
			result += "T"
		}
	}
	return result
}
