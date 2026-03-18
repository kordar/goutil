package goutil

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash/crc32"
)

func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

func SHA1B(b []byte) string {
	o := sha1.New()
	o.Write(b)
	return hex.EncodeToString(o.Sum(nil))
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5B(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func HashCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func CRC32(bytes []byte) uint32 {
	return crc32.ChecksumIEEE(bytes)
}

func CRC32Str(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
