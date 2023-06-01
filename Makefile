dev:
	cd cmd; go run main.go

build:
	cd cmd; go build -o ../build/app

start:
	cd build; ./app