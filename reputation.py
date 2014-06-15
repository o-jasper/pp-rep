##
##  Copyright (C) 11-06-2014 Jasper den Ouden.
##
##  This is free software: you can redistribute it and/or modify
##  it under the terms of the GNU General Public License as published
##  by the Free Software Foundation, either version 3 of the License, or
##  (at your option) any later version.
##

import pathfind
from pathfind import PathedNode


class RepNode(PathedNode):

    def cost(self, at_cost, to, i, goal):
        assert i < len(self.data)
        assert self.data[i] > 0
    # The higher the trust in a direction, the better the path.
        if self.run_i < self.context.from_run_i:
            return at_cost + 1/(1 + self.data[i])
        else:  # When already used path, take it as a maximum.
            return at_cost + 1

    def multi_pathfind(self, goal, n):
        global pathing.g_run_i  # Move it on further.
        self.context.from_run_i = pathing.g_run_i + 1

        # NOTE/TODO this simply tries it many times.
        # Do we want to search from middles instead?
        return map(lambda(i): self.pathfind(goal), range(n))


def unzip(pairs):  # Unzips a list.
    r = [], []
    for el in pairs:
        r[0].append(el[0])
        r[1].append(el[1])
    return r


# In simulation/looking at the data fields, 'asks' who a reputation contract in
# ethereum reputes.
class EthereumContext:
    from_run_i  = 0
    known_nodes = {}

    # Makes a node. Note that the reputations of the node is only obtained later.
    def figure_addr(self, addr):
        if addr not in self.known_nodes:
            self.known_nodes[addr] = RepNode(addr)
        return self.known_nodes[addr]

    # Finds information in the Ethereum state.
    def figure_edges(self, node):
        # TODO simulate a call that asks for all the reputations in the return value.

        repdata, addrs = unzip(reputations)
        return repdata, map(self.figure_addr, addrs)
        
        
# TODO next step is to gather low cost paths, multiple if needed.
# 
