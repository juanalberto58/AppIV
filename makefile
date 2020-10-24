build:
	go build ./src/Yourlife 
	mv Yourlife src/Yourlife

run: 
	go run ./src/Yourlife

test: 
	go test ./src/Yourlife