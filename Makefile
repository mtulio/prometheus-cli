# Defaut APP name to build
APP_NAME ?= prometheus-cli
BIN_PATH := $(PWD)/bin/

include Makefile-glob.mk

# Builder
.PHONY: build run clean
build: clean
	@test -d $(PWD)/bin || mkdir $(PWD)/bin
	@$(foreach dircmd,$(shell ls cmd/), \
		cd cmd/$(dircmd); \
		GOOS=$(GOOS) GOARCH=$(GOARCH) \
			go build \
			-ldflags "$(LDFLAGS)" \
			$(BUILD_TAGS) \
			-o $(BIN_PATH)/$(APP_NAME) && strip $(BIN_PATH)/$(APP_NAME) \
	; cd -)

run:
	$(BIN_PATH)/$(APP_NAME)

clean:
	@rm -rf bin/* dist/* |true


# ##################
# Release Management
#

# Release
tag:
	$(call deps_tag,$@)
	git tag -a $(shell cat VERSION) -m "$(message)"
	git push origin $(shell cat VERSION)


# Goreleaser
# https://goreleaser.com/introduction/
GORELEASE_ROOT ?= ../../
gorelease-init:
	goreleaser init

release: dev-release-latest
	@$(foreach dircmd,$(shell ls cmd/), \
	cd cmd/$(dircmd); \
	. $(GORELEASE_ROOT)/hack/env-build.sh && \
		goreleaser --rm-dist -f $(GORELEASE_ROOT)/.goreleaser.yml \
	; cd -)

release-snap:
	@$(foreach dircmd,$(shell ls cmd/), \
	cd cmd/$(dircmd); \
	. $(GORELEASE_ROOT)/hack/env-build.sh && \
		goreleaser --rm-dist --snapshot -f $(GORELEASE_ROOT)/.goreleaser.yml \
	; cd -)

# DEV Release builder
# Using ghr to avoid dependencies in goreleaser
dev-release: build
	ghr $(RELEASE_VERSION) bin/

dev-release-latest: build
	ghr --recreate latest bin/


# ##################
# Makefile functions

define deps_tag
	@if [[ "$(message)"x == "x" ]]; then \
		echo -e "\n Error: the commit message was not provided."; \
		$(call show_usage) \
		exit 1; \
	fi
endef