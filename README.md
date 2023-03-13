# go-graph

Go graph is a graph database written in Go.

Its goals are:
1. Highly parallelisable
2. Transactional (including schema changes)
3. Cypher compliant
4. Batteries-included (e.g. migrations)

I have no formal experience with database design, so this is basically a way to
explore the concepts while bringing my own thoughts to the table. Sometimes it's
fun to do this kind of thing with little experience, a little bit like an
[outsider artist](https://en.wikipedia.org/wiki/Outsider_art).

# Roadmap

## In memory

The first goal is to design an appropriate in-memory data layout for graphs.

The result should be an API that allows for reads and writes to that in-memory store.

Out of scope:
1. Persistence to disk
2. Querying via Cypher (querying will only be possible by calling methods on the in memory data structures)
3. Graph traversal (inc. parallelism)

## Persistence

The second goal is to persist the graph state to disk. This is, after all,
arguably the most important part of a database.

The result should be an addition to the API that allows for persistence and
restoration from data on disk. It should be able to handle in-memory caching and
invalidation.

## Traversal

The third goal is to be able to perform traversals.

## Cypher

The fourth goal is to expose the graph via Cypher queries.