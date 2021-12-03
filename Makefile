BUILD_FLAGS = -ldflags="-s -w"
GO_FILES = $(shell find -iname '*.go' -type f)
GO_ENTRY = $(shell find ./cli/kakoune-wiki/ -iname '*.go' -type f)
GO_TEST_FLAGS = -race ./...
GO_COVER_FLAGS = -covermode atomic
OPEN = xdg-open

CLEAN = kakoune-wiki cover.out cover.html

kakoune-wiki: $(GO_FILES)
	go build $(BUILD_FLAGS) -o $@ $(GO_ENTRY)

.PHONY: test
test: $(GO_FILES)
	go test $(GO_TEST_FLAGS)

cover.out: $(GO_FILES) test
	go test $(GO_TEST_FLAGS) -coverprofile $@ $(GO_COVER_FLAGS)

cover.html: cover.out
	go tool cover -html=$^ -o $@

coverage: cover.html
	$(OPEN) $^

.PHONY: lint
lint:
	golangci-lint run

.PHONY: format
format: $(GO_FILES)
	goimports -w $^

.PHONY: clean
clean:
	rm -f $(CLEAN)
