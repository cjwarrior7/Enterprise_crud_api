BINARY=accountingService

.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u github.com/alecthomas/gometalinter
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/dep/cmd/dep
	gometalinter --install --update
	@$(MAKE) dep

.PHONY: dep
dep: ## Run dep ensure and prune
	dep ensure
	dep prune

.PHONY: fmt
fmt: ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file";gofmt -s -w "$$file"; done

.PHONY: test
test: ## Run all the tests

.PHONY: lint
lint: ## Run all the linters
	gometalinter.v2 --vendor --disable-all \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=staticcheck \
		--enable=gofmt \
		--enable=goimports \
		--enable=misspell \
		--enable=errcheck \
		--enable=golint \
		--enable=vet \
		--enable=vetshadow \
		--exclude=gocyclo  \
		--deadline=10m \
		./...

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: build
build: clean test ## Build binary
	go build -v -o ${BINARY}

.PHONY: clean
clean: ## Remove temporary files
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
