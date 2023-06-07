COLOUR_BLUE=\033[0;34m
COLOUR_END=\033[0m

docs:
	@go get -u golang.org/x/tools/cmd/godoc
	@echo "please refer this link to see the docs: $(COLOUR_BLUE)http://localhost:6060/pkg/weavestore/$(COLOUR_END)"
	@godoc -http=:6060

test:
	go test -v ./...

setup:
	go install github.com/kisielk/errcheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@2022.1

check:
	python3 pre-commit-3.0.3.pyz run --all-files