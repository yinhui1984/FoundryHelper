APPNAME := foundryhelper
CURRENT_DIR := $(shell pwd)

all: clean build run

build:
	go build -o ./bin/$(APPNAME) ./src/
	cp ./config.json ./bin/config.json

clean:
	rm -rf ./bin/*

run:
	cd ~/Downloads && $(CURRENT_DIR)/bin/$(APPNAME)