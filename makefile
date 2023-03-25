TARGET = pgii
BUILD_PATH ?= build

VERSION := v1.0
PACKAGE_NAME := pgii-$(VERSION)

all: $(TARGET)

.PHONY: build_prepare
build_prepare:
	@mkdir -p $(BUILD_PATH)/$(PACKAGE_NAME)

.PHONY: $(TARGET)
$(TARGET): build_prepare
	@go build -gcflags="all=-N -l" -o ${BUILD_PATH}/${PACKAGE_NAME} src/main.go
	@echo "Build successfully"

.PHONY: pack
pack:
	@if [ -e out ] ; then rm -rf out; fi
	@mkdir out
	@cp $(TARGET) ./out/$(TARGET)
	tar -C out -zcf $(TARGET)-v$(VERSION).tar.gz .

.PHONY: clean
clean:
	@rm -rf ./build/*
	@rm -rf ./$(TARGET)
	@go clean -i .
	@echo "Clean successfully"