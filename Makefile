# Makefile to build questlister

APP_NAME := ql

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
DB_DIR := ~/.config/questlister
BIN_DIR := ~/bin/

.PHONY: all build clean install fmt

all: build

build:
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$(APP_NAME) $(GOBASE)/cmd

fmt:
	@go fmt ./...

install:
	mkdir -p $(DB_DIR)
	mkdir -p $(BIN_DIR)
	cp $(GOBIN)/$(APP_NAME) $(BIN_DIR)
	
clean:
	rm -f $(GOBASE)/bin/*
