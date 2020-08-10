GOLANGCI_LINT_VERSION := "v1.27.0"
JSON_CLI_VERSION := "v1.6.7"

gen-2.0.0:
	@test -s $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) || (curl -sSfL https://github.com/swaggest/json-cli/releases/download/$(JSON_CLI_VERSION)/json-cli -o $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) && chmod +x $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION))
	cd resources/schema/ && $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) gen-go asyncapi-2.0.0.json --output ../../spec-2.0.0/entities.go --fluent-setters --package-name spec --root-name AsyncAPI
	gofmt -w ./spec-2.0.0/entities.go

gen-1.2.0:
	json-cli gen-go resources/schema/asyncapi.json --output ./spec/entities.go --fluent-setters --package-name spec --root-name AsyncAPI \
		--renames AsyncAPIAsyncapi100:Asyncapi100 AsyncAPIAsyncapi110:Asyncapi110 AsyncAPIAsyncapi120:Asyncapi120
	gofmt -w ./spec/entities.go

lint:
	@test -s $(GOPATH)/bin/golangci-lint-$(GOLANGCI_LINT_VERSION) || (curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /tmp $(GOLANGCI_LINT_VERSION) && mv /tmp/golangci-lint $(GOPATH)/bin/golangci-lint-$(GOLANGCI_LINT_VERSION))
	@$(GOPATH)/bin/golangci-lint-$(GOLANGCI_LINT_VERSION) run ./...
