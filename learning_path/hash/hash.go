package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))

	v := h.Sum32()
	fmt.Println(v)

	k := sha1.New()
	k.Write([]byte("test"))

	bs := k.Sum([]byte{})
	fmt.Println(bs)
}
