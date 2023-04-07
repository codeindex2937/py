package py

type Counter[K comparable] struct {
	d map[K]int
}

func NewCounter[K comparable]() *Counter[K] {
	return &Counter[K]{
		d: make(map[K]int),
	}
}

func (c *Counter[K]) Add(ks ...K) {
	for _, k := range ks {
		c.d[k]++
	}
}

func (c *Counter[K]) Get(k K) int {
	if v, ok := c.d[k]; ok {
		return v
	}
	return 0
}

func (c *Counter[K]) Size() int {
	return len(c.d)
}

func (c *Counter[K]) Keys() []K {
	ks := []K{}
	for k := range c.d {
		ks = append(ks, k)
	}
	return ks
}

func (c *Counter[K]) Values() []int {
	vs := []int{}
	for _, v := range c.d {
		vs = append(vs, v)
	}
	return vs
}

func (c *Counter[K]) Items() map[K]int {
	items := map[K]int{}
	for k, v := range c.d {
		items[k] = v
	}
	return items
}
