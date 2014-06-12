
# Per-participant Reputation system approach for Ethereum

This project is to help figure out per-participant reputation systems. I.e. each
participant indicates its own opinion about what the reputation of other nodes
should be, and gets implied reputations from it.

It should be possible to 'prove' to a contract that there is reputation from one
to another. 

This is essentially pathfinding through opinions of the participants, though
it is a superset, because different paths are complementary. This pathfinding is
done client-side and the paths are put in the message for checking.
