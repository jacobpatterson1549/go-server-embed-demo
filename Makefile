.PHONY: serve clean

GO_ENV := # GOOS=windows GOARCH=amd64
GO := $(GO_ENV) go
GO_SRC := *.go
SERVE_SRC := html/* README.md
OBJ := main # set to main.exe for windows
TEST := main.test
BUILD_DIR := build

$(BUILD_DIR)/$(OBJ): $(GO_SRC) $(SERVE_SRC) $(BUILD_DIR)/$(TEST) | $(BUILD_DIR)
	$(GO) build -o $@ $<

$(BUILD_DIR)/$(TEST): $(GO_SRC) $(SERVE_SRC) | $(BUILD_DIR)
	$(GO) test $< > $@

$(BUILD_DIR):
	mkdir $@

serve: $(BUILD_DIR)/$(OBJ)
	$<

clean:
	rm -r -f $(BUILD_DIR)
