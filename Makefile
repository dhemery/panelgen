BUILD_DIR=_build
FACEPLATE_BUILD_DIR=$(BUILD_DIR)/faceplates
CONTROL_BUILD_DIR=$(BUILD_DIR)/frames

MODULE_SLUGS=$(shell go run .)

INSTALL_DIR=_install
IMAGE_INSTALL_DIR=$(INSTALL_DIR)/images
ASSET_INSTALL_DIR=$(INSTALL_DIR)/svg

PANEL_SOURCE_DIR=internal/panel

FACEPLATES=$(patsubst %, $(FACEPLATE_BUILD_DIR)/%.svg, $(MODULE_SLUGS))

$(FACEPLATE_BUILD_DIR)/%.svg: $(PANEL_SOURCE_DIR)/%.go
	go run . $(patsubst $(PANEL_SOURCE_DIR)/%.go, %, $^)

panels: $(FACEPLATES)

clean:
	rm -rf $(BUILD_DIR)

clobber: clean
	rm -rf $(INSTALL_DIR)

.PHONY: clean clobber

.DEFAULT_GOAL := panels
