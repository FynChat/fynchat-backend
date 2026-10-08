.PHONY:	build

build:
	go build -buildvcs=false -o bin/fynchat ./cmd/fynchat
