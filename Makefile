.PHONY: deps test clean

GOPATH ?= /go
GOBIN  := $(GOPATH)/bin
PATH   := $(GOPATH)/bin:$(PATH)
PROJ   := f5er
DOCKER_USERNAME ?= Monkey
DOCKER_PASSWORD ?= Magic

LDFLAGS := -ldflags "-X main.commit=`git rev-parse HEAD`"

all: fmt test $(PROJ)

fmt:
	GOPATH=$(GOPATH) go fmt *.go
	GOPATH=$(GOPATH) go fmt f5/*.go
	GOPATH=$(GOPATH) go vet 
	GOPATH=$(GOPATH) go vet ./f5

test: fmt
	@echo "+++ Is this thing working? :hammer_and_wrench:"
	GOPATH=$(GOPATH) go test -cover -v 

$(PROJ):
	CGO_ENABLED=0 GOPATH=$(GOPATH) go build $(LDFLAGS) -o $@ -v
	touch $@ && chmod 755 $@

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o $(PROJ)-linux-amd64 -v
	touch $(PROJ)-linux-amd64 && chmod 755 $(PROJ)-linux-amd64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o $(PROJ)-windows-amd64.exe -v
	touch $(PROJ)-windows-amd64.exe

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 GOPATH=$(GOPATH) go build -o $(PROJ)-darwin-amd64 -v
	touch $(PROJ)-darwin-amd64 && chmod 755 $(PROJ)-darwin-amd64

clean:
	rm -rf $(PROJ) $(PROJ)-windows-amd64.exe $(PROJ)-linux-amd64 $(PROJ)-darwin-amd64 dist

