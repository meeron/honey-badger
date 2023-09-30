dev:
	go run .

build:
	cd src && go build -o ../bin/hb -ldflags "-X main.version=$(ver)" .
	cp ./src/config.json ./bin/config.json

run: build
	./bin/hb

proto:
	protoc --go_out=./src/pb --go_opt=paths=source_relative --go-grpc_out=./src/pb --go-grpc_opt=paths=source_relative honey_badger.proto

bench: build
	./bin/hb -bench 127.0.0.1:18950

test:
	cd src && go test ./... -v -race

docker:
	docker build --build-arg ver=$(ver) -t meeron/honey-badger:$(ver) -t meeron/honey-badger:latest .

build-dotnet-client:
	dotnet restore ./clients/dotnet/HoneyBadger.Client.sln
	dotnet build ./clients/dotnet/HoneyBadger.Client.sln -c $(c)

test-dotnet-client: build-dotnet-client
	dotnet test ./clients/dotnet/HoneyBadger.Client.sln

pack-dotnet-client: build-dotnet-client
	dotnet pack ./clients/dotnet/HoneyBadger.Client.sln --no-build -o ./nugets -c $(c)
