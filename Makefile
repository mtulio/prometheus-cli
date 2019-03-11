# Defaut APP name to build
APP_NAME ?= prometheus-cli
BIN_PATH := $(PWD)/bin/

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