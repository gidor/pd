BIN="./bin"
SRC=$(shell find . -name "*.go")

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

ifeq (, $(shell which richgo))
$(warning "could not find richgo in $(PATH), run: go get github.com/kyoh86/richgo")
endif

# .PHONY: fmt lint test tr_cli install_deps clean
.PHONY: fmt lint cli install_deps clean

default: cli

all: fmt test cli
# all: fmt test tr_cli

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

test: install_deps lint
	$(info ******************** running tests ********************)
	richgo test -v ./...

cli: install_deps
	$(info ******************** building CLI ********************)
	mkdir -p $(BIN)
	go build -o $(BIN)/cnv

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

clean:
	rm -rf $(BIN)