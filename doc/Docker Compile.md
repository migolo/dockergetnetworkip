All else builds
docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang go get
docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang go build -v -o dockergetnetworkip

For Docker Alpine
docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" blang/golang-alpine go get
docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" blang/golang-alpine go build -v -o dockergetnetworkip