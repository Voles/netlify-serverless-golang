build:
	mkdir -p functions
	cd go-src/hello-world;go get ./...
	cd go-src/hello-world; go build -o ../../functions/hello-world
