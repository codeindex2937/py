package py

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Sum[K addable](vs ...K) K {
	var s K
	for _, v := range vs {
		s += v
	}
	return s
}

func Min[K numeric](vs ...K) (K, error) {
	if len(vs) < 1 {
		return 0, ValueError
	}

	r := vs[0]
	for _, v := range vs {
		if r > v {
			r = v
		}
	}
	return r, nil
}

func Max[K numeric](vs ...K) (K, error) {
	if len(vs) < 1 {
		return 0, ValueError
	}

	r := vs[0]
	for _, v := range vs {
		if r < v {
			r = v
		}
	}
	return r, nil
}

func Any(vs ...bool) bool {
	for _, v := range vs {
		if v {
			return true
		}
	}
	return false
}

func All(vs ...bool) bool {
	for _, v := range vs {
		if !v {
			return false
		}
	}
	return true
}

func MapSlice[K any, V any](mapFn func(K) V, vs ...K) []V {
	r := []V{}
	for _, v := range vs {
		r = append(r, mapFn(v))
	}
	return r
}

func Filter[K any](filterFn func(k K) bool, vs ...K) []K {
	r := []K{}
	for _, v := range vs {
		if filterFn(v) {
			r = append(r, v)
		}
	}
	return r
}

func Sorted[K constraints.Ordered](ks []K) []K {
	clone := make([]K, len(ks))
	copy(clone, ks)
	slices.Sort(clone)
	return clone
}
