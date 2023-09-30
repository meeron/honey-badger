## HoneyBadger.Client
Honey Badger is simple and fast cache server with persistent storage build on top of [BadgerDB](https://github.com/dgraph-io/badger). It uses [gRPC](https://grpc.io/) as transport protocol.
Check how it works [here](https://github.com/meeron/honey-badger).

Check also `IDistributedCache` [implementation](https://www.nuget.org/packages/HoneyBadger.Client.Caching).

Usage
```csharp
var client = new HoneyBadgerClient("127.0.0.1:18950");

await client.Data.SetAsync("db", "key", "data");
var data = client.Data.GetStringAsync("db", "key");
```
