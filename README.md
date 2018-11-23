# Gouda: Golang Utilities for Data Analysis

A collection of Golang libraries implementing various techniques for data analysis, including machine learning.

To install:
```
go get github.com/pforemski/gouda
```

This is work in progress. Expect breaking changes. [Embrace for impact](https://en.wikipedia.org/wiki/Sully_(film)).

## Features

Currently, it includes the following modules:

* [point](./point/): for operations on generic, n-dimensional points
 
* [kdtree](./kdtree/): a [kd-tree](https://en.wikipedia.org/wiki/K-d_tree) implementation

* data clustering:
  * [kmeans](./kmeans/): the
  [k-means](https://en.wikipedia.org/wiki/K-means_clustering) data clustering algorithm
  * [dbscan](./dbscan/): the [DBSCAN](https://en.wikipedia.org/wiki/DBSCAN)
  data clustering algorithm using kd-tree

* [interpolation](./interpolation/): numerical [interpolation](https://en.wikipedia.org/wiki/Interpolation)
  methods (wip, currently just the Lagrange polynomial)

## Documentation

**API documentation**:
* See [godoc.org](https://godoc.org/github.com/pforemski/gouda) for the root of API documentation. Navigate to module directories for detailed information.

More documentation & some examples available in the README files of each module:
* [point](./point/README.md)
* [kdtree](./kdtree/README.md)
* [kmeans](./kmeans/README.md)
* [dbscan](./dbscan/README.md)

## Author

Copyright (C) 2018 by Pawel Foremski, [@pforemski](https://twitter.com/pforemski).

Licensed under GNU GPL v3.
