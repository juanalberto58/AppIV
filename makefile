test: 
	go test -v ./src/YourLife
	
build:
	go build ./src/YourLife/
	mv YourLife src/YourLife

run: 
	go run ./src/YourLife

travis:
	docker run -t -v `pwd`:/test juanalberto58/appiv

