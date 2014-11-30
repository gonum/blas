package testblas

import (
	"fmt"
	"testing"

	"github.com/gonum/blas"
)

type Dtbsver interface {
	Dtbsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n, k int, a []float64, lda int, x []float64, incX int)
}

func DtbsvTest(t *testing.T, blasser Dtbsver) {
	for i, test := range []struct {
		ul   blas.Uplo
		tA   blas.Transpose
		d    blas.Diag
		n, k int
		a    [][]float64
		lda  int
		x    []float64
		incX int
		ans  []float64
	}{
		{
			ul: blas.Upper,
			tA: blas.NoTrans,
			d:  blas.NonUnit,
			n:  5,
			k:  1,
			a: [][]float64{
				{1, 3, 0, 0, 0},
				{0, 6, 7, 0, 0},
				{0, 0, 2, 1, 0},
				{0, 0, 0, 12, 3},
				{0, 0, 0, 0, -1},
			},
			x:    []float64{1, 2, 3, 4, 5},
			incX: 1,
			ans:  []float64{2.479166666666667, -0.493055555555556, 0.708333333333333, 1.583333333333333, -5.000000000000000},
		},
	} {
		aFlat := flattenTriBanded(test.a, test.k, test.ul)
		//aFlat = []float64{1, 6, 2, 12, -1, 0, 3, 7, 1, 3}
		fmt.Println(aFlat)
		xCopy := sliceCopy(test.x)

		// TODO: Have tests where the banded matrix is constructed explicitly
		// to allow testing for lda =! k+1
		blasser.Dtbsv(test.ul, test.tA, test.d, test.n, test.k, aFlat, test.k+1, xCopy, test.incX)
		if !dSliceTolEqual(test.ans, xCopy) {
			t.Errorf("Case %v: Want %v, got %v", i, test.ans, xCopy)
		}
	}
}
