package py

import "encoding/json"

type DefaultMap[K comparable, V any] struct {
	d         map[K]V
	defaultFn func() V
}

func NewDefaultMap[K comparable, V any](defaultFn func() V) *DefaultMap[K, V] {
	return &DefaultMap[K, V]{
		d:         map[K]V{},
		defaultFn: defaultFn,
	}
}

func NewDefaultMapWith[K comparable, V any](defaultFn func() V) func() *DefaultMap[K, V] {
	return func() *DefaultMap[K, V] {
		return &DefaultMap[K, V]{
			d:         map[K]V{},
			defaultFn: defaultFn,
		}
	}
}

func (m *DefaultMap[K, V]) Get(k K) V {
	if v, ok := m.d[k]; ok {
		return v
	} else {
		v = m.defaultFn()
		m.d[k] = v
		return v
	}
}

func (m *DefaultMap[K, V]) Set(k K, v V) {
	m.d[k] = v
}

func (m DefaultMap[K, V]) Size() int {
	return len(m.d)
}

func (m DefaultMap[K, V]) Keys() []K {
	ks := []K{}
	for k := range m.d {
		ks = append(ks, k)
	}
	return ks
}

func (m DefaultMap[K, V]) Values() []V {
	r := []V{}
	for _, v := range m.d {
		r = append(r, v)
	}
	return r
}

func (m DefaultMap[K, V]) Items() map[K]V {
	r := map[K]V{}
	for k, v := range m.d {
		r[k] = v
	}
	return r
}

func (m DefaultMap[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.d)
}
