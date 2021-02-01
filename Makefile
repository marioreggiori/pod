.PHONY: all docs build

all: docs build

docs:
	go run main.go doc-gen

build: 
	go build -o app main.go