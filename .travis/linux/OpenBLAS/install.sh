set -ex

# fetch fortran to build OpenBLAS
sudo apt-get update -qq && sudo apt-get install -qq gfortran

# fetch OpenBLAS
pushd ~
sudo git clone --depth=1 git://github.com/xianyi/OpenBLAS

# make OpenBLAS
pushd OpenBLAS
echo OpenBLAS $(git rev-parse HEAD)
sudo make FC=gfortran COMMON_OPT=-O0 -s libs netlib shared &> /dev/null && sudo make PREFIX=/usr -s install
popd

# fetch cblas reference lib
curl http://www.netlib.org/blas/blast-forum/cblas.tgz | tar -zx

# make cblas and install
pushd CBLAS
sudo mv Makefile.LINUX Makefile.in
sudo make BLLIB=/usr/lib/libopenblas.a CFLAGS="-O0 -DADD_" FFLAGS=-O0 -s alllib
sudo mv lib/cblas_LINUX.a /usr/lib/libcblas.a
popd
popd

# install gonum/blas against OpenBLAS
export CGO_LDFLAGS="-L/usr/lib -lopenblas"
go get github.com/gonum/blas
pushd cgo
go install -v -x
popd
