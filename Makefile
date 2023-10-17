################################################################################
# Global: Variables                                                            #
################################################################################

# Go build metadata variables
GOARCH            := $(shell go env GOARCH)
GOOS              := $(shell go env GOOS)
BUILDTYPE_DIR:=release

# Build output variables
CLI_BINARY := copa-grype
OUT_DIR    := ./dist
BINS_OUT_DIR      := $(OUT_DIR)/$(GOOS)_$(GOARCH)/$(BUILDTYPE_DIR)

################################################################################
# Target: build (default action)                                               #
################################################################################
.PHONY: build
build: 

	$(info $(INFOMARK) Building $(CLI_BINARY) ...)
	go build -o $(BINS_OUT_DIR)/$(CLI_BINARY);