GOROOT := /usr/local/go
GOPATH := $(shell pwd)
GOBIN  := $(GOPATH)/bin
PATH   := $(GOROOT)/bin:$(PATH)
DEPS   := github.com/jmcvetta/napping github.com/spf13/cobra github.com/spf13/viper github.com/pr8kerl/f5er/f5 github.com/inconshreveable/mousetrap github.com/fatih/structs

LDFLAGS := -ldflags "-X main.commit=`git rev-parse HEAD`"

all: f5er

deps: $(DEPS)
	GOPATH=$(GOPATH) go get -u $^

f5er: main.go commands.go stack.go
    # always format code
		GOPATH=$(GOPATH) go fmt $^
    # binary
		GOPATH=$(GOPATH) go build $(LDFLAGS) -o $@ -v $^
		touch $@

linux64: main.go commands.go stack.go
    # always format code
		GOPATH=$(GOPATH) go fmt $^
		# vet it
		GOPATH=$(GOPATH) go tool vet $^
    # binary
		GOOS=linux GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o f5er-linux-amd64.bin -v $^
		touch f5er-linux-amd64.bin

win64: main.go commands.go stack.go
    # always format code
		GOPATH=$(GOPATH) go fmt $^
		# vet it
		GOPATH=$(GOPATH) go tool vet $^
    # binary
		GOOS=windows GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o f5er-win-amd64.exe -v $^
		touch f5er-win-amd64.exe

.PHONY: $(DEPS) clean

clean:
	rm -f f5er f5er-win-amd64.exe f5er-linux-amd64.bin
