APPNAME := stackguru-go

DESTDIR ?= /usr/local

SOURCEDIRS := $(glide novendor)
SOURCES := $(shell find $(SOURCEDIRS) -name '*.go')

VERSION := $(shell git describe)

GO_LDFLAGS = -X main.Version=$(VERSION)

ifdef GO_LDFLAGS
	GO_FLAGS += -ldflags '$(GO_LDFLAGS)'
endif

.DEFAULT_GOAL: all

all: build

build: $(APPNAME)

$(APPNAME): $(SOURCES)
	go build $(GO_FLAGS) -o bin/$(APPNAME)

.PHONY: install
install: all
	install bin/$(APPNAME) $(DESTDIR)/bin

.PHONY: test
test:
	go test $(shell glide novendor)
