##
##  Copyright (C) 11-06-2014 Jasper den Ouden.
##
##  This is free software: you can redistribute it and/or modify
##  it under the terms of the GNU General Public License as published
##  by the Free Software Foundation, either version 3 of the License, or
##  (at your option) any later version.
##

from pathfind import PathedNode


class RepNode(PathedNode):

    def cost(self, to, i, goal):
        assert i < len(self.data)
    # The higher the trust in a direction, the better the path.
        return self.context.fun(self.data[i])


def unzip(pairs):  # Unzips a list.
    r = [], []
    for el in pairs:
        r[0].append(el[0])
        r[1].append(el[1])
    return r


# In simulation/looking at the data fields, 'asks' who a reputation contract in
# ethereum reputes.
class EthereumContext:
    known_nodes = {}

    # Function that governs the effect of setting the reputation param.
    def fun(self, rep):
        assert rep > 0
        return 1/(rep + 1)

    # Makes a node. Note that the reputations of the node is only obtained later.
    def figure_addr(self, addr):
        if addr not in self.known_nodes:
            self.known_nodes[addr] = RepNode(addr)
        return self.known_nodes[addr]

    def figure_edges(self, node):
        # TODO simulate a call that asks for all the reputations in the return value.

        repdata, addrs = unzip(reputations)
        return repdata, map(self.figure_addr, addrs)

# TODO next step is to gather low cost paths, multiple if needed.
# 
