set -ex
uname -a
sudo apt-get remove libblas3gf libatlas3gf-base

curl -O http://mirrors.xmission.com/ubuntu/pool/main/b/blas/libblas-common_1.2.20110419-10_amd64.deb
sudo dpkg -i libblas-common_1.2.20110419-10_amd64.deb
sudo apt-get install -f

curl -O http://mirrors.xmission.com/ubuntu/pool/main/b/blas/libblas3_1.2.20110419-10_amd64.deb
sudo dpkg -i libblas3_1.2.20110419-10_amd64.deb
sudo apt-get install -f

curl -O http://mirrors.xmission.com/ubuntu/pool/universe/b/blas/libblas3gf_1.2.20110419-10_all.deb
sudo dpkg -i libblas3gf_1.2.20110419-10_all.deb
sudo apt-get install -f

curl -O http://mirrors.xmission.com/ubuntu/pool/universe/o/openblas/libopenblas-base_0.2.14-1ubuntu1_amd64.deb
sudo dpkg -i libopenblas-base_0.2.14-1ubuntu1_amd64.deb
sudo apt-get install -f
sudo ln -s /usr/lib/libopenblas.so.0 /usr/lib/libopenblas.so

# install gonum/blas against OpenBLAS
export CGO_LDFLAGS="-L/usr/lib -lopenblas"
go get github.com/gonum/blas
pushd cgo
go install -v -x
popd
