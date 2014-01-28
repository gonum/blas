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

	_rotg (                                      a, b )                S
	_rotmg(                              d1, d2, a, b )                S
	_rot  ( n,         x, incX, y, incY,               c, s )          S
	_rotm ( n,         x, incX, y, incY,                      param )  S
	_swap ( n,         x, incX, y, incY )                              S, C
	_scal ( n,  alpha, x, incX )                                       S, C, Cs
	_copy ( n,         x, incX, y, incY )                              S, C
	_axpy ( n,  alpha, x, incX, y, incY )                              S, C
	_dot  ( n,         x, incX, y, incY )                              S, Ds
	_dotu ( n,         x, incX, y, incY )                              C
	_dotc ( n,         x, incX, y, incY )                              C
	__dot ( n,  alpha, x, incX, y, incY )                              Sds
	_nrm2 ( n,         x, incX )                                       S, Sc
	_asum ( n,         x, incX )                                       S, Sc
	I_amax( n,         x, incX )                                       s, c

Level 2 BLAS

	        options                   dim   b-width scalar matrix  vector   scalar vector   prefixes

	_gemv ( order,        trans,      m, n,         alpha, a, lda, x, incX, beta,  y, incY ) S, C
	_gbmv ( order,        trans,      m, n, kL, kU, alpha, a, lda, x, incX, beta,  y, incY ) S, C
	_hemv ( order, uplo,                 n,         alpha, a, lda, x, incX, beta,  y, incY ) C
	_hbmv ( order, uplo,                 n, k,      alpha, a, lda, x, incX, beta,  y, incY ) C
	_hpmv ( order, uplo,                 n,         alpha, ap,     x, incX, beta,  y, incY ) C
	_symv ( order, uplo,                 n,         alpha, a, lda, x, incX, beta,  y, incY ) S
	_sbmv ( order, uplo,                 n, k,      alpha, a, lda, x, incX, beta,  y, incY ) S
	_spmv ( order, uplo,                 n,         alpha, ap,     x, incX, beta,  y, incY ) S
	_trmv ( order, uplo, trans, diag,    n,                a, lda, x, incX )                 S, C
	_tbmv ( order, uplo, trans, diag,    n, k,             a, lda, x, incX )                 S, C
	_tpmv ( order, uplo, trans, diag,    n,                ap,     x, incX )                 S, C
	_trsv ( order, uplo, trans, diag,    n,                a, lda, x, incX )                 S, C
	_tbsv ( order, uplo, trans, diag,    n, k,             a, lda, x, incX )                 S, C
	_tpsv ( order, uplo, trans, diag,    n,                ap,     x, incX )                 S, C

	        options                   dim   scalar vector   vector   matrix  prefixes

	_ger  ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) S
	_geru ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) C
	_gerc ( order,                    m, n, alpha, x, incX, y, incY, a, lda ) C
	_her  ( order, uplo,                 n, alpha, x, incX,          a, lda ) C
	_hpr  ( order, uplo,                 n, alpha, x, incX,          ap )     C
	_her2 ( order, uplo,                 n, alpha, x, incX, y, incY, a, lda ) C
	_hpr2 ( order, uplo,                 n, alpha, x, incX, y, incY, ap )     C
	_syr  ( order, uplo,                 n, alpha, x, incX,          a, lda ) S
	_spr  ( order, uplo,                 n, alpha, x, incX,          ap )     S
	_syr2 ( order, uplo,                 n, alpha, x, incX, y, incY, a, lda ) S
	_spr2 ( order, uplo,                 n, alpha, x, incX, y, incY, ap )     S

Level 3 BLAS

	        options                                 dim      scalar matrix  matrix  scalar matrix  prefixes

	_gemm ( order,             transA, transB,      m, n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) S, Z
	_symm ( order, side, uplo,                      m, n,    alpha, a, lda, b, ldb, beta,  c, ldc ) S, Z
	_hemm ( order, side, uplo,                      m, n,    alpha, a, lda, b, ldb, beta,  c, ldc ) C
	_syrk ( order,       uplo, trans,                  n, k, alpha, a, lda,         beta,  c, ldc ) S, Z
	_herk ( order,       uplo, trans,                  n, k, alpha, a, lda,         beta,  c, ldc ) C
	_syr2k( order,       uplo, trans,                  n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) S, Z
	_her2k( order,       uplo, trans,                  n, k, alpha, a, lda, b, ldb, beta,  c, ldc ) C
	_trmm ( order, side, uplo, transA,        diag, m, n,    alpha, a, lda, b, ldb )                S, Z
	_trsm ( order, side, uplo, transA,        diag, m, n,    alpha, a, lda, b, ldb )                S, Z

Meaning of prefixes

	S - float32	C - complex64

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

// Type SrotmParams contains Givens transformation parameters returned
// by the Float32 Srotm method.
type SrotmParams struct {
	Flag
	H [4]float32 // Column-major 2 by 2 matrix.
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

// Float32 implements the single precision real BLAS routines.
type Float32 interface {
	Float32Level1
	Float32Level2
	Float32Level3
}

// Float32Level1 implements the single precision real BLAS Level 1 routines.
type Float32Level1 interface {
	Sdsdot(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32
	Dsdot(n int, x []float32, incX int, y []float32, incY int) float64
	Sdot(n int, x []float32, incX int, y []float32, incY int) float32
	Snrm2(n int, x []float32, incX int) float32
	Sasum(n int, x []float32, incX int) float32
	Isamax(n int, x []float32, incX int) int
	Sswap(n int, x []float32, incX int, y []float32, incY int)
	Scopy(n int, x []float32, incX int, y []float32, incY int)
	Saxpy(n int, alpha float32, x []float32, incX int, y []float32, incY int)
	Srotg(a, b float32) (c, s, r, z float32)
	Srotmg(d1, d2, b1, b2 float32) (p SrotmParams, rd1, rd2, rb1 float32)
	Srot(n int, x []float32, incX int, y []float32, incY int, c, s float32)
	Srotm(n int, x []float32, incX int, y []float32, incY int, p SrotmParams)
	Sscal(n int, alpha float32, x []float32, incX int)
}

// Float32Level2 implements the single precision real BLAS Level 2 routines.
type Float32Level2 interface {
	Sgemv(o Order, tA Transpose, m, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int)
	Sgbmv(o Order, tA Transpose, m, n, kL, kU int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int)
	Strmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float32, lda int, x []float32, incX int)
	Stbmv(o Order, ul Uplo, tA Transpose, d Diag, n, k int, a []float32, lda int, x []float32, incX int)
	Stpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float32, x []float32, incX int)
	Strsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []float32, lda int, x []float32, incX int)
	Stbsv(o Order, ul Uplo, tA Transpose, d Diag, n, k int, a []float32, lda int, x []float32, incX int)
	Stpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []float32, x []float32, incX int)
	Ssymv(o Order, ul Uplo, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int)
	Ssbmv(o Order, ul Uplo, n, k int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int)
	Sspmv(o Order, ul Uplo, n int, alpha float32, ap []float32, x []float32, incX int, beta float32, y []float32, incY int)
	Sger(o Order, m, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int)
	Ssyr(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, a []float32, lda int)
	Sspr(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, ap []float32)
	Ssyr2(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int)
	Sspr2(o Order, ul Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32)
}

// Float32Level3 implements the single precision real BLAS Level 3 routines.
type Float32Level3 interface {
	Sgemm(o Order, tA, tB Transpose, m, n, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int)
	Ssymm(o Order, s Side, ul Uplo, m, n int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int)
	Ssyrk(o Order, ul Uplo, t Transpose, n, k int, alpha float32, a []float32, lda int, beta float32, c []float32, ldc int)
	Ssyr2k(o Order, ul Uplo, t Transpose, n, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int)
	Strmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha float32, a []float32, lda int, b []float32, ldb int)
	Strsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha float32, a []float32, lda int, b []float32, ldb int)
}

// Complex64 implements the single precision complex BLAS routines.
type Complex64 interface {
	Complex64Level1
	Complex64Level2
	Complex64Level3
}

// Complex64Level1 implements the single precision complex BLAS Level 1 routines.
type Complex64Level1 interface {
	Cdotu(n int, x []complex64, incX int, y []complex64, incY int) (dotu complex64)
	Cdotc(n int, x []complex64, incX int, y []complex64, incY int) (dotc complex64)
	Scnrm2(n int, x []complex64, incX int) float32
	Scasum(n int, x []complex64, incX int) float32
	Icamax(n int, x []complex64, incX int) int
	Cswap(n int, x []complex64, incX int, y []complex64, incY int)
	Ccopy(n int, x []complex64, incX int, y []complex64, incY int)
	Caxpy(n int, alpha complex64, x []complex64, incX int, y []complex64, incY int)
	Cscal(n int, alpha complex64, x []complex64, incX int)
	Csscal(n int, alpha float32, x []complex64, incX int)
}

// Complex64Level2 implements the single precision complex BLAS routines Level 2 routines.
type Complex64Level2 interface {
	Cgemv(o Order, tA Transpose, m, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Cgbmv(o Order, tA Transpose, m, n, kL, kU int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Ctrmv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex64, lda int, x []complex64, incX int)
	Ctbmv(o Order, ul Uplo, tA Transpose, d Diag, n, k int, a []complex64, lda int, x []complex64, incX int)
	Ctpmv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex64, x []complex64, incX int)
	Ctrsv(o Order, ul Uplo, tA Transpose, d Diag, n int, a []complex64, lda int, x []complex64, incX int)
	Ctbsv(o Order, ul Uplo, tA Transpose, d Diag, n, k int, a []complex64, lda int, x []complex64, incX int)
	Ctpsv(o Order, ul Uplo, tA Transpose, d Diag, n int, ap []complex64, x []complex64, incX int)
	Chemv(o Order, ul Uplo, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Chbmv(o Order, ul Uplo, n, k int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Chpmv(o Order, ul Uplo, n int, alpha complex64, ap []complex64, x []complex64, incX int, beta complex64, y []complex64, incY int)
	Cgeru(o Order, m, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int)
	Cgerc(o Order, m, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int)
	Cher(o Order, ul Uplo, n int, alpha float32, x []complex64, incX int, a []complex64, lda int)
	Chpr(o Order, ul Uplo, n int, alpha float32, x []complex64, incX int, a []complex64)
	Cher2(o Order, ul Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int)
	Chpr2(o Order, ul Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, ap []complex64)
}

// Complex64Level3 implements the single precision complex BLAS Level 3 routines.
type Complex64Level3 interface {
	Cgemm(o Order, tA, tB Transpose, m, n, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int)
	Csymm(o Order, s Side, ul Uplo, m, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int)
	Csyrk(o Order, ul Uplo, t Transpose, n, k int, alpha complex64, a []complex64, lda int, beta complex64, c []complex64, ldc int)
	Csyr2k(o Order, ul Uplo, t Transpose, n, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int)
	Ctrmm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int)
	Ctrsm(o Order, s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int)
	Chemm(o Order, s Side, ul Uplo, m, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int)
	Cherk(o Order, ul Uplo, t Transpose, n, k int, alpha float32, a []complex64, lda int, beta float32, c []complex64, ldc int)
	Cher2k(o Order, ul Uplo, t Transpose, n, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta float32, c []complex64, ldc int)
}
