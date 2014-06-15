## Other ways to do it?
Maybe an approach like
'[hanging blocks](http://o-jasper.github.io/blog/2014/06/03/hanging_blocks.html)'
but where the entirety of the block can be calculated from Ethereum state.
Initially, i was fooled into thinking this fixed the availabilty-of-the-block
problem, but unfortunately, with incorrect data, you can still not produce
Merkle paths about it, so you cannot refute.

Anyway, it is an entire record of all calculated implies trusts/reputations, and
the creator will be punished if any is wrong. A requester will just get a 
Merkle-tree-proven-chunk of the file. He can recalculate it, and if it doesnt check
out report to the contract for revocation of the entire portion of data under 
that Merkle root.
(Computation power to check depends on reputation algorithm, since it only needs
a single refutation, any particular calculation can be more expensive)

For other contracts, they just accept it if it is in an old enough non-retracted
record. (contract writers themselves decide the threshhold. I would go for a day
for many purposes.)

### Efficiency
The number of interconnections is **N^2** of course. I think it is easiest to
chunk in parts of **M** participants, and then each person has a long
list of implied trusts, in decending order. For 10^6 people..Another 20
splittings.. Each checksum being 32 bytes.. That is ~1kB for showing a single
reputation between two persons.

Low amounts of trust are possibly not very useful, the lower end could be omitted,
or the lowness determined by a small payment. The probability of use could also
be tied to reputation, and be used to create a 'splitting', that makes some of the
lower-probability options more expensive to access in exchange of the
higher-probability ones.

Can try think about what good values of **M** are, frequently used addresses
want to be on lower **M** entries than infrequent ones. Higher **M** have the
advantage that refutation can be made more profitable.

### Refutation to the contract managing it
The contract has to be able to check it. For this the state of opinions  when it
was created is necessary. The whole thing will probably require 'storing' what
all values with implied `block.number % 10 == 0` for some amount of blocks, 
and 'calculated hanging blocks' are created for those. Older ones are obsolete.

### Availability problem
As said, it still suffers from the problem that you cannot refute if you do not
have the data, because you cannot contruct Merkle paths in that case. You could
take every time someone uses the contract to get at a reputation link as a vote
for that version. But that only mildly improves on the voting mechanism,
because we dont know if the mayor users may be the abusive ones..
