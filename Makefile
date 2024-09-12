#GOLANGCI_LINT_VERSION := "v1.61.0" # Optional configuration to pinpoint golangci-lint version.

# The head of Makefile determines location of dev-go to include standard targets.
GO ?= go
export GO111MODULE = on

ifneq "$(GOFLAGS)" ""
  $(info GOFLAGS: ${GOFLAGS})
endif

ifneq "$(wildcard ./vendor )" ""
  $(info Using vendor)
  modVendor =  -mod=vendor
  ifeq (,$(findstring -mod,$(GOFLAGS)))
      export GOFLAGS := ${GOFLAGS} ${modVendor}
  endif
  ifneq "$(wildcard ./vendor/github.com/bool64/dev)" ""
  	DEVGO_PATH := ./vendor/github.com/bool64/dev
  endif
endif

ifeq ($(DEVGO_PATH),)
	DEVGO_PATH := $(shell GO111MODULE=on $(GO) list ${modVendor} -f '{{.Dir}}' -m github.com/bool64/dev)
	ifeq ($(DEVGO_PATH),)
    	$(info Module github.com/bool64/dev not found, downloading.)
    	DEVGO_PATH := $(shell export GO111MODULE=on && $(GO) get github.com/bool64/dev && $(GO) list -f '{{.Dir}}' -m github.com/bool64/dev)
	endif
endif

-include $(DEVGO_PATH)/makefiles/main.mk
-include $(DEVGO_PATH)/makefiles/lint.mk
-include $(DEVGO_PATH)/makefiles/test-unit.mk
-include $(DEVGO_PATH)/makefiles/bench.mk
-include $(DEVGO_PATH)/makefiles/reset-ci.mk

# Add your custom targets here.

## Run tests
test: test-unit

JSON_CLI_VERSION := v1.11.0

## Generate bindings for v2.4.0 spec.
gen-2.4.0:
	@test -s $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) || (curl -sSfL https://github.com/swaggest/json-cli/releases/download/$(JSON_CLI_VERSION)/json-cli -o $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) && chmod +x $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION))
	cd resources/schema/ && ./prepare_bindings.sh && $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) gen-go asyncapi-2.4.0-fixed.json --output ../../spec-2.4.0/entities.go  --fluent-setters --package-name spec --root-name AsyncAPI --config ./asyncapi-2.4.0-gen-cfg.json --schema-resolver ./bindings-resolver.json
	make fix-lint


## Generate bindings for v2.1.0 spec.
gen-2.1.0:
	@test -s $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) || (curl -sSfL https://github.com/swaggest/json-cli/releases/download/$(JSON_CLI_VERSION)/json-cli -o $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) && chmod +x $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION))
	cd resources/schema/ && $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) gen-go asyncapi-2.1.0.json --patches asyncapi-2.1.0-patch.json --output ../../spec-2.1.0/entities.go --validate-required --fluent-setters --package-name spec --root-name AsyncAPI
	make fix-lint

## Generate bindings for v2.0.0 spec.
gen-2.0.0:
	@test -s $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) || (curl -sSfL https://github.com/swaggest/json-cli/releases/download/$(JSON_CLI_VERSION)/json-cli -o $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) && chmod +x $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION))
	cd resources/schema/ && $(GOPATH)/bin/json-cli-$(JSON_CLI_VERSION) gen-go asyncapi-2.0.0.json --output ../../spec-2.0.0/entities.go --validate-required --fluent-setters --package-name spec --root-name AsyncAPI
	make fix-lint

## Generate bindings for v1.2.0 spec.
gen-1.2.0:
	json-cli gen-go resources/schema/asyncapi.json --output ./spec/entities.go --fluent-setters --package-name spec --root-name AsyncAPI \
		--renames AsyncAPIAsyncapi100:Asyncapi100 AsyncAPIAsyncapi110:Asyncapi110 AsyncAPIAsyncapi120:Asyncapi120
	make fix-lint
