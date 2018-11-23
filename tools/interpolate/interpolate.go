package main

import (
	"math"
	"fmt"
	"strconv"
	"strings"
	"flag"
	"github.com/pforemski/gouda/point"
	"github.com/pforemski/gouda/interpolate"
)

var (
	optPoints = flag.Int("points", 100, "number of points to generate")
)
 
func main() {
	flag.Parse()

	var points point.Points
	for _,arg := range flag.Args() {
		d := strings.SplitN(arg, "=", 2)
		if len(d) != 2 { continue }

		x, err := strconv.ParseFloat(d[0], 64)
		if err != nil { panic(err) }

		y, err := strconv.ParseFloat(d[1], 64)
		if err != nil { panic(err) }

		points = append(points, point.New(x, y))
	}

	//fmt.Printf("points: %s\n", points)

	ip, err := interpolate.NewLagrange(points)
	if err != nil { panic(err) }

	i := 0
	firstx := points[i].V[0]
	lastx  := points[len(points)-1].V[0]
	step   := (lastx - firstx) / float64(*optPoints)
	for x := firstx; x < lastx; x += step {
		idiff := math.Abs(x - points[i].V[0])
		if x >= points[i].V[0] {
			fmt.Printf("%g\t%g\n", points[i].V[0], ip.Interpolate(points[i].V[0]))
			i++
		}
		if idiff > step/2 {
			fmt.Printf("%g\t%g\n", x, ip.Interpolate(x))
		}
	}
	fmt.Printf("%g\t%g\n", lastx, ip.Interpolate(lastx))
}