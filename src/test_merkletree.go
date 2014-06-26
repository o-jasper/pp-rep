package main

import (
	"fmt"
	"math/rand"
	"merkletree"
)

func rand_bytes(n int) []byte {
	out := []byte{}
	for i := 0 ; i < n ; i++ {
		out = append(out, byte(rand.Intn(256)))
	}
	return out
}

func bytes_as_hex(input []byte) string {
	hex := "0123456789ABCDEF"
	out := ""
	for i := range input {
		val := uint8(input[i])
		out = string(append([]byte(out), hex[val%16]))
		out = string(append([]byte(out), hex[val/16]))
	}
	return out
}

func main() {
	a := rand_bytes(32)
	fmt.Println("a: ", bytes_as_hex(a))
	b := rand_bytes(32)
	fmt.Println("b: ", bytes_as_hex(b))
	
	gen := merkletree.MerkleTreeGen{List:[]MerkleTreePortion{}}
	gen.AddChunk(a, true)
	gen.AddChunk(b, true)

	fmt.Println(gen.List)
}
