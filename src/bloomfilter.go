package bloomfilter

import (
	"fmt"
	_ "hash"
	_ "hash/fnv"
)

type Interface interface {
	Add (item []byte)
	Test (item []byte) bool
}

type BloomFilter struct {
	arrayBuffer []byte
	a           uint
	m           uint
	k           uint
	n           uint
	v 			uint
}

func (bf *BloomFilter) New(size) *BloomFilter {
	return &BloomFilter{
		arrayBuffer: make([]bool, size),
		a:           0,
		m:           size,
		k:           3, // we have 3 hash functions for now
		n:           uint(0),
		v:           0,
	}
}

func (bf *BloomFilter) locations(v) float64 {
	fmt.Printf(v)
	return v
}

func (bf *BloomFilter) add(v) float64 {
	fmt.Printf(v)
	return v
}

func (bf *BloomFilter) test(v) float64 {
	fmt.Printf(v)
	return v
}

func (bf *BloomFilter) size(v) float64  {
    fmt.Printf(v)
	return v
}

func popcnt (v) {
	fmt.Printf(v)
}

func fnv_1a (v, seed) {
	fmt.Printf(v)
	fmt.Printf(seed)
}

func fnv_multiply(a) float64 {
	fmt.Printf(a)
	return a
}

func fnv_mix(a) float64 {
	fmt.Printf(a)
	return a

}

func main() {
	return
}