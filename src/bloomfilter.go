package bloomfilter

import (
	"fmt"
	_ "hash"
	_ "hash/fnv"
	"math"
)

type Interface interface {
	Add (item []byte)
	Test (item []byte) bool
}

type BloomFilter struct {
	arrayBuffer []bool
	m 			uint
	k           uint
	n           uint
	v           uint
}



func New(size uint, k uint) BloomFilter {

	//var n = math.Ceil(float64(size) / 32)
	var kbytes = math.Ceil(math.Log(math.Ceil(math.Log(float64(size)) / math.Ln2 / 8)) / math.Ln2)
	kbytes = 1 << uint(kbytes)
	var arrayBuffer = make([]bool, uint(kbytes) * k)
	//var bucket = [uint(n)]int{}
	return BloomFilter{
		arrayBuffer: arrayBuffer,
		m:           size,
		k:           k, // we have 3 hash functions for now
		n:           uint(0),
	}
}

func (bf BloomFilter) locations(v string) float64 {
	
	//var k = bf.k
	//var m = bf.m
	//var a = fnv_1a(v)
	//var b = fnv_1a(v, 1576284489)
	//
	//return
}

func (bf BloomFilter) add(v) float64 {
	fmt.Printf(v)
	return v
}

func (bf BloomFilter) test(v) float64 {
	fmt.Printf(v)
	return v
}

func (bf BloomFilter) size(v) float64  {
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