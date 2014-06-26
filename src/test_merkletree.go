package main

import (
	"fmt"
	"flag"
	"math/rand"
	"merkletree"
	"crypto/sha256"
	"encoding/hex"
)

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

func to_bytes(x [sha256.Size]byte) []byte {
	ret := []byte{}
	for i := range x {
		ret = append(ret, x[i])
	}
	return ret
}

//Adds a lot of chunks and lists the tree leaves.
func run_test(seed int64, n_min int32, n_max int32, N int) {
	fmt.Println("Seed:", seed)
	r := rand.New(rand.NewSource(seed))

	gen := merkletree.NewMerkleTreeGen()  //Put chunks in.
	list := []*merkletree.MerkleNode{}
	for i:= 0 ; i < N ; i++ {
		chunk := rand_chunk(r, n_min, n_max)
		list = append(list, gen.AddChunk(chunk, true))
	}
	roothash := gen.Finish().Hash  //Get the root hash.
	fmt.Println("Root:", hex.EncodeToString(roothash[:]))

	fmt.Println("---")
//Reset random function, doing exact same to it.
	r = rand.New(rand.NewSource(seed))
	for i := 0 ; i < N ; i++ {
		chunk := rand_chunk(r, n_min, n_max)
		if !list[i].IsValid(-1) || !list[i].CorrespondsToChunk(chunk) {
			fmt.Println("Chunk %v didnt check out.", i)
		}
		
		path := list[i].Path()
//		root := merkletree.ExpectedRoot(merkletree.H(chunk), path)
//		fmt.Println(hex.EncodeToString(to_bytes(root)))

		if !merkletree.CorrectRoot(roothash, chunk, path) {
			fmt.Println(" - One of the Merkle Paths did not check out!")
		}
	}
	fmt.Println("---")
	fmt.Println("No messages above implies success.")
}

func main() {
	var seed int64
	flag.Int64Var(&seed, "seed", rand.Int63(), "Random seed for test.")
	var n_min int64
	flag.Int64Var(&n_min, "n_min", 1, "Minimum length of random chunk.")
	var n_max int64
	flag.Int64Var(&n_max, "n_max", 256, "Maximum length of random chunk.")
	var N int
	flag.IntVar(&N, "N", 256, "Number of chunks.")
	
	flag.Parse()

	run_test(seed, int32(n_min), int32(n_max), N)
}
