BUILD_DIR=_build
FACEPLATE_BUILD_DIR=$(BUILD_DIR)/faceplates
CONTROL_BUILD_DIR=$(BUILD_DIR)/controls
INSTALL_DIR=_install
IMAGE_INSTALL_DIR=$(INSTALL_DIR)/images
ASSET_INSTALL_DIR=$(INSTALL_DIR)/svg

svg:
	go run .

clean:
	rm -rf $(BUILD_DIR) out

clobber: clean
	rm -rf $(INSTALL_DIR)

$(IMAGE_INSTALL_DIR) $(ASSET_INSTALL_DIR) $(FACEPLATE_BUILD_DIR): $(ASSET_INSTALL_DIR)
	mkdir -p $@

.PHONY: clean clobber svg
