package testblas

import (
	"math"
	"testing"

	"github.com/gonum/blas"
)

// throwPanic will throw unexpected panics if true, or will just report them as errors if false
const throwPanic = true

func dTolEqual(a, b float64) bool {
	m := math.Max(math.Abs(a), math.Abs(b))
	if m > 1 {
		a /= m
		b /= m
	}
	if math.Abs(a-b) < 1e-14 {
		return true
	}
	return false
}

func dSliceTolEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !dTolEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func dStridedSliceTolEqual(n int, a []float64, inca int, b []float64, incb int) bool {
	ia := 0
	ib := 0
	if inca <= 0 {
		ia = -(n - 1) * inca
	}
	if incb <= 0 {
		ib = -(n - 1) * incb
	}
	for i := 0; i < n; i++ {
		if !dTolEqual(a[ia], b[ib]) {
			return false
		}
		ia += inca
		ib += incb
	}
	return true
}

func dSliceEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !(a[i] == b[i]) {
			return false
		}
	}
	return true
}

func dCopyTwoTmp(x, xTmp, y, yTmp []float64) {
	if len(x) != len(xTmp) {
		panic("x size mismatch")
	}
	if len(y) != len(yTmp) {
		panic("y size mismatch")
	}
	for i, val := range x {
		xTmp[i] = val
	}
	for i, val := range y {
		yTmp[i] = val
	}
}

// returns true if the function panics
func panics(f func()) (b bool) {
	defer func() {
		err := recover()
		if err != nil {
			b = true
		}
	}()
	f()
	return
}

func testpanics(f func(), name string, t *testing.T) {
	b := panics(f)
	if !b {
		t.Errorf("%v should panic and does not", name)
	}
}

func sliceOfSliceCopy(a [][]float64) [][]float64 {
	n := make([][]float64, len(a))
	for i := range a {
		n[i] = make([]float64, len(a[i]))
		copy(n[i], a[i])
	}
	return n
}

func sliceCopy(a []float64) []float64 {
	n := make([]float64, len(a))
	copy(n, a)
	return n
}

func flatten(a [][]float64) []float64 {
	if len(a) == 0 {
		return nil
	}
	m := len(a)
	n := len(a[0])
	s := make([]float64, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s[i*n+j] = a[i][j]
		}
	}
	return s
}

func unflatten(a []float64, m, n int) [][]float64 {
	s := make([][]float64, m)
	for i := 0; i < m; i++ {
		s[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			s[i][j] = a[i*n+j]
		}
	}
	return s
}

func flattenTriBanded(a [][]float64, k int, ul blas.Uplo) []float64 {
	n := len(a)
	if n != len(a[0]) {
		panic("tri banded must be square")
	}
	aflat := make([]float64, (k+1)*n) // a is an lda x n matrix
	if ul == blas.Upper {
		for j := 0; j < n; j++ {
			m := k - j
			min := j - k
			if j-k < 0 {
				min = 0
			}
			for i := min; i < j+1; i++ {
				aflat[(m+i)*n+j] = a[i][j]
			}
		}
		return aflat
	}
	panic("not coded for lower")
}
