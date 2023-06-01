dev:
	cd cmd; go run main.go

docs:
	swag init --dir ./cmd

build:
	cd cmd; go build -o ../build/app

start:
	cd build; ./app