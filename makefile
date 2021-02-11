test: 
	go test -v ./src/YourLife ./src/
	
build:
	go build ./src/YourLife/
	go build ./api/
	mv YourLife src/YourLife

run: 
	go run ./src/main.go

