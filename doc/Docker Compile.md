All else builds
```
docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang:1.7.1 go get
docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang:1.7.1 go build -v -o ./build/dockergetnetworkip_amd64
```
For Docker Alpine
```
docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang:1.7.1-alpine go get
docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang:1.7.1-alpine go build -v -o ./build/dockergetnetworkip_alpine
```
