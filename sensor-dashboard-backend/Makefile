APP_NAME=sensor-dashboard
BUILD_DIR=bin
SRC_DIR=./cmd
CONFIG_FILE=./config.yaml

GO=go
GOBUILD=$(GO) build
GORUN=$(GO) run
GOFMT=$(GO) fmt
GOTEST=$(GO) test ./...

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

run:
	$(GORUN) $(SRC_DIR) cfg=$(CONFIG_FILE)

fmt:
	$(GOFMT) ./...



clean:
	rm -rf $(BUILD_DIR)
