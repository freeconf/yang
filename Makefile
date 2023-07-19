export YANGPATH=$(abspath ./yang)

VER = 0.1.0

test:
	go test ./...

# Just the popular ones.  You can easily build binary for missing platform
PLATFORMS = \
  darwin-amd64 \
  darwin-arm64 \
  linux-amd64 \
  windows-amd64

BIN_TARGETS = $(foreach P,$(PLATFORMS), bin/fc-yang-$(VER)-$(P))

dist: $(BIN_TARGETS)

bin/fc-yang-$(VER)-darwin-amd64: BUILD_ENV=GOARCH=amd64 GOOS=darwin
bin/fc-yang-$(VER)-darwin-arm64: BUILD_ENV=GOARCH=amd64 GOOS=darwin
bin/fc-yang-$(VER)-windows-amd64: BUILD_ENV=GOARCH=amd64 GOOS=windows
bin/fc-yang-$(VER)-windows-amd64: BIN_EXT=.exe
bin/fc-yang-$(VER)-linux-amd64: BUILD_ENV=GOARCH=amd64 GOOS=linux

.PHONY: $(BIN_TARGETS)

$(BIN_TARGETS):
	test -d $(dir $@) || mkdir -p $(dir $@)
	$(BUILD_ENV) go build $(BUILD_OPTS) -o $@$(BIN_EXT) cmd/fc-yang/main.go


# Getting this error. Might have been introduced in 1.20 and no known solution or workaround
# https://github.com/golang/go/issues/45361
test-coverage:
	go test -coverprofile test-coverage.out ./...
	go tool cover -html=test-coverage.out -o test-coverage.html
	go tool cover -func test-coverage.out

