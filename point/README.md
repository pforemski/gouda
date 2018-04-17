# gouda/point

An implementation of various operations on generic n-dimensional points (and ranges of points).

## Features
* data sampling (e.g. random, uniform)
* quick median search (via sampling)
* various distance measures (e.g. Euclidean, Taxicab)
* basic statistics (e.g. standard deviation, sum of squared errors)
* auxiliary data attached to points

## Documentation

**API documentation**:
* See [godoc.org](https://godoc.org/github.com/pforemski/gouda/point)

Some basic examples (run with `go test -v`):
* creating points: [point_test.go](point_test.go)
* operations & stats: [math_test.go](math_test.go)
* distance measures: [distance_test.go](distance_test.go)

## Author

Copyright (C) 2018 by Pawel Foremski, [@pforemski](https://twitter.com/pforemski).

Licensed under GNU GPL v3. 
