package py

type Set[K comparable] struct {
	d map[K]struct{}
}

func NewSet[K comparable]() *Set[K] {
	return &Set[K]{
		d: map[K]struct{}{},
	}
}

func (s *Set[K]) Has(k K) bool {
	_, exists := s.d[k]
	return exists
}

func (s *Set[K]) Add(values ...K) bool {
	oldSize := len(s.d)
	for _, v := range values {
		s.d[v] = struct{}{}
	}
	return len(s.d) != oldSize
}

func (s *Set[K]) Size() int {
	return len(s.d)
}

func (s *Set[K]) Union(other Set[K]) {
	for k := range other.d {
		s.d[k] = other.d[k]
	}
}

func (s Set[K]) Values() []K {
	r := []K{}
	for v := range s.d {
		r = append(r, v)
	}
	return r
}

func (s Set[K]) Intersection(ks ...K) *Set[K] {
	r := NewSet[K]()
	for _, k := range ks {
		if _, ok := s.d[k]; ok {
			r.d[k] = struct{}{}
		}
	}
	return r
}
