GOPATH := /go
GOBIN  := $(GOPATH)/bin
PATH   := $(GOPATH)/bin:$(PATH)
PROJ   := f5er

LDFLAGS := -ldflags "-X main.commit=`git rev-parse HEAD`"

all: deps fmt test $(PROJ)

deps: $(DEPS)
	GOPATH=$(GOPATH) glide install

test: deps
		GOPATH=$(GOPATH) go test -cover -v $(shell glide novendor)

fmt:
		GOPATH=$(GOPATH) go fmt $(glide novendor)
		GOPATH=$(GOPATH) go tool vet *.go f5/*.go

$(PROJ): deps 
		GOPATH=$(GOPATH) go build $(LDFLAGS) -o $@ -v $(glide novendor)
		touch $@ && chmod 755 $@

build:
		GOOS=linux GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o $(PROJ) -v $(glide novendor)
		touch $(PROJ) && chmod 755 $(PROJ)

linux: deps
		GOOS=linux GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o $(PROJ)-linux-amd64 -v $(glide novendor)
		touch $(PROJ)-linux-amd64 && chmod 755 $(PROJ)-linux-amd64

windows: deps
		GOOS=windows GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o $(PROJ)-windows-amd64.exe -v $(glide novendor)
		touch $(PROJ)-windows-amd64.exe

darwin: deps
		GOOS=darwin GOARCH=amd64 GOPATH=$(GOPATH) go build -o $(PROJ)-darwin-amd64 -v $(glide novendor)
		touch $(PROJ)-darwin-amd64 && chmod 755 $(PROJ)-darwin-amd64

.PHONY: $(DEPS) clean

clean:
		rm -rf $(PROJ) $(PROJ)-win-amd64.exe $(PROJ)-linux-amd64 $(PROJ)-darwin-amd64 .glide vendor

