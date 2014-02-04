// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package blas provides interfaces for the BLAS linear algebra standard.

All methods must perform appropriate parameter checking and panic if
provided parameters that do not conform to the requirements specified
by the BLAS standard.

Quick Reference Guide to the BLAS from http://www.netlib.org/lapack/lug/node145.html

Level 1 BLAS

	        dim scalar vector   vector   scalars              5-element prefixes
	                                                          struct

	_rotg (                                      a, b )                D
	_rotmg(                              d1, d2, a, b )                D
	_rot  ( n,         x, incX, y, incY,               c, s )          D
	_rotm ( n,         x, incX, y, incY,                      param )  D
	_swap ( n,         x, incX, y, incY )                              D, Z
	_scal ( n,  alpha, x, incX )                                       D, Z, Zd
	_copy ( n,         x, incX, y, incY )                              D, Z
	_axpy ( n,  alpha, x, incX, y, incY )                              D, Z
	_dot  ( n,         x, incX, y, incY )                              D,
	_dotu ( n,         x, incX, y, incY )                              Z
	_dotc ( n,         x, incX, y, incY )                              Z
	_nrm2 ( n,         x, incX )                                       D, Dz
	_asum ( n,         x, incX )                                       D, Dz
	I_amax( n,         x, incX )                                       d, z

Level 2 BLAS

	        options                   dim   b-width scalar matrix  vector   scalar vector   prefixes

	_gemv ( order,        trans,      m, n,         alpha, a, lda, x, incX, beta,  y, incY ) D, Z
	_gbmv ( order,        trans,      m, n, kL, kU, alpha, a, lda, x, incX, beta,  y, incY ) D, Z
	_hemv ( order, uplo,                 n,         alpha, a, lda, x, incX, beta,  y, incY ) Z
	_hbmv ( order, uplo,                 n, k,      alpha, a, lda, x, incX, beta,  y, incY ) Z
	_hpmv ( order, uplo,                 n,         alpha, ap,     x, incX, beta,  y, incY ) Z
	_symv ( order, uplo,                 n,         alpha, a, lda, x, incX, beta,  y, incY ) D
	_sbmv ( order, uplo,                 n, k,      alpha, a, lda, x, incX, beta,  y, incY ) D
	_spmv ( order, uplo,                 n,         alpha, ap,     x, incX, beta,  y, incY ) D
	_trmv ( order, uplo, trans, diag,    n,                a, lda, x, incX )                 D, Z
	_tbmv ( order, uplo, trans, diag,    n, k,             a, lda, x, incX )                 D, Z
	_tpmv ( order, uplo, trans, diag,    n,                ap,     x, incX )                 D, Z
	_trsv ( order, uplo, trans, diag,    n,                a, lda, x, incX )                 D, Z
	_tbsv ( order, uplo, trans, diag,    n, k,             a, lda, x, incX )                 D, Z
	_tpsv ( order, uplo, trans, diag,    n,                ap,     x, incX )                 D, Z

	        options                   dim   scalar vector   vector   matrix  prefixes

	_ger  ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) D
	_geru ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) Z
	_gerc ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) Z
	_her  ( order, uplo,                 n, alpha, x, incX,          a, lda ) Z
	_hpr  ( order, uplo,                 n, alpha, x, incX,          ap )     Z
	_her2 ( order, uplo,                 n, alpha, x, incX, y, incY, a, lda ) Z
	_hpr2 ( order, uplo,                 n, alpha, x, incX, y, incY, ap )     Z
	_syr  ( order, uplo,                 n, alpha, x, incX,          a, lda ) D
	_spr  ( order, uplo,                 n, alpha, x, incX,          ap )     D
	_syr2 ( order, uplo,                 n, alpha, x, incX, y, incY, a, lda ) D
	_spr2 ( order, uplo,                 n, alpha, x, incX, y, incY, ap )     D

Level 3 BLAS

	        options                                 dim      scalar matrix  matrix  scalar matrix  prefixes

	_gemm ( order,             transA, transB,      m, n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) D, Z
	_symm ( order, side, uplo,                      m, n,    alpha, a, lda, b, ldb, beta,  c, ldc ) D, Z
	_hemm ( order, side, uplo,                      m, n,    alpha, a, lda, b, ldb, beta,  c, ldc ) Z
	_syrk ( order,       uplo, trans,                  n, k, alpha, a, lda,         beta,  c, ldc ) D, Z
	_herk ( order,       uplo, trans,                  n, k, alpha, a, lda,         beta,  c, ldc ) Z
	_syr2k( order,       uplo, trans,                  n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) D, Z
	_her2k( order,       uplo, trans,                  n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) Z
	_trmm ( order, side, uplo, transA,        diag, m, n,    alpha, a, lda, b, ldb )                D, Z
	_trsm ( order, side, uplo, transA,        diag, m, n,    alpha, a, lda, b, ldb )                D, Z

Meaning of prefixes

	D - float64	Z - complex128

Matrix types

	GE - GEneral 		GB - General Band
	SY - SYmmetric 		SB - Symmetric Band 	SP - Symmetric Packed
	HE - HErmitian 		HB - Hermitian Band 	HP - Hermitian Packed
	TR - TRiangular 	TB - Triangular Band 	TP - Triangular Packed

Options

	trans 	= NoTrans, Trans, ConjTrans
	uplo 	= Upper, Lower
	diag 	= Nonunit, Unit
	side 	= Left, Right (A or op(A) on the left, or A or op(A) on the right)

For real matrices, Trans and ConjTrans have the same meaning.
For Hermitian matrices, trans = Trans is not allowed.
For complex symmetric matrices, trans = ConjTrans is not allowed.
*/
package blas

// Flag constants indicate Givens transformation H matrix state.
type Flag int

const (
	Identity    Flag = iota - 2 // H is the identity matrix; no rotation is needed.
	Rescaling                   // H specifies rescaling.
	OffDiagonal                 // Off-diagonal elements of H are units.
	Diagonal                    // Diagonal elements of H are units.
)

// Type DrotmParams contains Givens transformation parameters returned
// by the Float64 Drotm method.
type DrotmParams struct {
	Flag
	H [4]float64 // Column-major 2 by 2 matrix.
}

// Type Order is used to specify the matrix storage format. An implementation
// may not implement both orders and must panic if a routine is called using
// an unimplemented order.
type Order int

const (
	RowMajor Order = 101 + iota
	ColMajor
)

// Type Transpose is used to specify the transposition operation for a
// routine.
type Transpose int

const (
	NoTrans Transpose = 111 + iota
	Trans
	ConjTrans
)

// Type Uplo is used to specify whether the matrix is an upper or lower
// triangular matrix.
type Uplo int

const (
	Upper Uplo = 121 + iota
	Lower
)

// Type Diag is used to specify whether the matrix is a unit or non-unit
// triangular matrix.
type Diag int

const (
	NonUnit Diag = 131 + iota
	Unit
)

// Type side is used to specify from which side a multiplication operation
// is performed.
type Side int

const (
	Left Side = 141 + iota
	Right
)

// Float64 implements the single precision real BLAS routines.
type Float64 interface {
	Float64Level1
	Float64Level2
	Float64Level3
}

// Float64Level1 implements the double precision real BLAS Level 1 routines.
type Float64Level1 interface {
	Ddot(n int, x []float64, incX int, y []float64, incY int) float64
	Dnrm2(n int, x []float64, incX int) float64
	Dasum(n int, x []float64, incX int) float64
	Idamax(n int, x []float64, incX int) int
	Dswap(n int, x []float64, incX int, y []float64, incY int)
	Dcopy(n int, x []float64, incX int, y []float64, incY int)
	Daxpy(n int, alpha float64, x []float64, incX int, y []float64, incY int)
	Drotg(a, b float64) (c, s, r, z float64)
	Drotmg(d1, d2, b1, b2 float64) (p DrotmParams, rd1, rd2, rb1 float64)
	Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64)
	Drotm(n int, x []float64, incX int, y []float64, incY int, p DrotmParams)
	Dscal(n int, alpha float64, x []float64, incX int)
}

// Float64Level2 implements the double precision real BLAS Level 2 routines.
type Float64Level2 interface {
	Dgemv(o Order, tA Transpose, m, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Dgbmv(o Order, tA Transpose, m, n, kL, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Dtrmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int)
	Dtbmv(o Order, ul Uplo, tA Transpose, d Diag, n, k int, a []float64, lda int, x []float64, incX int)
	Dtpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int)
	Dtrsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float64, lda int, x []float64, incX int)
	Dtbsv(o Order, ul Uplo, tA Transpose, d Diag, n, k int, a []float64, lda int, x []float64, incX int)
	Dtpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float64, x []float64, incX int)
	Dsymv(o Order, ul Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Dsbmv(o Order, ul Uplo, n, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
	Dspmv(o Order, ul Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int)
	Dger(o Order, m, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int)
	Dsyr(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int)
	Dspr(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, ap []float64)
	Dsyr2(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int)
	Dspr2(o Order, ul Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64)
}

// Float64Level3 implements the double precision real BLAS Level 3 routines.
type Float64Level3 interface {
	Dgemm(o Order, tA, tB Transpose, m, n, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Dsymm(o Order, s Side, ul Uplo, m, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Dsyrk(o Order, ul Uplo, t Transpose, n, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int)
	Dsyr2k(o Order, ul Uplo, t Transpose, n, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
	Dtrmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha float64, a []float64, lda int, b []float64, ldb int)
	Dtrsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha float64, a []float64, lda int, b []float64, ldb int)
}

// Complex128 implements the double precision complex BLAS routines.
type Complex128 interface {
	Complex128Level1
	Complex128Level2
	Complex128Level3
}

// Complex128Level1 implements the double precision complex BLAS Level 1 routines.
type Complex128Level1 interface {
	Zdotu(n int, x []complex128, incX int, y []complex128, incY int) (dotu complex128)
	Zdotc(n int, x []complex128, incX int, y []complex128, incY int) (dotc complex128)
	Dznrm2(n int, x []complex128, incX int) float64
	Dzasum(n int, x []complex128, incX int) float64
	Izamax(n int, x []complex128, incX int) int
	Zswap(n int, x []complex128, incX int, y []complex128, incY int)
	Zcopy(n int, x []complex128, incX int, y []complex128, incY int)
	Zaxpy(n int, alpha complex128, x []complex128, incX int, y []complex128, incY int)
	Zscal(n int, alpha complex128, x []complex128, incX int)
	Zdscal(n int, alpha float64, x []complex128, incX int)
}

// Complex128Level2 implements the double precision complex BLAS Level 2 routines.
type Complex128Level2 interface {
	Zgemv(o Order, tA Transpose, m, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Zgbmv(o Order, tA Transpose, m, n int, kL int, kU int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Ztrmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex128, lda int, x []complex128, incX int)
	Ztbmv(o Order, ul Uplo, tA Transpose, d Diag, n, k int, a []complex128, lda int, x []complex128, incX int)
	Ztpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex128, x []complex128, incX int)
	Ztrsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex128, lda int, x []complex128, incX int)
	Ztbsv(o Order, ul Uplo, tA Transpose, d Diag, n, k int, a []complex128, lda int, x []complex128, incX int)
	Ztpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex128, x []complex128, incX int)
	Zhemv(o Order, ul Uplo, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Zhbmv(o Order, ul Uplo, n, k int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Zhpmv(o Order, ul Uplo, n int, alpha complex128, ap []complex128, x []complex128, incX int, beta complex128, y []complex128, incY int)
	Zgeru(o Order, m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int)
	Zgerc(o Order, m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int)
	Zher(o Order, ul Uplo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int)
	Zhpr(o Order, ul Uplo, n int, alpha float64, x []complex128, incX int, a []complex128)
	Zher2(o Order, ul Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int)
	Zhpr2(o Order, ul Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, ap []complex128)
}

// Complex128Level3 implements the double precision complex BLAS Level 3 routines.
type Complex128Level3 interface {
	Zgemm(o Order, tA, tB Transpose, m, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
	Zsymm(o Order, s Side, ul Uplo, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
	Zsyrk(o Order, ul Uplo, t Transpose, n, k int, alpha complex128, a []complex128, lda int, beta complex128, c []complex128, ldc int)
	Zsyr2k(o Order, ul Uplo, t Transpose, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
	Ztrmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int)
	Ztrsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int)
	Zhemm(o Order, s Side, ul Uplo, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
	Zherk(o Order, ul Uplo, t Transpose, n, k int, alpha float64, a []complex128, lda int, beta float64, c []complex128, ldc int)
	Zher2k(o Order, ul Uplo, t Transpose, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta float64, c []complex128, ldc int)
}
