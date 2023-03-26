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

func Repeat[K any](k K, times int) []K {
	result := []K{}
	for i := 0; i < times; i++ {
		result = append(result, k)
	}
	return result
}

func Range[K numeric](start, stop, step K) []K {
	result := []K{}
	for k := start; k < stop; {
		result = append(result, k)
	}
	return result
}

func Combination[K any](slice []K, k int) [][]K {
	if k < 0 || k > len(slice) {
		return [][]K{}
	}

	index := make([]int, k)
	for i := 0; i < k; i++ {
		index[i] = i
	}

	var result [][]K

	for {
		comb := make([]K, k)
		for i, j := range index {
			comb[i] = slice[j]
		}
		result = append(result, comb)

		i := k - 1
		for ; i >= 0; i-- {
			if index[i]+1 != len(slice)-(k-i-1) {
				break
			}
		}

		if i < 0 {
			break
		}

		index[i]++
		for j := i + 1; j < k; j++ {
			index[j] = index[j-1] + 1
		}
	}

	return result
}

func Product[K any](slices ...[]K) [][]K {
	result := [][]K{}
	for _, s := range slices {
		next := [][]K{}
		for _, v := range s {
			for _, r := range result {
				next = append(next, append(r, v))
			}
		}
		result = next
	}
	return result
}

func GroupBy[K any, V comparable](slice []K, vFn func(k K) V) map[V][]K {
	result := map[V][]K{}
	for _, k := range slice {
		g := vFn(k)
		if vs, ok := result[g]; ok {
			result[g] = append(vs, k)
		} else {
			result[g] = []K{k}
		}
	}
	return result
}

func Zip[K any](slices ...[]K) [][]K {
	result := [][]K{}
	if len(slices) < 1 {
		return result
	}

	minLen, _ := Min(MapSlice(func(slice []K) int { return len(slice) }, slices...)...)
	for i := 0; i < minLen; i++ {
		r := []K{}
		for _, slice := range slices {
			r = append(r, slice[i])
		}
		result = append(result, r)
	}
	return result
}

func ZipLongest[K any](k K, slices ...[]K) [][]K {
	result := [][]K{}
	if len(slices) < 1 {
		return result
	}

	maxLen, _ := Max(MapSlice(func(slice []K) int { return len(slice) }, slices...)...)
	for i := 0; i < maxLen; i++ {
		r := []K{}
		for _, slice := range slices {
			if len(slice) <= i {
				r = append(r, k)
			} else {
				r = append(r, slice[i])
			}
		}
		result = append(result, r)
	}
	return result
}
