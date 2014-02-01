go-ci-demo
==========

Continuous Integration around a go project.

This is how it's handled on the jenkins script:

    make clean
    export GOROOT=/usr/local/go
    export PATH=$PATH:$GOROOT/bin
    export GOPATH=`pwd`/go
    go get -d
    make
    make stop PORT=8601 || printf ""

Demonstration page: http://go-ci-demo.webingenia.com/
