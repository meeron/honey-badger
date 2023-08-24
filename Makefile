dev:
	go run .

build:
	go build -o ./.out/hb .

run: build
	./.out/hb

