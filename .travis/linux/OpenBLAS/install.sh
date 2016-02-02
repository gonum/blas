sudo apt-get update -qq
sudo apt-get install -qq gfortran
pushd ~
sudo git clone --depth=1 git://github.com/xianyi/OpenBLAS
pushd OpenBLAS
echo OpenBLAS $(git rev-parse HEAD)
sudo make FC=gfortran &> /dev/null
sudo make PREFIX=/usr install
popd
curl http://www.netlib.org/blas/blast-forum/cblas.tgz | tar -zx
pushd CBLAS
sudo mv Makefile.LINUX Makefile.in
sudo BLLIB=/usr/lib/libopenblas.a make alllib
sudo mv lib/cblas_LINUX.a /usr/lib/libcblas.a
popd
popd
export CGO_LDFLAGS="-L/usr/lib -lopenblas -lm"
go get github.com/gonum/blas
pushd cgo
go install -v -x
popd
