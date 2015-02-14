// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package native

import (
	"testing"

	testblas "github.com/gonum/blas/testblas/single"
)

func TestSasum(t *testing.T) {
	testblas.DasumTest(t, impl)
}

func TestSaxpy(t *testing.T) {
	testblas.DaxpyTest(t, impl)
}

func TestSdot(t *testing.T) {
	testblas.DdotTest(t, impl)
}

func TestSnrm2(t *testing.T) {
	testblas.Dnrm2Test(t, impl)
}

func TestIsamax(t *testing.T) {
	testblas.IdamaxTest(t, impl)
}

func TestSswap(t *testing.T) {
	testblas.DswapTest(t, impl)
}

func TestScopy(t *testing.T) {
	testblas.DcopyTest(t, impl)
}

func TestSrotg(t *testing.T) {
	testblas.DrotgTest(t, impl)
}

func TestSrotmg(t *testing.T) {
	testblas.DrotmgTest(t, impl)
}

func TestSrot(t *testing.T) {
	testblas.DrotTest(t, impl)
}

func TestSrotm(t *testing.T) {
	testblas.DrotmTest(t, impl)
}

func TestSscal(t *testing.T) {
	testblas.DscalTest(t, impl)
}
