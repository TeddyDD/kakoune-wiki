BUILD_FLAGS = -ldflags="-s -w"
GO_FILES = $(shell find -iname '*.go' -type f)

kakoune-wiki: $(GO_FILES)
	go build $(BUILD_FLAGS) -o $@ $^

.PHONY: format
format: $(GO_FILES)
	goimports -w $^

.PHONY: clean
clean:
	rm -f kakoune-wiki
