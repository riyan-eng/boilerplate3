dev:
	cd cmd; go run main.go

docs:
	swag init -d ./cmd -o ./pkg/docs

build:
	cd cmd; go build -o ../build/app

start:
	cd build; ./app