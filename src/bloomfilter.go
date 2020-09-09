package bloomfilter

import (
	"hash/fnv"
	"math"
)

//type Bloom interface {
//	Add(v string)
//	Locations(v string) []uint
//	Test(v string) bool
//	Size() float64
//}

type BloomFilter struct {
	arrayBuffer []bool
	_locations  []uint
	bucket      []uint
	m           uint64
	k           uint
	n           uint
	v           uint
}

func NewBloom(size uint64, k uint) *BloomFilter {
	
	var n = math.Ceil(float64(size) / 32)
	var kbyte = math.Ceil(math.Log(math.Ceil(math.Log(float64(size))/math.Ln2/8)) / math.Ln2)
	var i uint = 1
	var kbytes = i << uint(kbyte)
	var arrayBuffer = make([]bool, uint(kbytes)*k)
	var bucket = make([]uint, int(n))
	var _locations = make([]uint, k)
	return &BloomFilter{
		arrayBuffer: arrayBuffer,
		_locations:  _locations,
		bucket:      bucket,
		m:           size,
		k:           k, // we have 3 hash functions for now
		n:           uint(0),
		v:           0,
	}
}

func fingerprint(b []byte) uint64 {
	hash := fnv.New64a()
	_, _ = hash.Write(b)
	return hash.Sum64()
}

func (bf *BloomFilter) Locations(v string) []uint {
	
	var k = bf.k
	var m = bf.m
	var r = bf._locations
	
	var a = fingerprint([]byte(v))
	var b = fingerprint([]byte(v))
	var x = a % m
	var i uint = 0
	for ; i < k; i++ {
		if x < 0 {
			r[i] = uint(x + m)
		} else {
			r[i] = uint(x)
		}
		x = (x + b) % m
	}
	return r
}

func (bf *BloomFilter) Add(v string) {
	var l = bf.Locations(v + "")
	var k = bf.k
	var bucket = bf.bucket
	var i uint = 0
	for i = 0; i < k; i++ {
		bucket[int(math.Floor(float64(l[i]/32)))] |= 1 << (l[i] % 32)
	}
}

func (bf *BloomFilter) Test(v string) bool {
	var l = bf.Locations(v + "")
	var k = bf.k
	var bucket = bf.bucket
	var i uint = 0
	for ; i < k; i++ {
		var b = l[i]
		if (bucket[int(math.Floor(float64(b/32)))] & (1 << (b % 32))) == 0 {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) Size() float64 {
	var bucket = bf.bucket
	var bits uint = 0
	var n = len(bucket)
	var result float64
	for i := 0; i < n; i++ {
		bits += PopCount(bucket[i])
	}
	result = -(float64(bf.m) * math.Log(float64(1-bits)/float64(bf.m)) / float64(bf.k))
	return result
}

func PopCount(v uint) uint {
	v -= (v >> 1) & 0x55555555
	v = (v & 0x33333333) + ((v >> 2) & 0x33333333)
	return ((v + (v>>4)&0xf0f0f0f) * 0x1010101) >> 24
}
