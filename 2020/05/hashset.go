package main

type hashset struct {
	m map[int]bool
}

func NewHashset() hashset {
	return hashset{
		m: make(map[int]bool),
	}
}
func (h hashset) Contains(id int) bool {
	return h.m[id]
}

func (h hashset) Add(id int) {
	h.m[id] = true
}
