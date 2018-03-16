package bloomfilter

import (
	"math/rand"
	"testing"
)

// 100 Million
var MemberSize uint = 100000000

// 100 Thousand
var SampleSize int = 100000

//the number of hashing functions
var DefaultHashFunctions uint = 3


func RandomBytes(size int) []byte {
	b := make([]byte, size)
	rand.Read(b)
	return b
}

func RandomString() string {
	return string(RandomBytes(10))
}

func TestExistance(t *testing.T) {
	//bf := New(MemberSize, DefaultHashFunctions)
	//
	//for i := 0; i < SampleSize; i++ {
	//	items
	}
}