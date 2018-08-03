all: carbonapi carbonzipper

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
        EXTRA_PKG_CONFIG_PATH=/opt/X11/lib/pkgconfig
endif

VERSION ?= $(shell git describe --abbrev=4 --dirty --always --tags)

GO ?= go

SOURCES=$(shell find . -name '*.go')

PKG_CARBONAPI=github.com/go-graphite/carbonapi/cmd/carbonapi
PKG_CARBONZIPPER=github.com/go-graphite/carbonapi/cmd/carbonzipper

carbonapi: $(SOURCES)
	PKG_CONFIG_PATH="$(EXTRA_PKG_CONFIG_PATH)" $(GO) build -tags cairo -ldflags '-X main.BuildVersion=$(VERSION)' $(PKG_CARBONAPI)

carbonzipper: $(SOURCES)
	$(GO) build --ldflags '-X main.BuildVersion=$(VERSION)' $(PKG_CARBONZIPPER)

debug: $(SOURCES)
	PKG_CONFIG_PATH="$(EXTRA_PKG_CONFIG_PATH)" $(GO) build -tags cairo -ldflags '-X main.BuildVersion=$(VERSION)' -gcflags=all='-l -N' $(PKG_CARBONAPI) $(PKG_CARBONZIPPER)

nocairo: $(SOURCES)
	$(GO) build -ldflags '-X main.BuildVersion=$(VERSION)'

test:
	PKG_CONFIG_PATH="$(EXTRA_PKG_CONFIG_PATH)" $(GO) test -tags cairo ./... -race

clean:
	rm -f carbonapi carbonzipper
