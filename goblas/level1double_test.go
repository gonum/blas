// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Autogenerate benchmark code
//go:generate go run ../testblas/benchautogen/autogen_bench_level1double.go -output level1doubleBench_auto_test.go goblas

package goblas

import (
	"testing"

	"github.com/gonum/blas/testblas"
)

var blasser = Blas{}

func TestDasum(t *testing.T) {
	testblas.DasumTest(t, blasser)
}

func TestDaxpy(t *testing.T) {
	testblas.DaxpyTest(t, blasser)
}

func TestDdot(t *testing.T) {
	testblas.DdotTest(t, blasser)
}

func TestDnrm2(t *testing.T) {
	testblas.Dnrm2Test(t, blasser)
}

func TestIdamax(t *testing.T) {
	testblas.IdamaxTest(t, blasser)
}

func TestDswap(t *testing.T) {
	testblas.DswapTest(t, blasser)
}

func TestDcopy(t *testing.T) {
	testblas.DcopyTest(t, blasser)
}

func TestDrotg(t *testing.T) {
	testblas.DrotgTest(t, blasser)
}

func TestDrotmg(t *testing.T) {
	testblas.DrotmgTest(t, blasser)
}

func TestDrot(t *testing.T) {
	testblas.DrotTest(t, blasser)
}

func TestDrotm(t *testing.T) {
	testblas.DrotmTest(t, blasser)
}

func TestDscal(t *testing.T) {
	testblas.DscalTest(t, blasser)
}
