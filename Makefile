TAG=1.0.0

.PHONY: dist

dist:
	GOOS=linux GOARCH=amd64 go build -o pagerduty-cli_$(TAG)_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o pagerduty-cli_$(TAG)_darwin_amd64