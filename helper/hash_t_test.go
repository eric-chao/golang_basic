package helper

import (
	"fmt"
	"testing"
)

var bucketSize = 10

// murmur hash 均匀分布
func TestMurmurHash(t *testing.T) {
	var bucketMap = map[uint64]int{}
	for i := 15000000000; i < 15000000000+10000000; i++ {
		hashInt := murmur64(fmt.Sprintf("%d", i)) % uint64(bucketSize)
		bucketMap[hashInt]++
	}
	t.Log(bucketMap)
}

func TestHash(t *testing.T) {
	var str = "xxxyyyzzz"
	var murmur32 = murmur64(str)
	var md5hash = md5Hash(str)
	t.Logf("murmur32: %x", murmur32)
	t.Logf("md5hash: %x", md5hash)
}
