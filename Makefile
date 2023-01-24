install:
	go get ./cmd/practice-server

build: install
	go build -o bin/practice-server ./cmd/practice-server

run: build
	./bin/practice-server