# It's just to make things simpler. This makefile is pretty much useless.

APP=./go-ci-demo
PORT=5000

ifeq "$(BUILD_NUMBER)" ""
	BUILD_NUMBER=0
endif

ifeq "$(GIT_COMMIT)" ""
	GIT_COMMIT=00000000-0000-0000-0000-000000000000
endif

FILES=$(shell ls *.go)

$(APP): $(FILES) Makefile build.go
	go build -x

build.go:
	printf 'package main;\nconst BUILD_NUMBER = $(BUILD_NUMBER)\nconst BUILD_GIT_COMMIT = "$(GIT_COMMIT)"' >build.go && go fmt build.go
	
start: $(APP)
	$(APP) -port $(PORT)

stop:
	fuser -k $(PORT)/tcp

clean:
	- rm $(APP) *~
