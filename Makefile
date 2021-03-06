install:
	go get -v -t -d ./...

lint:
	golint ./...

test:
	go test -v ./...

benchmark:
	go test -v -bench=. -run=- ./...
