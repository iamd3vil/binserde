PHONY : build run fresh test clean pack-releases install-deps

BIN := binserde.bin

LAST_COMMIT := $(shell git rev-parse --short HEAD)
LAST_COMMIT_DATE := $(shell git show -s --format=%ci ${LAST_COMMIT})
VERSION := $(shell git describe --tags)
BUILDSTR := ${VERSION} (Commit: ${LAST_COMMIT_DATE} (${LAST_COMMIT}), Build: $(shell date +"%Y-%m-%d% %H:%M:%S %z"))

STATIC := ./templates:/templates

build:
	go build -o ${BIN} -ldflags="-X 'main.buildString=${BUILDSTR}'" ./cmd/generator/
	stuffbin -a stuff -in ${BIN} -out ${BIN} ${STATIC}

run:
	./${BIN}

fresh: clean build

test: build
	rm -rf ./test/binserde_gen.go
	./binserde.bin --dir test --file test/binserde_gen.go
	go test -v ./test

bench: build
	rm -rf ./test/binserde_gen.go
	./binserde.bin --dir test --file test/binserde_gen.go
	go test -v -bench=BenchmarkMarshalUnmarshal ./test

clean:
	go clean
	rm -f ${BIN}

install-deps:
	go install github.com/knadh/stuffbin/...


# pack-releases runs stuffbin packing on a given list of
# binaries. This is used with goreleaser for packing
# release builds for cross-build targets.
pack-releases:
	$(foreach var,$(RELEASE_BUILDS),stuffbin -a stuff -in ${var} -out ${var} ${STATIC};)