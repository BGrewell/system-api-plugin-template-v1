GOCMD=go
GOPLUGIN=$(GOCMD) build -buildmode=plugin

build:
  [ -d bin ] || mkdir bin
  $(GOPLUGIN) -o bin/template.so main.go
