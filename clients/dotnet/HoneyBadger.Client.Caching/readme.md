## HoneyBadger.Client.Caching
Honey Badger is simple and fast cache server with persistent storage build on top of [BadgerDB](https://github.com/dgraph-io/badger). It uses [gRPC](https://grpc.io/) as transport protocol.
Check how it works [here](https://github.com/meeron/honey-badger)

Usage
```
services.AddHoneyBadgerDistributedCache("127.0.0.1:18950", "cache-db");
```
