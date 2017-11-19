.PHONY: deps test publish clean

GOPATH ?= /go
GOBIN  := $(GOPATH)/bin
PATH   := $(GOPATH)/bin:$(PATH)
PROJ   := f5er

LDFLAGS := -ldflags "-X main.commit=`git rev-parse HEAD`"

all: deps fmt test $(PROJ)

deps:
	@echo "--- collecting ingredients :bento:"
	GOPATH=$(GOPATH) glide install

fmt:
	GOPATH=$(GOPATH) go fmt $(glide novendor)
	GOPATH=$(GOPATH) go tool vet *.go f5/*.go

test: fmt deps
	@echo "+++ Is this thing working? :hammer_and_wrench:"
	GOPATH=$(GOPATH) go test -cover -v $(glide novendor)

$(PROJ): deps 
	GOPATH=$(GOPATH) go build $(LDFLAGS) -o $@ -v $(glide novendor)
	touch $@ && chmod 755 $@

linux: deps
	GOOS=linux GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o $(PROJ)-linux-amd64 -v $(glide novendor)
	touch $(PROJ)-linux-amd64 && chmod 755 $(PROJ)-linux-amd64

windows: deps
	GOOS=windows GOARCH=amd64 GOPATH=$(GOPATH) go build $(LDFLAGS) -o $(PROJ)-windows-amd64.exe -v $(glide novendor)
	touch $(PROJ)-windows-amd64.exe

darwin: deps
	GOOS=darwin GOARCH=amd64 GOPATH=$(GOPATH) go build -o $(PROJ)-darwin-amd64 -v $(glide novendor)
	touch $(PROJ)-darwin-amd64 && chmod 755 $(PROJ)-darwin-amd64

ifdef TRAVIS_TAG
publish: deps
	@echo "+++ release :octocat:"
	docker-compose run releaser --skip-validate --rm-dist
endif

clean:
		rm -rf $(PROJ) $(PROJ)-win-amd64.exe $(PROJ)-linux-amd64 $(PROJ)-darwin-amd64 .glide vendor dist
