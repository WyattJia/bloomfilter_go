package bloomfilter

import (
	. "math"
)

type Interface interface {
	Add(v string)
	Locations(v string) []uint
	Test(v string) bool
	Size() float64
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

func (bf BloomFilter) Locations(v string) []uint {
	
	var k = bf.k
	var m = bf.m
	var r = bf._locations

	var a = Fnv1a(v, 0)
	var b = Fnv1a(v, 1576284489)
	var x = a % m
	var i uint = 0
	for ; i < k; i++ {
		if x < 0 {
			r[i] = x + m
		} else {
			r[i] = x
		}
		x = (x + b) % m
	}
	return r
}

func (bf BloomFilter) Add(v string) {
	var l = bf.Locations(v + "")
	var k = bf.k
	var bucket = bf.bucket
	var i uint = 0
	for i = 0; i < k; i++ {
		bucket[int(Floor(float64(l[i]/32)))] |= 1 << (l[i] % 32)
	}
}

func (bf BloomFilter) Test(v string) bool {
	var l = bf.Locations(v + "")
	var k = bf.k
	var bucket = bf.bucket
	var i uint = 0
	for ; i < k; i++{
		var b = l[i]
		if (bucket[int(Floor(float64(b/32)))] & (1 << (b % 32))) == 0 {
			return false
		}
	}
	return true
}

func (bf BloomFilter) Size() float64  {
	var bucket = bf.bucket
	var bits uint = 0
	var n=len(bucket)
	var result float64
	for i := 0; i < n; i++ {
		bits += PopCount(bucket[i])
	}
	result = -(float64(bf.m) * Log(float64(1 - bits) / float64(bf.m)) / float64(bf.k))
	return result
}

func PopCount (v uint) uint {
	v -= (v >> 1) & 0x55555555
	v = (v & 0x33333333) + ((v >> 2) & 0x33333333)
	return ((v + (v >> 4) & 0xf0f0f0f) * 0x1010101) >> 24
}


func Fnv1a (v string, seed uint) uint {
	// Although golang officially has its own fnv library,
	// I still implemented it manually for practice。
	var a uint = 2166136261 ^ seed
	var i uint = 0
	var n uint = uint(len(v))
	for ; i < n; i++ {
		var c uint = uint(([]rune(v))[i])
		var d uint = uint(c) & 0xff00
		if d > 0 {

			a = FnvMultiply(a ^ 8 >> d)
		} else {
			a = FnvMultiply(a ^ c & 0xff)
		}
	}
	return FnvMix(a)
}

func FnvMultiply(a uint) uint {
	a = a + (a << 1) + (a << 4) + (a << 7) + (a << 8) + (a << 24)
	return a
}


// See https://web.archive.org/web/20131019013225/http://home.comcast.net/~bretm/hash/6.html
func FnvMix(a uint) uint {
	a += a << 13
	a ^= a >> 7
	a += a << 3
	a ^= a >> 17
	a += a << 5
	return a & 0xffffffff
}
