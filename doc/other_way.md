## Other ways to do it?
Maybe an approach like
'[hanging blocks](http://o-jasper.github.io/blog/2014/06/03/hanging_blocks.html)'
but where the entirety of the block can be calculated from Ethereum state,
making availability not an issue.

It is an entire record of all calculated implies trusts/reputations, and the
creator will be punished if any is wrong. A requester will just get a bit of the
file as proven to be a bit by the Merkle tree mechanism. He can recalculate it,
and if it doesnt check out report to the contract for revocation.
(Computation power to check depends on reputation algorithm, since it only needs
a single refutation, it can be more expensive on the blockchain per-time used)

For other contracts, they just accept it if it is in an old enough non-retracted
record. (contract writers themselves decide the threshhold. I would go for a day
for many purposes.)

### Efficiency
The number of interconnections is **N^2** of course. I think it is easiest to
chunk in parts of **M &le;1024** participants, and then each person has a long
list of implied trusts, in decending order. For 10^6 people..Another 20
splittings.. Each checksum being 32 bytes.. That is ~1kB for showing a single
reputation.

Low amounts of trust are possibly not very useful, the lower end could be omitted,
or the lowness determined by a small payment. The probability of use could also
be tied to reputation, and be used to create a 'splitting', that makes some of the
lower-probability options more expensive to access in exchange of the
higher-probability ones.

### Testing by the contract managing it
The contract has to be able to check it. For this the state of opinions  when it
was created is necessary. The whole thing will probably require 'storing' what
all values with implied `block.number % 10 == 0` for some amount of blocks, 
and 'calculated hanging blocks' are created for those. Older ones are obsolete.
