test: 
	go test -v ./src ./src/YourLife
	
build:
	go build ./src/YourLife/
	go build ./api/
	mv YourLife src/YourLife
	
install:

run: 
	go run ./src/main.go

