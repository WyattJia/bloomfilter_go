package bloomfilter

import (
	"fmt"
	"github.com/Go-zh/tools/go/analysis/passes/vet/testdata/divergent"
	"github.com/boltdb/bolt"
	"hash/fnv"

	//_ "hash"
	//"hash/fnv"
	//_ "hash/fnv"
	. "math"
)

type Interface interface {
	Add (item []byte)
	Test (item []byte) bool
}

type BloomFilter struct {
	arrayBuffer []bool
	_locations  []uint
	bucket 		[]uint
	m 			uint
	k           uint
	n           uint
	v           uint
}



func New(size uint, k uint) BloomFilter {

	var n = Ceil(float64(size) / 32)
	var kbytes = Ceil(Log(Ceil(Log(float64(size)) /Ln2/ 8)) / Ln2)
	kbytes = 1 << uint(kbytes)
	var arrayBuffer = make([]bool, uint(kbytes) * k)
	//var bucket = [uint(n)]int{}
	var bucket = make([]uint, n)
	var _locations = make([]uint, k)
	return BloomFilter{
		arrayBuffer: arrayBuffer,
		_locations:  _locations,
		bucket:		 bucket,
		m:           size,
		k:           k, // we have 3 hash functions for now
		n:           uint(0),
		v:           0,
	}
}

func (bf BloomFilter) locations(v string) []uint {
	
	var k = bf.k
	var m = bf.m
	var r = bf._locations

	var a = fnv_1a(v, 0)
	var b = fnv_1a(v, 1576284489)
	var x = a % m

	for i = 0; i < k; i++ {
		if x < 0 {
			r[i] = x + m
		} else {
			r[i] = x
		}
		x = (x + b) % m
	}
	return r
}

func (bf BloomFilter) add(v string) {
	var l = bf.locations(v + "")
	k = bf.k
	bucket = bf.bucket
	for i = 0; i < k; i++ {
		bucket[Floor(float64(l[i]/32))] |= 1 << (l[i] % 32)
	}
}

func (bf BloomFilter) test(v) float64 {
	fmt.Printf(v)
	return v
}

func (bf BloomFilter) size(v) float64  {
    fmt.Printf(v)
	return v
}

func popcnt (uint v) uint {
	v -= (v >> 1) & 0x55555555
	v = (v & 0x33333333) + ((v >> 2) & 0x33333333)
	return ((v + (v >> 4) & 0xf0f0f0f) * 0x1010101) >> 24
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