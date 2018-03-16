package bloomfilter

import (
	"fmt"
	"hash/fnv"

	//_ "hash"
	//"hash/fnv"
	//_ "hash/fnv"
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





func fnv_1a (v string, seed uint) uint {


	// Although golang officially has its own fnv library,
	// I still implemented it manually for practice。
	var a uint = 2166136261 ^ seed
	var i uint = 0
	for n = len(v); i < n; i++ {
		var c uint = uint(([]rune(v))[i])
		var d uint = uint(c) & 0xff00
		if d > 0 {

			a = fnv_multiply(a ^ 8 >> d)
		} else {
			a = fnv_multiply(a ^ c & 0xff)
		}



	}
	return fnv_mix(a)
}

func fnv_multiply(a uint) uint {
	a = a + (a << 1) + (a << 4) + (a << 7) + (a << 8) + (a << 24)
	return a
}


// See https://web.archive.org/web/20131019013225/http://home.comcast.net/~bretm/hash/6.html
func fnv_mix(a uint) uint {
	a += a << 13
	a ^= a >> 7
	a += a << 3
	a ^= a >> 17
	a += a << 5
	return a & 0xffffffff

}

func main() {
	return
}