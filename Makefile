test:
	@go test ./...

deps-get:
	@go get github.com/spf13/cobra
	@go get github.com/spf13/pflag
	@go get github.com/stretchr/testify/assert

.PHONY: test deps-get
