GOROOT := /usr/local/go
GOPATH := $(shell pwd)
GOBIN  := $(GOPATH)/bin
PATH   := $(GOROOT)/bin:$(PATH)
DEPS   = code.google.com/p/gopass github.com/jmcvetta/napping github.com/spf13/cobra github.com/spf13/viper

all: f5er

update: $(DEPS)
	GOPATH=$(GOPATH) go get -u $^

f5er: src/f5er.go
    # always format code
		GOPATH=$(GOPATH) go fmt $^
    # binary
		GOPATH=$(GOPATH) go build -o $@ -v $^
		touch $@

.PHONY: $(DEPS)

clean:
	rm -f f5er
