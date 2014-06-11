##
##  Copyright (C) 11-06-2014 Jasper den Ouden.
##
##  This is free software: you can redistribute it and/or modify
##  it under the terms of the GNU General Public License as published
##  by the Free Software Foundation, either version 3 of the License, or
##  (at your option) any later version.
##

# Generic pathfind.

from heapq import heappush, heappop  # Heap, used as priority list.

g_run_i = 0  # TODO when it wraps, all has to be reset!


class PathedNode:  # Dijkstras algorithm for finding paths.
    def __init__(self, addr, edges=None, context=None, data=None):
        self.addr = addr
        self.data = data
        self.context = context
        self.edges = edges
        self.costval = 0  # Values used for the pathfinding.
        self.run_i   = 0

    def cost(self, at_cost, to, i, goal):  # Cost between two positions.
        return at_cost + 1

    def step_from(self, heap, at_cost, run_i, goal):
        self.costval = at_cost
        self.run_i   = run_i
        if self in goal:  # Found it, return it!
            return self
        if self.edges is None:
            self.data, self.edges = self.context.figure_edges(self)

        i = 0
        for el in self.edges:
            i += 1
            if el.run_i != run_i:  # Add if wasnt already there.
                heappush(heap, (self.cost(at_cost, el, i, goal), el))
            elif el in goal:
                return el
        return None

    # Pathfind from the node itself to the goal.
    def pathfind_prep(self, goal, run_i):
        heap = []
        got = self.step_from(heap, 0, run_i, goal)
        while got is None:  # Keep picking the lowest total cost.
            try:
                cost, cur = heappop(heap)
            except IndexError:
                return False  # Could not find a path.
            got = cur.step_from(heap, cost, run_i, goal)
        return True

    def pathfind_track(self, goal, run_i):
        cur = self
        path = []
        while cur not in goal:
            first = True
            min_costval, nxt = 0, None
            for el in cur.edges:
                if el.run_i == run_i and (el.costval < min_costval or first):
                    min_costval = el.costval
                    nxt = el
            path.append(cur)
            cur = nxt
            assert cur is not None  # Otherwise pathfind_prep returns False.
        path.append(cur)
        return path

 # Pathfinds from self to goal, returns None if no path, the backtrack otherwise.
 # goal.costval will contain the amount it costs.
    def pathfind(self, goal, run_i=None):
        if run_i is None:  # Note all nodes will have to be reset at some point.
            global g_run_i
            g_run_i += 1
            run_i = g_run_i
        if self.pathfind_prep(goal, run_i):
            return self.pathfind_track(goal, run_i)
