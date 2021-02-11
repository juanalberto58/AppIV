test: 
	go test -v ./src/YourLife
	go test -v ./src/
	
build:
	go build ./src/YourLife/
	go build ./api/
	mv YourLife src/YourLife

run: 
	go run ./src/main.go

