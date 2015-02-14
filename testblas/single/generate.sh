#!/usr/bin/env bash
#
# Copyright ©2015 The gonum Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

echo Generating level1.go
cat ../level1double.go |
	gofmt -r 'float64 -> float32' |
	gofmt -r 'DrotmParams -> SrotmParams' |
	gofmt -r 'Dasum -> Sasum' |
	gofmt -r 'Daxpy -> Saxpy' |
	gofmt -r 'Ddot -> Sdot' |
	gofmt -r 'Dnrm2 -> Snrm2' |
	gofmt -r 'Idamax -> Isamax' |
	gofmt -r 'Dswap -> Sswap' |
	gofmt -r 'Dcopy -> Scopy' |
	gofmt -r 'Drotg -> Srotg' |
	gofmt -r 'Drotmg -> Srotmg' |
	gofmt -r 'Drot -> Srot' |
	gofmt -r 'Drotm -> Srotm' |
	gofmt -r 'Dscal -> Sscal' |
	sed -e 's_"math"_math "github.com/gonum/blas/native/internal/math32"_' |
	gofmt -r 'testblas -> single' |
	goimports > level1.go

echo Generating common.go
cat ../common.go |
	gofmt -r 'float64 -> float32' |
	gofmt -r '1e-14 -> 1e-6' |
	sed -e 's_"math"_math "github.com/gonum/blas/native/internal/math32"_' |
	gofmt -r 'testblas -> single' |
	goimports > common.go
