// See https://aka.ms/new-console-template for more information

using HoneyBadger.Client;

const string Db = "dotnet-client";

var client = new HoneyBadgerClient("127.0.0.1:18950");

// Ping
var status = await client.PingAsync();
Console.WriteLine($"Ping: {status}");

// Check set byte[] data
await client.Data.SetAsync(Db, "byte[]", new byte[] { 1, 2, 3 });
var byteData = await client.Data.GetAsync(Db, "byte[]");
Console.WriteLine($"byte[]: {BitConverter.ToString(byteData!)}");

// Check string data
await client.Data.SetAsync(Db, "string", "string data");
var stringData = await client.Data.GetStringAsync(Db, "string");
Console.WriteLine($"string: {stringData}");

// Check TTL work
await client.Data.SetAsync(Db, "with-ttl", "present", TimeSpan.FromSeconds(1));
await Task.Delay(1500);
var withTtlData = await client.Data.GetStringAsync(Db, "with-ttl");
Console.WriteLine($"This should be null (ttl): {withTtlData ?? "<null>"}");

// Check deleting data
await client.Data.SetAsync(Db, "will-be-deleted", "not important");
await client.Data.DeleteAsync(Db, "will-be-deleted");
var deletedData = await client.Data.GetStringAsync(Db, "will-be-deleted");
Console.WriteLine($"This should be null (delete): {deletedData ?? "<null>"}");

// Check getting and deleting by prefix
await client.Data.SetAsync(Db, "prefixed-1", "data 1");
await client.Data.SetAsync(Db, "prefixed-2", "data 2");
await client.Data.SetAsync(Db, "prefixed-3", "data 3");
var prefixedData = await client.Data.GetStringsByPrefixAsync(Db, "prefixed-");
var prefixedDataDebug = string.Join(",", prefixedData.Select(x => $"<{x.Key},{x.Value}>"));
Console.WriteLine($"Prefixed data: {prefixedDataDebug}");

await client.Data.DeleteByPrefixAsync(Db, "prefixed-");
var deletedPrefixedData = await client.Data.GetByPrefixAsync(Db, "prefixed-");
Console.WriteLine($"Prefixed data count (delete): {deletedPrefixedData.Count}");

// Check set byte[] data in batch mode
await client.Data.SetBatchAsync(Db, new Dictionary<string, byte[]>
{
    { "batch-b-1", new byte[] { 1, 2, 3 }},
    { "batch-b-2", new byte[] { 4, 5, 6 } },
    { "batch-b-3", new byte[] { 7, 8, 9 } }
});
var batchedByteData = await client.Data.GetByPrefixAsync(Db, "batch-b");
var batchedByteDataDebug = string.Join(",", batchedByteData.Select(x => $"<{x.Key},{BitConverter.ToString(x.Value)}>"));
Console.WriteLine($"Batched byte[] data: {batchedByteDataDebug}");

// Check set string data in batch mode
await client.Data.SetBatchAsync(Db, new Dictionary<string, string>
{
    { "batch-s-1", "batch data 1" },
    { "batch-s-2", "batch data 2" },
    { "batch-s-3", "batch data 3" }
});
var batchedStringData = await client.Data.GetStringsByPrefixAsync(Db, "batch-s");
var batchedStringDataDebug = string.Join(",", batchedStringData.Select(x => $"<{x.Key},{x.Value}>"));
Console.WriteLine($"Batched string data: {batchedStringDataDebug}");

// Check create db
var createStatus = await client.Db.Create("new-db", true);
Console.WriteLine($"Db created: {createStatus}");

// Check delete db
var createDelete = await client.Db.Drop("new-db");
Console.WriteLine($"Db deleted: {createDelete}");
