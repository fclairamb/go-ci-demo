# It's just to make things simpler. This makefile is pretty much useless.

APP=./server
PORT=5000

ifeq "$(BUILD_VERSION)" ""
	BUILD_VERSION=0
endif

ifeq "$(GIT_COMMIT)" ""
	GIT_COMMIT="00000000-0000-0000-0000-000000000000"
endif

FILES=$(shell ls *.go) build.go

$(APP): $(FILES)
	go build -x -o $(APP)

build.go: Makefile
	printf 'package main;\nconst BUILD_VERSION = $(BUILD_VERSION)\nconst BUILD_GIT_COMMIT = "$(GIT_COMMIT)"' >build.go && go fmt build.go
	
start: $(APP)
	$(APP) -port $(PORT)

stop:
	fuser -k $(PORT)/tcp

clean:
	- rm $(APP) *~
