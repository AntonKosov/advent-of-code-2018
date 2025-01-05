package slice

import "iter"

func FromSeq[T any](seq iter.Seq[T]) []T {
	var s []T
	for v := range seq {
		s = append(s, v)
	}

	return s
}
