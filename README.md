# Data-Structures & Algorithms [![Build Status](https://travis-ci.org/dansackett/algorithms.svg?branch=master)](https://travis-ci.org/dansackett/algorithms)

In an effort to learn more Golang and understand computer science concepts
better, I'm pairing the two in order to multiply my efforts. This repo's
design is based off that of [arnauddri](https://github.com/arnauddri/algorithms).
I merely took the structure and CS concepts as reference and ignored the
actual implementations.

## What I Want to Learn

* How to write Idiomatic Golang code
* How to write tests for Golang code
* How to implement common data structures
* Which algorithms are most efficient for sorting / searching

## Using this Project

I've provided a file to set the `$GOPATH` for your project in the
`bin/postactivate` hook. I use this with Python Virtualenvwrapper to handle
things, but you can just as easily run:

```
$ bash bin/postactivate
```

This will allow you to `go get ...` items from Github and have them locally.
Not sure if this is best practice as it's something I've struggled to
understand. Nevertheless, I've also provided a postdeactivate hook to unset
the variable.

```
$ bash bin/postdeactivate
```

Ideally you want to have these fire off when CDing into / out of your directory.
[Autoenv](https://github.com/kennethreitz/autoenv) is one solution for this,
but I prefer [Vrtualenv with Virtualenvwrapper](http://programeveryday.com/post/setting-up-a-python-development-environment-to-make-developing-a-breeze/).

## Table of Contents

### Data-Structures

* [Binary Search Tree](https://github.com/dansackett/algorithms/tree/master/data-structures/binary-tree) ([Wiki](http://en.wikipedia.org/wiki/Binary_search_tree))
* [Node](https://github.com/dansackett/algorithms/tree/master/data-structures/node) ([Wiki](http://en.wikipedia.org/wiki/Node_%28computer_science%29))
* Hash Tables
* Linked List
* Matrix
* Min/Max Heap
* Priority Queue
* Queue
* Stack
* Trie
* Graph

### Graph algorithms

**Searching:**
* Depth First Search
* Breadth First Search
* Kosaraju's Algorithm (find all SCCs)

**Shortest path:**
* Breadth First Search Shortest Path
* Dijkstra

**Sorting:**
* Topological Sort

### Maths algorithms

* Binary GCD algorithm
* Closest pairs
* FastPower
* Fibonacci
* Fisher-Yates Shuffle
* Erastothenes Sieve
* Extented GCD algorithm
* Karatsuba's Multiplication
* Newton's Square Root
* Permutations Count
* Strassen's matrix multiplication
* Randomized Selection

### Sorting algorithms

* Bubble Sort
* Heap Sort
* Quick Sort
* Merge Sort
* Insertion Sort
* Shell Sort
* Selection Sort

### Searching algorithms

* Binary Search
