package main

import (
	"fmt"
	"iter"
	"slices"
)

func Equal(lhs iter.Seq[int], rhs iter.Seq[int]) bool {
	lhsNext, lhsStop := iter.Pull(lhs)
	rhsNext, rhsStop := iter.Pull(rhs)

	defer func() {
		lhsStop()
		rhsStop()
	}()

	for {
		lhsValue, lhsOk := lhsNext()
		rhsValue, rhsOk := rhsNext()

		if !lhsOk {
			return !rhsOk
		}

		if lhsOk != rhsOk || lhsValue != rhsValue {
			return false
		}
	}
}

func main() {
	seq1 := slices.Values([]int{1, 2, 3})
	seq2 := slices.Values([]int{1, 2, 3})

	fmt.Println(Equal(seq1, seq2))
}
