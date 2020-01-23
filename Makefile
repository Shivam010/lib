GO = GO111MODULE=on go

fmt:
	${GO} fmt ./...

vet: fmt
	${GO} vet ./...

clean: vet
	${GO} mod tidy

build: clean
	${GO} build ./...

test: build
	${GO} test -v -cover ./...

.PHONY: test build clean vet fmt
