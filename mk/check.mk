CLANG_FORMAT_PATH ?= clang-format

.PHONY: fmt
fmt: fmt/go fmt/proto ## Dev: Run various format tools

.PHONY: fmt/go
fmt/go: ## Dev: Run go fmt
	go fmt $(GOFLAGS) ./...

.PHONY: fmt/proto
fmt/proto: ## Dev: Run clang-format on .proto files
	which $(CLANG_FORMAT_PATH) && find . -name '*.proto' | xargs -L 1 $(CLANG_FORMAT_PATH) -i || true

.PHONY: tidy
tidy:
	@TOP=$(shell pwd) && \
	for m in $$(find . -name go.mod) ; do \
		( cd $$(dirname $$m) && go mod tidy ) ; \
	done

.PHONY: shellcheck
shellcheck:
	find . -name "*.sh" -not -path "./.git/*" -exec shellcheck -P SCRIPTDIR -x {} +

.PHONY: golangci-lint
golangci-lint: ## Dev: Runs golangci-lint linter
	# Starting with golangci-lint v1.47.1, the CI job runs OOM if all of these
	# linters are used together. The first set is the largest that ran without
	# OOM.
	$(GOLANGCI_LINT_DIR)/golangci-lint run \
		--disable-all \
		--enable bodyclose,contextcheck,errcheck,gci,gocritic,gofmt,gomodguard,govet,importas,ineffassign,misspell,typecheck,unconvert,unparam,whitespace \
		--timeout=10m -v
	$(GOLANGCI_LINT_DIR)/golangci-lint run \
		--disable-all \
		--enable gosimple,staticcheck,unused \
		--timeout=10m -v

.PHONY: helm-lint
helm-lint:
	for c in ./deployments/charts/*; do \
  		if [ -d $$c ]; then \
			helm lint --strict $$c; \
		fi \
	done

.PHONY: ginkgo/unfocus
ginkgo/unfocus:
	$(GOPATH_BIN_DIR)/ginkgo unfocus

.PHONY: format
format: fmt generate docs tidy ginkgo/unfocus

.PHONY: kube-lint
kube-lint:
	$(GOPATH_BIN_DIR)/kube-linter lint .

.PHONY: check
check: format helm-lint golangci-lint shellcheck kube-lint ## Dev: Run code checks (go fmt, go vet, ...)
	git diff --quiet || test $$(git diff --name-only | grep -v -e 'go.mod$$' -e 'go.sum$$' | wc -l) -eq 0 || ( echo "The following changes (result of code generators and code checks) have been detected:" && git --no-pager diff && false ) # fail if Git working tree is dirty
