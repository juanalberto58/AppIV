test: 
	go test -v ./src ./src/YourLife
	
build:
	go build ./src/main.go
	
install:

start:
	./main

run: 
	go run ./src/main.go

