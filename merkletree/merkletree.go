package merkletree

//NOTE: untested.

import (
	"hash"
	"sha256"
)

//NOTE: Hash64 is too few bits imo...

type MerkleTreePortion struct { //Subtree in there somewhere.
	hash  [sha256.Size]byte
	n     int
}

// A node of Merkle tree, note that the below omits a lot.
type MerkleNode struct {
	hash   []byte
	left   *MerkleNode
	right  *MerkleNode
	interesting  bool
}

//The algo is to make a 'mountain range'(forgot link)
// and put the latest together if we have them. 
type MerkleTreeGen struct {
	list  []MerkleTreePortion
	tree  *MerkleNode
}

//Get the twig part of the merkle path of the _next_ one.
func (gen MerkleTreeGen) merkletwig() [][sha256.Size]byte {
	list := [][sha256.Size]byte
	for i, p := range gen.list {
		list = append(list, p.hash)
	}
	return list
}

func (gen MerkleTreeGen) addchunk(chunk []byte, interest bool) {
	h := sha256.Sum256(chunk)
	if len(gen.list) == 0 || gen.list[0].n != 1 {
		gen.list = append(gen.list, MerkleTreePortion{hash:h, n:1})
	} else { //Combine the first one.
		// assert gen.list[0].n == 1
		combined := append(h, gen.list[0].hash...)
		gen.list[0].hash = sha256.Sum256(combined)
		// And potentially combine more.
		for len(gen.list) >= 2 && gen.list[1].n == gen.list[0].n {
			combined := append(gen.list[0].hash, gen.list[1].hash...)
			gen.list = gen.list[1:] // Cut off the first one.
			gen.list[0].hash = sha256.Sum256(combined)  //(n is good).
		}
	}
}

//Coerce the last parts together.
func (gen MerkleTreeGen) finish(){
	//TODO
}


func gen_path(chunk []byte)
