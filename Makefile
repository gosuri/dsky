test:
	@go test ./...

deps-get:
	@go get github.com/spf13/cobra
	@go get github.com/spf13/pflag

.PHONY: test deps-get
