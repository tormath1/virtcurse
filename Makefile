VERSION= $(shell git describe --tags --always)
LDFLAGS=-ldflags "-X github.com/tormath1/virtcurse/cmd.Version=${VERSION}"

build:
	go build -o virtcurse $(LDFLAGS) main.go

lint:
	goimports -w ./
