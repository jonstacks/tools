PLATFORM=`uname`
VERSION=`cat VERSION.txt`

default: help

.PHONY: help
help:
	@echo "deps     - Installs project dependencies"
	@echo "install  - Install command line tools"
	@echo "release  - Packages up the command binaries into a zip file"

deps:
	dep ensure -v

install: deps
	go install ./cmd/jira

release:
	@mkdir -p release
	go build -o release/jira ./cmd/jira
	tar -czvf tools-$(PLATFORM)-$(VERSION).tar.gz ./release
	@rm -rf ./release
