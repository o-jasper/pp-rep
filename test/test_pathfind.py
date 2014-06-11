import os.path
import sys

sys.path.append(os.path.join(os.path.dirname(__file__), '..'))

from pathfind import PathedNode

def connect_chain(to_chain):
    for i in range(len(to_chain)-1):
        to_chain[i].edges.append(to_chain[i+1])
    return to_chain

def make_chain(n):
    return connect_chain(map(lambda(x): PathedNode(x,[]), range(n)))

def test_chain(n):
    chain = make_chain(n)
    for k2 in range(n - 1):  #Overly exhaust stupidly identical possibilities.
        for k1 in range(n - k2):
            path = chain[k1].pathfind([chain[-k2-1]])
            assert len(path) == n - k2 - k1
   # Cannot go backward.
    assert chain[-1].pathfind([chain[0]]) is None

test_chain(20)

def ensure_node(name, nodes):
    if name not in nodes:
        nodes[name] = PathedNode(name, [])
    return nodes[name]

#Creates netowrk by bunches of cains.
def create_network(lists, nodes=None):
    nodes = nodes or {}
    for chain in lists:
        for i in range(len(chain)-1):
            cur = ensure_node(chain[i], nodes)
            nxt = ensure_node(chain[i+1], nodes)
            if nxt not in cur.edges:
                cur.edges.append(nxt)
    return nodes

def listnames(nodes):
    return map(lambda(el): el.addr, nodes)

net = create_network([[1,2,3,4,5,6,7,8,9], [3,'a', 'b', 7]])
# Thou shalt take the shortcut.
assert listnames(net[1].pathfind([net[9]])) == [1,2,3,'a','b',7,8,9]

