# Copyright 2019 Red Hat Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

BINS = $(wildcard velero-*)

REPO ?= github.com/fusor/openshift-migration-plugin

BUILD_IMAGE ?= openshift/origin-release:golang-1.11

IMAGE ?= docker.io/fusor/openshift-migration-plugin

ARCH ?= amd64

BUILDTAGS ?= containers_image_ostree_stub exclude_graphdriver_devicemapper exclude_graphdriver_btrfs containers_image_openpgp exclude_graphdriver_overlay

SOURCES:=$(shell find . -name '*.go' -not -path "*/vendor/*" -not -path "*/.go*")

SOURCE_DIRS=velero-plugins

REMOVELINKS=remove-known-symlinks-after-dep-ensure.sh

all: $(addprefix build-, $(BINS))

build-%:
	$(MAKE) --no-print-directory BIN=$* build

build: _output/$(BIN)

_output/$(BIN): $(BIN)/*.go
	mkdir -p .go/src/$(REPO) .go/pkg .go/.cache .go/std/$(ARCH) _output
	cp -rp * .go/src/$(REPO)
	docker run \
				 --rm \
				 -u $$(id -u):$$(id -g) \
				 -v $$(pwd)/.go/pkg:/go/pkg:z \
				 -v $$(pwd)/.go/src:/go/src:z \
				 -v $$(pwd)/.go/std:/go/std:z \
				 -v $$(pwd)/.go/.cache:/go/.cache:z \
				 -v $$(pwd)/_output:/go/src/$(REPO)/_output:z \
				 -v $$(pwd)/.go/std/$(ARCH):/usr/local/go/pkg/linux_$(ARCH)_static:z \
				 -e CGO_ENABLED=0 \
				 -w /go/src/$(REPO) \
				 $(BUILD_IMAGE) \
				 go build -installsuffix "static" -tags "$(BUILDTAGS)" -i -v -o _output/$(BIN) ./$(BIN)

container:
	docker build -t $(IMAGE) -f Dockerfile .

test:
	go test -installsuffix "static" -tags "$(BUILDTAGS)" ./velero-plugins/...

ci: lint fmtcheck vet all test

clean:
	rm -rf .go _output

lint: ## Run golint
	@golint -set_exit_status $(addsuffix /... , $(SOURCE_DIRS))

fmt: ## Run go fmt
	@gofmt -d $(SOURCES)

fmtcheck: ## Check go formatting
	@gofmt -l $(SOURCES) | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

vet: ## Run go vet
	@go vet -installsuffix "static" -tags "$(BUILDTAGS)" $(addsuffix /..., $(addprefix ./, $(SOURCE_DIRS)))

help: ## Show this help screen
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''

rmlinks:
	-bash hack/$(REMOVELINKS)
