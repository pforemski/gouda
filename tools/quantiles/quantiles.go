package main

import (
	"time"
	"math/rand"
	"fmt"
	"flag"
	"github.com/axiomhq/quantiles"
	"os"
	"strconv"
	"bufio"
	"sort"
)

var (
	optR    = flag.Int("P", 1000, "number of random points to generate (0=read stdin)")
	optQ    = flag.Int("Q", 10, "number of quantiles to generate")
	optM    = flag.Float64("M", 100.0, "max. random number")
	optPrint = flag.Bool("print", false, "print sorted numbers")
	optErr   = flag.Bool("err", false, "show estimation errors")
)
 
func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	Q := quantiles.NewDefault()
	var R sort.Float64Slice

	if *optR == 0 {
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			r, err := strconv.ParseFloat(stdin.Text(), 64)
			if err != nil { continue }
			Q.Push(r, 1.0)
			if *optErr || *optPrint { R = append(R, r) }
		}
	} else {
		for i := 0; i < *optR; i++ {
			r := rand.Float64() * *optM
			Q.Push(r, 1.0)
			if *optErr || *optPrint { R = append(R, r) }
		}
	}

	Q.Finalize()
	quants,err := Q.GenerateQuantiles(int64(*optQ))
	if err != nil { panic(err) }

	if *optPrint || *optErr { R.Sort() }
	if *optPrint { fmt.Printf("%g\n", R) }
	rstep := float64(len(R)) / float64(*optQ)

	qstep := 100.0 / float64(*optQ)
	for q := 0; q <= *optQ; q++ {
		if *optErr {
			ri := int(float64(q) * rstep)
			if ri >= len(R) { ri = len(R) - 1 }
			fmt.Printf("%g\t%.4f\t%g\n", float64(q)*qstep, quants[q], quants[q]-R[ri])			
		} else {
			fmt.Printf("%g\t%.4f\n", float64(q)*qstep, quants[q])
		}
	}
}