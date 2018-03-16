package bloomfilter

import (
	"math/rand"
	"sort"
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
	bf := New(MemberSize, DefaultHashFunctions)

	for i := 0; i < SampleSize; i++ {
		item := RandomString()
		bf.Add(item)

		if bf.Test(item) != true {
			t.Errorf("'%q' not found", item)
		}

		// Now lets create some items that don't exist
		item2 := RandomString()

		// Test that item does NOT exist
		if bf.Test(item2) == true {
			t.Errorf("'%q' should not be found", item2)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	bf := New(MemberSize, DefaultHashFunctions)
	for i := 0; i < b.N; i++ {
		bf.Add(RandomString())
	}
}

func BenchmarkTest(b *testing.B) {
	bf := New(MemberSize, DefaultHashFunctions)
	for i := 0; i < b.N; i++ {
		bf.Add(RandomString())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Test(RandomString())
	}
}

func use(interface{}) {}

func BenchmarkBinarySearch(b *testing.B) {
	var strings []string
	for i := 0; i < SampleSize; i++ {
		item := RandomString()
		strings = append(strings, item)
	}

	// Sort by byte order
	sort.Strings(strings)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		item := RandomString()
		use(sort.SearchStrings(strings, item))
	}
}


