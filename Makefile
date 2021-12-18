.DEFAULT_GOAL := images

MODULE_SLUGS=$(shell go run .)

BUILD_DIR=$(abspath _build)
IMAGE_BUILD_DIR=$(BUILD_DIR)/images
FRAME_BUILD_DIR=$(BUILD_DIR)/frames

IMAGES=$(patsubst %, $(IMAGE_BUILD_DIR)/%.svg, $(MODULE_SLUGS))

FRAMES=$(patsubst %, $(FRAME_BUILD_DIR)/%/port.svg, $(MODULE_SLUGS))

INSTALL_DIR=$(abspath _install)
IMAGE_INSTALL_DIR=$(INSTALL_DIR)/images
ASSET_INSTALL_DIR=$(INSTALL_DIR)/svg

INSTALLED_IMAGES=$(patsubst %, $(IMAGE_INSTALL_DIR)/%.svg, $(MODULE_SLUGS))
INSTALLED_FRAMES=$(patsubst $(FRAME_BUILD_DIR)/%, $(ASSET_INSTALL_DIR)/%, $(FRAMES))
INSTALLED_FACEPLATES=$(patsubst %, $(ASSET_INSTALL_DIR)/%.svg, $(MODULE_SLUGS))

PANEL_SOURCE_DIR=internal/panel

$(IMAGE_BUILD_DIR)/%.svg: $(PANEL_SOURCE_DIR)/%.go
	go run . $(patsubst $(PANEL_SOURCE_DIR)/%.go, %, $^)

images: $(IMAGES)

$(IMAGE_BUILD_DIR) $(ASSET_INSTALL_DIR):
	mkdir -p $(dir $@)

$(INSTALLED_IMAGES): $(IMAGE_BUILD_DIR)

$(IMAGE_INSTALL_DIR)/%: $(IMAGE_BUILD_DIR)/%
	mkdir -p $(dir $@)
	./scripts/install-svg.sh $(patsubst $(IMAGE_INSTALL_DIR)%, $(IMAGE_BUILD_DIR)/%, $@) $@

$(ASSET_INSTALL_DIR)/%.svg: $(IMAGE_BUILD_DIR)/%.svg
	mkdir -p $(dir $@)
	./scripts/install-faceplate.sh $(patsubst $(ASSET_INSTALL_DIR)/%.svg, $(IMAGE_BUILD_DIR)/%.svg, $@) $@

$(ASSET_INSTALL_DIR)/%: $(FRAME_BUILD_DIR)/%
	mkdir -p $(dir $@)
	./scripts/install-frames.sh $(dir $<) $(ASSET_INSTALL_DIR)

$(INSTALLED_FACEPLATES):

install: images $(INSTALLED_IMAGES) $(INSTALLED_FRAMES) $(INSTALLED_FACEPLATES)

clean:
	rm -rf $(BUILD_DIR)

clobber: clean
	rm -rf $(INSTALL_DIR)

.PHONY: clean clobber
