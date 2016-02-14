# Graph

A graph is a container which holds vertices or nodes. These vertices can be
connected by different edges and when we have these edges, we can traverse
a graph finding the shortest path to different vertices as well as creating
lines between vertices. Graphs can be directed or undirected. The
difference being that a directed graph has edges that go one way and
undirected graphs can go both backwards and forwards along the edge.

In the following code, I have implemented a undirected graph. As a handy
way to view the graph and use it, I have provided two functions to return
it:

- `ToAdjacencyList()`
- `ToAdjacencyMatrix()`

An Adjacency list has vertices stored as records or objects, and every
vertex stores a list of adjacent vertices. This data structure allows the
storage of additional data on the vertices. Additional data can be stored
if edges are also stored as objects, in which case each vertex stores its
incident edges and each edge stores its incident vertices. An example is:

```
[1]
[0 2 3]
[1]
[1]
```

An Adjacency Matrix is a two-dimensional matrix, in which the rows represent
source vertices and columns represent destination vertices. Data on edges
and vertices must be stored externally. Only the cost for one edge can be
stored between each pair of vertices. An example is:


```
[0 1 0 0]
[1 0 1 1]
[0 1 0 0]
[0 1 0 0]
```

Adjacency lists are generally preferred because they efficiently represent
sparse graphs. An adjacency matrix is preferred if the graph is dense, that
is the number of edges is close to the number of vertices squared or if one
must be able to quickly look up if there is an edge connecting two vertices.

## Time Complexity

| Operations    | Adjacency List    | Adjacency Matrix  |
| ------------- |:-----------------:|:-----------------:|
| Store Graph   | `O(|V|+|E|)`      | `O(|V|^2)`        |
| Add Vertex    | `O(1)`            | `O(|V|^2)`        |
| Add Edge      | `O(1)`            | `O(1)`            |
| Remove Vertex | `O(|E|)`          | `O(|V|^2)`        |
| Remove Edge   | `O(|E|)`          | `O(1)`            |
