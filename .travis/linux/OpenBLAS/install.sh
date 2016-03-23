set -ex

# remove exiting blas packages to prevent conflicts
sudo apt-get remove libblas3gf libatlas3gf-base

# fetch fortran to build OpenBLAS
sudo apt-get update -qq
sudo apt-get install -qq gfortran

# source for blas and lapack packages
LAPACK_ROOT="http://mirrors.xmission.com/ubuntu/pool/main/l/lapack"
OPENBLAS_ROOT="http://mirrors.xmission.com/ubuntu/pool/universe/o/openblas"

packagelist=(
  "$LAPACK_ROOT/libblas-common_3.6.0-2ubuntu2_amd64.deb"
  "$LAPACK_ROOT/libblas3_3.6.0-2ubuntu2_amd64.deb"
  "$LAPACK_ROOT/libblas-dev_3.6.0-2ubuntu2_amd64.deb"
  "$OPENBLAS_ROOT/libopenblas-base_0.2.15-1build1_amd64.deb"
  "$OPENBLAS_ROOT/libopenblas-dev_0.2.15-1build1_amd64.deb"
  )

for i in ${packagelist[*]};
do
  fn=${i##*/}
  echo "fetching $fn"
  curl -O $i
  sudo dpkg -i $fn
done

# install gonum/blas against OpenBLAS
export CGO_LDFLAGS="-L/usr/lib -lopenblas"
go get github.com/gonum/blas
pushd cgo
go install -v -x
popd
