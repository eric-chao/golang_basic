package helper

import (
	"crypto/md5"
	"crypto/sha1"

	"github.com/spaolacci/murmur3"
)

func md5Hash(str string) [16]byte {
	return md5.Sum([]byte(str))
}

func sha1Hash(str string) [20]byte {
	return sha1.Sum([]byte(str))
}

func murmur32(str string) uint32 {
	return murmur3.Sum32([]byte(str))
}

func murmur64(str string) uint64 {
	return murmur3.Sum64([]byte(str))
}