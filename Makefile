build:
	mkdir -p functions
	cd endpoints/hello-world;go get ./...
	cd endpoints/hello-world;go build -o ../../functions/hello-world
