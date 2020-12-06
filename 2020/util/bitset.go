package util

type BitSet interface {
	Has(i int) bool
	Set(i int, v bool)
	Count() int
	And(other ...BitSet) BitSet
	Or(other ...BitSet) BitSet
	Xor(other ...BitSet) BitSet
}

type bitset struct {
	bits   []bool
	length int
}

func NewBitSet(length int) BitSet {
	return bitset{
		length: length,
		bits:   make([]bool, length),
	}
}

func (b bitset) Has(i int) bool {
	return i < b.length && b.bits[i]
}

func (b bitset) Set(i int, v bool) {
	if i >= b.length {
		return
	}
	b.bits[i] = v
}

func (b bitset) Count() int {
	c := 0
	for i := 0; i < b.length; i++ {
		if b.Has(i) {
			c++
		}
	}
	return c
}

func (b bitset) And(other ...BitSet) BitSet {
	result := bitset{
		bits:   b.bits,
		length: b.length,
	}
	for _, bits := range other {
		for i := 0; i < result.length; i++ {
			result.Set(i, result.Has(i) && bits.Has(i))
		}
	}

	return result
}

func (b bitset) Or(other ...BitSet) BitSet {
	result := bitset{
		bits:   b.bits,
		length: b.length,
	}
	for _, bits := range other {
		for i := 0; i < result.length; i++ {
			result.Set(i, result.Has(i) || bits.Has(i))
		}
	}

	return result
}

func (b bitset) Xor(other ...BitSet) BitSet {
	result := bitset{
		bits:   b.bits,
		length: b.length,
	}
	for _, bits := range other {
		for i := 0; i < result.length; i++ {
			result.Set(i, result.Has(i) != bits.Has(i))

		}
	}

	return result
}
