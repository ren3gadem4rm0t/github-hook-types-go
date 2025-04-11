APP_NAME := github-hook-types-go
VERSION  := 0.1.0

SHELL := /bin/bash

.PHONY: test lint sec clean major minor patch release coverage coverage-ci vet staticcheck ast fmt check-fmt
.DEFAULT_GOAL := help

test:
	@echo "Running tests..."
	go test -v -coverprofile=coverage.out ./...

lint:
	@echo "Running golangci-lint..."
	golangci-lint run

ast:
	@echo "Running security scan..."
	gosec  -exclude-dir=examples . ./webhook  ./...

fmt:
	@gofmt -s -w .

check-fmt:
	@gofmt -l . | tee /dev/stderr | grep -q '^' && echo "Code is not formatted" && exit 1 || echo "Code is formatted"

coverage:
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out ./... && \
	go tool cover -func=coverage/coverage.out && \
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html; \
	open coverage/coverage.html

coverage-ci:
	@mkdir -p coverage
	@go test -coverprofile=coverage/coverage.out ./... && \
	go tool cover -func=coverage/coverage.out

vet:
	@go vet ./...

staticcheck:
	@staticcheck ./...

clean:
	@echo "Cleaning up..."
	rm -rf bin coverage.out

# Bump version targets – bump the VERSION variable, commit, tag, and push.
define bump
	@old_version=$(VERSION); \
	IFS='.' read -r major minor patch <<< "$$old_version"; \
	if [ "$(1)" = "major" ]; then \
	  major=$$((major+1)); minor=0; patch=0; \
	elif [ "$(1)" = "minor" ]; then \
	  minor=$$((minor+1)); patch=0; \
	elif [ "$(1)" = "patch" ]; then \
	  patch=$$((patch+1)); \
	else \
	  echo "Usage: make {major|minor|patch}"; exit 1; \
	fi; \
	new_version="$$major.$$minor.$$patch"; \
	sed -i.bak "s/^VERSION[[:space:]]*:=[[:space:]].*/VERSION  := $$new_version/" Makefile && rm -f Makefile.bak; \
	git add Makefile; \
	git commit -m "Bump version to $$new_version"; \
	git tag "v$$new_version"; \
	git push origin main --tags; \
	echo "Version bumped from $$old_version to $$new_version."
endef

major:
	$(call bump,major)

minor:
	$(call bump,minor)

patch:
	$(call bump,patch)

release:
	@echo "Releasing with goreleaser..."
	@goreleaser release --clean
	@echo "Release complete."
