test: 
	go test -v ./src/main
	
build:
	go build ./src/main/
	go build ./api/
	mv YourLife src/main

run: 
	go run ./src/main

