BUILD_DIR=_build
IMAGE_BUILD_DIR=$(BUILD_DIR)/images
CONTROL_BUILD_DIR=$(BUILD_DIR)/frames

MODULE_SLUGS=$(shell go run .)

INSTALL_DIR=_install
IMAGE_INSTALL_DIR=$(INSTALL_DIR)/images
ASSET_INSTALL_DIR=$(INSTALL_DIR)/svg

PANEL_SOURCE_DIR=internal/panel

IMAGES=$(patsubst %, $(IMAGE_BUILD_DIR)/%.svg, $(MODULE_SLUGS))

$(IMAGE_BUILD_DIR)/%.svg: $(PANEL_SOURCE_DIR)/%.go
	go run . $(patsubst $(PANEL_SOURCE_DIR)/%.go, %, $^)

images: $(IMAGES)

clean:
	rm -rf $(BUILD_DIR)

clobber: clean
	rm -rf $(INSTALL_DIR)

.PHONY: clean clobber

 # def install_svg(from, to, *options)                                                                    
 #   from = from.expand_path.to_s                                                                         
 #   to = to.expand_path.to_s                                                                             
 #   sh './scripts/install-svg.sh', from, to, *options                                                    
 # end                                                                                                    
 #                                                                                                        
 # def install_faceplate(from, to)                                                                        
 #   install_svg from, to, '--export-id=faceplate', '--export-id-only'                                    
 # end                                                                                                    
 