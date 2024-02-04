

# Variables
DIR_ARTIFACTS := $(shell pwd)/artifacts
DIR_DATA := $(shell pwd)/data
API_JSON_PATH := $(DIR_ARTIFACTS)/api.json
LIB_DIR := lib
DIR_PACKAGE := $(shell pwd)/$(LIB_DIR)
SPEC_PATH := https://infrahub-api-doc.nexgencloud.com/api.json
PROJECT := github.com\/nexgen\/hyperstack-sdk-go
TIME_PACKAGE := time
DEPS := github.com/deepmap/oapi-codegen/cmd/oapi-codegen

# Environment variables (set these in your environment or here directly)
export HYPERSTACK_STAGING := true
# export HYPERSTACK_API_KEY :=

# Targets
init: # Install all the necessary dependencies
	go get $(DEPS)
	go install $(DEPS)

pull-api: # Pulls the latest api.json from the server
	@echo "Pulling the latest swagger spec and generating Go SDK..."
	curl "$(SPEC_PATH)" | python3 -c "import sys, urllib.parse; print(urllib.parse.unquote(sys.stdin.read()))" > "$(API_JSON_PATH)"
	sed 's/Update Keypair name response/UpdatedKeypairNameResponseAPIObject/g' "$(API_JSON_PATH)"
	sed 's/Import Keypair Response/ImportedKeypairResponseAPIObject/g' "$(API_JSON_PATH)"

build: # Generates go sdk file from api.json
	@echo -n "Generating Go SDK... "
	export DIR_PACKAGE=$(shell pwd)/$(LIB_DIR)
	export DIR_ARTIFACTS=$(shell pwd)/artifacts
	go run sdk_generator.go
	mkdir -p "$(DIR_PACKAGE)/$(TIME_PACKAGE)"
	find "$(DIR_PACKAGE)" -name "*.go" -exec sed -i -e "s/*time.Time/*$(TIME_PACKAGE).CustomTime/g" {} \;
	find "$(DIR_PACKAGE)" -name "*.go" -exec sed -i -e 's/"time"/"$(PROJECT)\/$(LIB_DIR)\/$(TIME_PACKAGE)"/g' {} \;
	sed 's/PACKAGE_REPLACE/$(TIME_PACKAGE)/g' "$(DIR_DATA)/CustomTime.go" > "$(DIR_PACKAGE)/$(TIME_PACKAGE)/CustomTime.go"
	@echo "done"

test: # Runs tests
	go test -v sdk_generator_test.go

all: pull-api init build test

.PHONY: init pull-api build test all
