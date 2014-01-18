go-ci-demo
==========

Continuous Integration around a go project

    make clean
    export GOROOT=/usr/local/go
    export PATH=$PATH:$GOROOT/bin
    export GOPATH=`pwd`/go
    go get -d
    make
    make stop PORT=8601 || printf ""
