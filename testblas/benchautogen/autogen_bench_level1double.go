// Copyright 2014 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file

// Script for automatic code generation of the benchmark routines
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"strconv"
)

var output = flag.String("output", "", "output file name")

var copyrightnotice = `// Copyright 2014 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file`

var autogen = `// This file is autogenerated by github.com/gonum/blas/testblas/benchautogen/autogen_bench_level1double.go`

var imports = `import(
	"math/rand"
	"testing"
	"github.com/gonum/blas"
)`

var randomSliceFunction = `func randomSlice(l, idx int) ([]float64) {
	if idx < 0{
		idx = -idx
	}
	s := make([]float64, l * idx)
	for i := range s {
		s[i] = rand.Float64()
	}
	return s
}`

const (
	posInc1 = 5
	posInc2 = 3
	negInc1 = -3
	negInc2 = -4
)

var level1Sizes = []struct {
	lower string
	upper string
	camel string
	size  int
}{
	{
		lower: "small",
		upper: "SmallSlice",
		camel: "Small",
		size:  10,
	},
	{
		lower: "medium",
		upper: "MediumSlice",
		camel: "Medium",
		size:  1000,
	},
	{
		lower: "large",
		upper: "LargeSlice",
		camel: "Large",
		size:  100000,
	},
	{
		lower: "huge",
		upper: "HugeSlice",
		camel: "Huge",
		size:  10000000,
	},
}

type level1functionStruct struct {
	camel      string
	sig        string
	call       string
	extraSetup string
	oneInput   bool
	extraName  string // if have a couple different cases for the same function
}

var level1Functions = []level1functionStruct{
	{
		camel:    "Ddot",
		sig:      "n int, x []float64, incX int, y []float64, incY int",
		call:     "n, x, incX, y, incY",
		oneInput: false,
	},
	{
		camel:    "Dnrm2",
		sig:      "n int, x []float64, incX int",
		call:     "n, x, incX",
		oneInput: true,
	},
	{
		camel:    "Dasum",
		sig:      "n int, x []float64, incX int",
		call:     "n, x, incX",
		oneInput: true,
	},
	{
		camel:    "Idamax",
		sig:      "n int, x []float64, incX int",
		call:     "n, x, incX",
		oneInput: true,
	},
	{
		camel:    "Dswap",
		sig:      "n int, x []float64, incX int, y []float64, incY int",
		call:     "n, x, incX, y, incY",
		oneInput: false,
	},
	{
		camel:    "Dcopy",
		sig:      "n int, x []float64, incX int, y []float64, incY int",
		call:     "n, x, incX, y, incY",
		oneInput: false,
	},
	{
		camel:      "Daxpy",
		sig:        "n int, alpha float64, x []float64, incX int, y []float64, incY int",
		call:       "n, alpha, x, incX, y, incY",
		extraSetup: "alpha := 2.4",
		oneInput:   false,
	},
	{
		camel:      "Drot",
		sig:        "n int, x []float64, incX int, y []float64, incY int, c, s float64",
		call:       "n, x, incX, y, incY, c, s",
		extraSetup: "c := 0.89725836967\ns:= 0.44150585279",
		oneInput:   false,
	},
	{
		camel:      "Drotm",
		sig:        "n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams",
		call:       "n, x, incX, y, incY, p",
		extraSetup: "p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375,0}}",
		oneInput:   false,
		extraName:  "OffDia",
	},
	{
		camel:      "Drotm",
		sig:        "n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams",
		call:       "n, x, incX, y, incY, p",
		extraSetup: "p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}",
		oneInput:   false,
		extraName:  "Dia",
	},
	{
		camel:      "Drotm",
		sig:        "n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams",
		call:       "n, x, incX, y, incY, p",
		extraSetup: "p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}",
		oneInput:   false,
		extraName:  "Resc",
	},
	{
		camel:      "Dscal",
		sig:        "n int, alpha float64, x []float64, incX int",
		call:       "n, alpha, x, incX",
		extraSetup: "alpha := 2.4",
		oneInput:   true,
	},
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s pkgname\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	f := os.Stdout
	if *output != "" {
		var err error
		f, err = os.Create(*output)
		if err != nil {
			log.Fatalf("creating %s failed: %v\n", *output, err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatalf("closing %s failed: %v\n", *output, err)
			}
		}()
	}
	var b bytes.Buffer
	if err := level1(&b, args[0]); err != nil {
		log.Fatalf("generating level1 benchmark failed: %v\n", err)
	}
	src, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatalf("formatting source failed: %v\n", err)
	}
	f.Write(src)
}

func printHeader(f io.Writer, name string) error {
	if _, err := io.WriteString(f, copyrightnotice); err != nil {
		return err
	}
	io.WriteString(f, "\n\n")
	io.WriteString(f, autogen)
	io.WriteString(f, "\n\n")
	io.WriteString(f, "package "+name)
	io.WriteString(f, "\n\n")
	io.WriteString(f, imports)
	io.WriteString(f, "\n\n")
	return nil
}

// Generate the benchmark scripts for level1
func level1(f io.Writer, pkgname string) error {
	// Generate level 1 benchmarks
	printHeader(f, pkgname)

	// Print all of the constants
	io.WriteString(f, "const (\n")
	io.WriteString(f, "\tposInc1 = "+strconv.Itoa(posInc1)+"\n")
	io.WriteString(f, "\tposInc2 = "+strconv.Itoa(posInc2)+"\n")
	io.WriteString(f, "\tnegInc1 = "+strconv.Itoa(negInc1)+"\n")
	io.WriteString(f, "\tnegInc2 = "+strconv.Itoa(negInc2)+"\n")
	for _, con := range level1Sizes {
		io.WriteString(f, "\t"+con.upper+" = "+strconv.Itoa(con.size)+"\n")
	}
	io.WriteString(f, ")\n")
	io.WriteString(f, "\n")

	// Write the randomSlice function
	io.WriteString(f, randomSliceFunction)
	io.WriteString(f, "\n\n")

	// Start writing the benchmarks
	for _, fun := range level1Functions {
		writeLevel1Benchmark(fun, f)
		io.WriteString(f, "\n/* ------------------ */ \n")
	}

	return nil
}

func writeLevel1Benchmark(fun level1functionStruct, f io.Writer) {
	// First, write the base benchmark file
	io.WriteString(f, "func benchmark"+fun.camel+fun.extraName+"(b *testing.B, ")
	io.WriteString(f, fun.sig)
	io.WriteString(f, ") {\n")

	io.WriteString(f, "b.ResetTimer()\n")
	io.WriteString(f, "for i := 0; i < b.N; i++{\n")
	io.WriteString(f, "\tblasser."+fun.camel+"(")

	io.WriteString(f, fun.call)
	io.WriteString(f, ")\n}\n}\n")
	io.WriteString(f, "\n")

	// Write all of the benchmarks to call it
	for _, sz := range level1Sizes {
		lambda := func(incX, incY, name string, twoInput bool) {
			io.WriteString(f, "func Benchmark"+fun.camel+fun.extraName+sz.camel+name+"(b *testing.B){\n")
			io.WriteString(f, "n := "+sz.upper+"\n")
			io.WriteString(f, "incX := "+incX+"\n")
			io.WriteString(f, "x := randomSlice(n, incX)\n")
			if twoInput {
				io.WriteString(f, "incY := "+incY+"\n")
				io.WriteString(f, "y := randomSlice(n, incY)\n")
			}
			io.WriteString(f, fun.extraSetup+"\n")
			io.WriteString(f, "benchmark"+fun.camel+fun.extraName+"(b, "+fun.call+")\n")
			io.WriteString(f, "}\n\n")
		}
		if fun.oneInput {
			lambda("1", "", "UnitaryInc", false)
			lambda("posInc1", "", "PosInc", false)
		} else {
			lambda("1", "1", "BothUnitary", true)
			lambda("posInc1", "1", "IncUni", true)
			lambda("1", "negInc1", "UniInc", true)
			lambda("posInc1", "negInc1", "BothInc", true)
		}
	}
}
