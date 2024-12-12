.PHONY: build run

all: build run

build:
	go build -o ./tmp/priceindex .

run:
	./tmp/priceindex

migrate:
	./tmp/priceindex migrate