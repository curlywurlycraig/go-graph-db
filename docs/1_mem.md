# Part 1: In Memory

The first goal is to design an appropriate in-memory data layout for graphs.

The result should be an API that allows for reads and writes to that in-memory store.

Out of scope:
1. Persistence to disk
2. Querying via Cypher (querying will only be possible by calling methods on the in memory data structures)
3. Graph traversal (inc. parallelism)

## Basics

Our graph database consists of nodes and edges, like all graph database. There
is a distinction between a node _type_ and a node _instance_. A bit like a table
and a row. The same is true for edges.

A node type has a _name_, and a _definition_. The definition of a node is its
attributes and their types. The same is true for edges.

A node instance has a unique ID, and values for each of the attributes.

An edge has a unique ID, values for each of its attributes, a _from_ ID and a
_to_ ID. The _from_ and _to_ IDs are of course the IDs of the nodes the edge
connects together.

Therefore, a traversal means finding edge types where the _from_ ID matches the
"current" node ID.

## Possible memory layouts

Using the _structure of arrays_ pattern, we could imagine an array for each
_node type_ and each _edge type_.

Adding a new node or edge _type_ means adding a new entry in the structure.
Adding a new node or edge _instance_ means inserting into one of those arrays.

## TODO

- [ ] Add some more tests and functionality
- [ ] Add some benchmarks
- [ ] Figure out O(1) lookup on node edges