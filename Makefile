export GITCOMMIT=$(shell git log -1 --pretty=format:"%H")
export GOLDFLAGS=-s -w -extldflags '-zrelro -znow' -X go.eqrx.net/mauzr.version=$(GITTAG) -X go.eqrx.net/mauzr.commit=$(GITCOMMIT)
export GOFLAGS=-trimpath
export GITTAG=$(shell git describe --tags --always)
export IMAGE=docker.pkg.github.com/eqrx/mauzr/mauzr

.PHONY: all
all: build

.PHONY: build
build:
	go build -ldflags "$(GOLDFLAGS)" ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: benchmark
benchmark:
	go test -bench=. ./...

.PHONY: test
test:
	go test -race ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: download
download:
	go mod download

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: build-image
build-image: build
	./build/buildah.sh

.PHONY: push-image
push-image:
	buildah manifest push --all $(IMAGE):$(GITTAG) docker://$(IMAGE):$(GITTAG)

.PHONY: update
update:
	go get -t -u=patch ./...
