package cblas

import (
	"testing"

	"github.com/gonum/blas/testblas"
)

func TestDgemv(t *testing.T) {
	testblas.DgemvTest(t, blasser)
}

func TestDger(t *testing.T) {
	testblas.DgerTest(t, blasser)
}

func TestDtxmv(t *testing.T) {
	testblas.DtxmvTest(t, blasser)
}

func TestDtbsv(t *testing.T) {
	testblas.DtbsvTest(t, blasser)
}
