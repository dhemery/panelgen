BUILD_DIR=_build
FACEPLATE_BUILD_DIR=$(BUILD_DIR)/faceplates
CONTROL_BUILD_DIR=$(BUILD_DIR)/frames
INSTALL_DIR=_install
IMAGE_INSTALL_DIR=$(INSTALL_DIR)/images
ASSET_INSTALL_DIR=$(INSTALL_DIR)/svg

svg:
	go run .

clean:
	rm -rf $(BUILD_DIR)

clobber: clean
	rm -rf $(INSTALL_DIR)

.PHONY: clean clobber svg
