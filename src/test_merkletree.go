package main

import (
	"fmt"
	"flag"
	"math/rand"
	"merkletree"
)

// Print helpers.
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

//Generating chunks.
func rand_bytes(r *rand.Rand, n int32) []byte {
	out := []byte{}
	for i := int32(0) ; i < n ; i++ {
		out = append(out, byte(r.Int63n(256)))
	}
	return out
}

func rand_range(r *rand.Rand, fr int32, to int32) int32 {
	return fr + r.Int31n(to - fr)
}

func rand_chunk(r *rand.Rand, n_min int32, n_max int32) []byte {
	return rand_bytes(r, rand_range(r, n_min, n_max))
}

//Adds a lot of chunks and lists the tree leaves.
func run_test(seed int64, n_min int32, n_max int32, N int) {
	r := rand.New(rand.NewSource(seed))

	gen := merkletree.NewMerkleTreeGen()  //Put chunks in.
	list := []*merkletree.MerkleNode{}
	for i:= 0 ; i < N ; i++ {
		list = append(list, gen.AddChunk(rand_chunk(r, n_min, n_max), true))
	}
	roothash := gen.Finish().Hash  //Get the root hash.

//Reset random function, doing exact same to it.
	r = rand.New(rand.NewSource(seed))
	for i := 0 ; i < N ; i++ {
		chunk := rand_chunk(r, n_min, n_max)
		if !merkletree.CorrectRoot(roothash, chunk, list[i].Path()) {
			fmt.Println("One of the Merkle Paths did not check out!")
		}
	}
}

func main() {
	var seed int64
	flag.Int64Var(&seed, "seed", rand.Int63(), "Random seed for test")
	flag.Parse()

	fmt.Println("Seed", seed)
	run_test(seed, 8, 32, 900)

	r := rand.New(rand.NewSource(seed))
	a := rand_bytes(r, 32)
	fmt.Println("a: ", bytes_as_hex(a))
	b := rand_bytes(r, 32)
	fmt.Println("b: ", bytes_as_hex(b))
	
	gen := merkletree.NewMerkleTreeGen()
	gen.AddChunk(a, true)
	gen.AddChunk(b, true)

	fmt.Println(gen.List)
}
