test: 
	go test -v ./src/YourLife
	
build:
	go build ./src/YourLife/
	mv YourLife src/YourLife

run: 
	go run ./src/YourLife

travis:
	docker run -t -v $TRAVIS_BUILD_DIR:/test juanalberto58/appiv

