package merkletree

//NOTE: untested.

import (
	"crypto/sha256"
)


func FirstBit(hash [sha256.Size]byte) bool {
	return hash[0]%2 == 1
}
func SetFirstBit(hash [sha256.Size]byte, to bool) [sha256.Size]byte {
	hash[0] -= hash[0]%2  // Always even.
	if to {
		hash[0] += 1
	}
	return hash
}

func H(a []byte) [sha256.Size]byte {
	return sha256.Sum256(a)
}
// TODO might want to add them and then sha instead, easier in EVM?
func H_2(a [sha256.Size]byte, b [sha256.Size]byte) [sha256.Size]byte {
	fa  := FirstBit(a)
	fb  := FirstBit(b)
	ret := sha256.Sum256(append(SetFirstBit(a, false), SetFirstBit(a, false)...))
	SetFirstBit(a, fa)
	SetFirstBit(b, fb)
	return ret
}


//NOTE: Hash64 is too few bits imo...

// A node of Merkle tree, note that the below omits a lot.
type MerkleNode struct {
	hash   [sha256.Size]byte
	left   *MerkleNode
	right  *MerkleNode
	up     *MerkleNode  // Otherwise it is messy to create the path.
}

func (self *MerkleNode) interest() bool {  // Even/odd is whether interest.
	return FirstBit(self.hash)
}

func new_MerkleNode(left *MerkleNode,right *MerkleNode) *MerkleNode {
	hash := SetFirstBit(H_2(left.hash, right.hash), left.interest() || right.interest())
	node := &MerkleNode{hash:hash, left:left, right:right, up:nil}
	left.up  = node
	right.up = node
	return node
}

func selective_new_MerkleNode(left *MerkleNode,right *MerkleNode) *MerkleNode {
	node := new_MerkleNode(left, right)
	node.left.effect_interest()
	node.right.effect_interest()
	return node
}

func (node MerkleNode) effect_interest() {
	if !node.interest() { //No interest in branches.
		node.left = nil
		node.right = nil
	}
}

type MerkleTreePortion struct { //Subtree in there somewhere.
	node   *MerkleNode
	depth  int
}


//The algo is to make a 'mountain range'(forgot link)
// and put the latest together if we have them. Meanwhile, it creates a tree, but
// drops anything in which there is no interest.
type MerkleTreeGen struct {
	list []MerkleTreePortion
}

// Adds chunk, returning the leaf the current is on.
func (gen MerkleTreeGen) AddChunk(chunk []byte, interest bool) *MerkleNode {
	h := SetFirstBit(H(chunk), interest)
	if len(gen.list) == 0 || gen.list[0].depth != 1 {
		add_node := &MerkleNode{hash:h, left:nil, right:nil, up:nil}

		gen.list = append(gen.list, MerkleTreePortion{node:add_node, depth:1})
		return add_node
	} else {
		// assert gen.list[0].depth == 1
		new_leaf := &MerkleNode{hash:h, left:nil, right:nil}

		new_node := selective_new_MerkleNode(new_leaf, gen.list[0].node)  //Combine the two.
		gen.list[0] = MerkleTreePortion{node:new_node, depth:2}

		// Combine more, while equal depth.
		for len(gen.list) >= 2 && gen.list[1].depth == gen.list[0].depth {
			new_node := selective_new_MerkleNode(gen.list[0].node, gen.list[1].node)
			gen.list = gen.list[1:] // Cut off the first one.
			gen.list[0] = MerkleTreePortion{node:new_node, depth:gen.list[0].depth + 1}
		}
		return new_leaf  //Return the leaf.
	}
}

// Coerce the last parts together, returning the root.
// NOTE: you can 'finish' and then continue to make what you put in already
// becomes a bit of a 'lob' that takes longer Merkle paths.
func (gen MerkleTreeGen) Finish() *MerkleNode {
	// assert len(gen.list) > 0
	for len(gen.list) >= 2  {
		new_node := selective_new_MerkleNode(gen.list[0].node, gen.list[1].node)
		gen.list = gen.list[1:]
		gen.list[0] = MerkleTreePortion{node:new_node, depth:gen.list[0].depth}
	}
	return gen.list[0].node
}

// Calculate the path, requires that you have calculated it up to the root.
// Note that the first bit that was used to indicate interest, now indicates
// it goes to the right.
func (node *MerkleNode) Path() [][sha256.Size]byte {
	if node.up == nil {
		var ret [][sha256.Size]byte
		return ret
	} else if node.up.right == node { //Going right.
		var hash [sha256.Size]byte
		copy(hash, node.up.left.hash)
		return append(node.up.Path(), SetFirstBit(hash, true))
	}	else if node.up.left == node { //Going left.
		var hash [sha256.Size]byte
		copy(hash, node.up.right.hash)
		return append(node.up.Path(), SetFirstBit(hash, true))
	}	else { // Information was not stored, or invalid Merkle tree.
		var ret [][sha256.Size]byte
		return ret
	}
}

//Calculate expected root, given the path.
func ExpectedRoot(H_leaf [sha256.Size]byte, path [][sha256.Size]byte) [sha256.Size]byte {
	x := H_leaf
	for i, h := range path {
		if FirstBit(h) {
			x = H_2(x, h)
		} else {
			x = H_2(h, x)
		}
	}
	return x
}

//Checks a root.
func CorrectRoot(root [sha256.Size]byte, leaf []byte, path [][sha256.Size]byte) bool {
	return ExpectedRoot(H(leaf), path) == root
}
