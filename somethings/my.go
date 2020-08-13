set GO111MODULE=on
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\wicky\AppData\Local\go-build
set GOENV=C:\Users\wicky\AppData\Roaming\go\env
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=E:\Go-workspace
set GOPRIVATE=
set GOPROXY=https://goproxy.cn
set GOROOT=c:\go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=c:\go\pkg\tool\windows_amd64
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=NUL
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fmessage-length=0 -fdebug-prefix-map=C:\Users\wicky\AppData\Local\Temp\go-build716206338=/tmp/go-build -gno-record-gcc-switches

E:\Go-workspace\bin>dir /b /l
dlv
go-outline
gocode-gomod
gocode
godef
goimports
golint
gopkgs
gopls
goreturns

The code in the workspace failed to compile (see the error message below).
If you believe this is a mistake, please file an issue: https://github.com/golang/go/issues/new.
go [-e -json -compiled=true -test=true -export=false -deps=true -find=false -- ./]: exit status 1: go: cannot find main module; see 'go help modules'
: packages.Load error

//go tools安装
mkdir -p $env:GOPATH/src/golang.org/x
cd $env:GOPATH/src/golang.org/x
git clone https://github.com/golang/tools.git
cd tools/cmd/
go install ...或go get ...

godoc -http=:6060
go doc url
