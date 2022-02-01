GOCMD=go
GOPLUGIN=$(GOCMD) build -buildmode=plugin
MOD_NAME := $(shell echo `git config --get remote.origin.url | sed 's/https\?:\/\///' | sed 's/.git//'`)

all: build

mod:
	@echo MOD $(MOD_NAME)
	[ -f go.mod ] || go mod init $(MOD_NAME)

deps:
	export GO111MODULE=on
	export GOPROXY=direct
	export GOSUMDB=off
	go mod tidy

build: mod deps
	[ -d bin ] || mkdir bin
	$(GOPLUGIN) -o bin/template.so main.go
