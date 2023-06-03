.PHONY: docs build

dev:
	go run main.go

docs:
	swag init -o ./pkg/docs

build:
	go build -o ./build/app

start:
	cd build; ./app