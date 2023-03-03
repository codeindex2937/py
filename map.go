package py

type Map[K comparable, V any] map[K]V
type iMap[K comparable, V any] interface {
	Items() map[K]V
	Set(K, V)
}

func (m Map[K, V]) Values() []V {
	r := []V{}
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func (m Map[K, V]) Items() map[K]V {
	r := map[K]V{}
	for k, v := range m {
		r[k] = v
	}
	return r
}

func (m Map[K, V]) Set(k K, v V) {
	m[k] = v
}

func UpdateMap[K comparable, V any](t iMap[K, V], s iMap[K, V]) {
	for k, v := range s.Items() {
		t.Set(k, v)
	}
}
