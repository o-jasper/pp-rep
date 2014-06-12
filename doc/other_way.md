## Other ways to do it?
Maybe an approach like
'[hanging blocks](http://o-jasper.github.io/blog/2014/06/03/hanging_blocks.html)'
but where the entirety of the block can be calculated from Ethereum state,
making availability not an issue.

It is an entire record of all calculated implies trusts/reputations, and the creator
will be punished if any is wrong. A requester will just get a bit of the file as
proven to be a bit by the Merkle tree mechanism. He can recalculate it, and if it
doesnt check out report to the contract for revocation.
(Computation power to check depends on reputation algorithm, since it only needs a
single refutation, it can be more expensive on the blockchain)

For other contracts, they just accept it if it is in an old enough non-retracted
record. (contract writers themselves decide the threshhold. I would go for a day
for many purposes.)

### Problem
the 'dumb' approach making a giant table requires &propto;n^2 items, which is
clearly prohibitive.
