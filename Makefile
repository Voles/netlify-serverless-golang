build:
	mkdir -p functions
	cd endpoints/items;go get ./...
	cd endpoints/items;go build -o ../../functions/items
