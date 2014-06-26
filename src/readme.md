## Idea: merkle tree maker, and a contract that verifies them
The merkle tree should eat 'chunks' and put them on the leaves. It is intended
to be able to mark chunks as interesting, so it doesnt have to actually store
the whole thing.

This is useful for where large portions of the merkle tree are infact simply
calculated entries.

The Ethereum contract merely has to check.

#### Uses

* Lots-a-stuff, i dont know well about, torrents use them for instance.

* Allowing Ethereum to have access to data. Also 'pre-emptively' for data that
  you dont think Ethereum contracts have any use for.

* [Hanging blocks](http://o-jasper.github.io/blog/2014/06/03/hanging_blocks.html)
  of various sorts.

* [Dropbox example](https://github.com/jorisbontje/cll-sim/blob/master/examples/decentralized-dropbox.cll), i suppose.

#### Additional
An idea is to take probabilities for chunks and make a somewhat lobsided tree
that minimizes the average length of proving merkle paths. The checking mechanism
doesnt care about the shape of the merkle tree, so it can be added later,
also, `merkletree.Finish` can sort-of be used prematurely for lobsidedness.

On the other hand, off-chain, the efficiency in memory is better if things you want
to keep track of are bunched together.
