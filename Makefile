# It's just to make things simpler. This makefile is pretty much useless.

APP=./server
PORT=5000

$(APP): server.go
	go build -o $(APP) server.go
	
start: $(APP)
	$(APP) -port $(PORT)

stop:
	fuser -k $(PORT)/tcp
