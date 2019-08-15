
PROJECTNAME := $(shell basename "$(PWD)")
BASE := $(shell pwd)
BIN := $(BASE)/bin
FILES := *.go

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

run:
	@go run *.go

build:
	@echo "Building binary"
	@go build $(LDFLAGS) -o $(BIN)/$(PROJECTNAME) $(FILES)

clean:
	@echo "Cleaning"
	@rm -fr $(BIN) 2> /dev/null
