test: 
	go test -v ./src/YourLife
	
build:
	go build ./src/YourLife/
	mv YourLife src/YourLife

run: 
	go run ./src/YourLife


