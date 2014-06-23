package merkletree

//NOTE: untested.

import (
	"hash"
	"sha256"
)

//NOTE: Hash64 is too few bits imo...

// A node of Merkle tree, note that the below omits a lot.
type MerkleNode struct {
	hash   [sha256.Size]byte
	left   *MerkleNode
	right  *MerkleNode
	up     *MerkleNode  // Otherwise it is messy to create the path.
	interest  bool
}

func new_MerkleNode(left *MerkleNode,right *MerkleNode) *MerkleNode {
	interest := left.interest || right.interest	
	hash = sha256.Sum256(append(left.hash, right.hash...))
	node = &MerkleNode{hash:hash, left:left, right:right, up:nil, interest:interest}
	left.up  = node
	right.up = node
	return node
}

func selective_new_MerkleNode(left *MerkleNode,right *MerkleNode) *MerkleNode {
	node = new_MerkleNode(left, right)
	node.left.effect_interest()
	node.right.effect_interest()
	return node
}

func (node MerkleNode) effect_interest()
{
	if !node.interest { //No interest in branches.
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
	list  []MerkleTreePortion
}

// Adds chunk, returning the leaf the current is on.
func (gen MerkleTreeGen) add_chunk(chunk []byte, interest bool) *MerkleNode {
	h := sha256.Sum256(chunk)
	if len(gen.list) == 0 || gen.list[0].depth != 1 {
		add_node := &MerkleNode{hash:h, left:nil, right:nil, up:nil, interest:false}
		gen.list = append(gen.list, MerkleTreePortion{node:add_node, n:1})
		return add_node
	} else {
		// assert gen.list[0].depth == 1
		new_leaf = MerkleNode{hash:h, left:nil, right:nil, interest:interest}

		new_node := selective_new_MerkleNode(new_leaf, gen.list[0].node)  //Combine the two.
		gen.list[0] = MerkleTreePortion{node:new_node, depth:2}

		// Combine more, while equal depth.
		for len(gen.list) >= 2 && gen.list[1].depth == gen.list[0].depth {
			new_node := selective_new_MerkleNode(gen.list[0], gen.list[1])
			gen.list = gen.list[1:] // Cut off the first one.
			gen.list[0] = new_node
		}
		return new_leaf  //Return the leaf.
	}
}

// Coerce the last parts together, returning the root.
func (gen MerkleTreeGen) finish() MerkleNode* {
	// assert len(gen.list) > 0
	for len(gen.list) >= 2  {
		new_node = selective_new_MerkleNode(gen.list[0], gen.list[1])
		gen.list = gen.list[1:]
		gen.list[0] = new_node
	}
	return gen.list[0].node
}

// Calculate the path, requires that you have calculated it up to the root.
func (node *MerkleNode) gen_path() [][sha256.Size]byte {
	if node.up != nil {
		return append(node.hash, node.up.gen_path()...)
	} else {
		var ret [][sha256.Size]byte = {}
		return ret
	}
}
