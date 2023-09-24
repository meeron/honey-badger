// See https://aka.ms/new-console-template for more information

using HoneyBadger.Client;

const string Db = "dotnet-client";

var client = new HoneyBadgerClient("127.0.0.1:18950");

var status = await client.PingAsync();

await client.Data.SetAsync(Db, "byte[]", new byte[] { 1, 2, 3 });
var byteData = await client.Data.GetAsync(Db, "byte[]");

await client.Data.SetAsync(Db, "string", "string data");
var stringData = await client.Data.GetStringAsync(Db, "string");

await client.Data.SetAsync(Db, "with-ttl", "present", TimeSpan.FromSeconds(1));
await Task.Delay(1500);
var withTtlData = await client.Data.GetStringAsync(Db, "with-ttl");

await client.Data.SetAsync(Db, "will-be-deleted", "not important");
await client.Data.DeleteAsync(Db, "will-be-deleted");
var deletedData = await client.Data.GetStringAsync(Db, "will-be-deleted");

await client.Data.SetAsync(Db, "prefixed-1", "data 1");
await client.Data.SetAsync(Db, "prefixed-2", "data 2");
await client.Data.SetAsync(Db, "prefixed-3", "data 3");
var prefixedData = await client.Data.GetStringsByPrefixAsync(Db, "prefixed-");
var prefixedDataDebug = string.Join(",", prefixedData.Select(x => $"<{x.Key},{x.Value}>"));

await client.Data.DeleteByPrefixAsync(Db, "prefixed-");
var deletedPrefixedData = await client.Data.GetByPrefixAsync(Db, "prefixed-");

await client.Data.SetBatchAsync(Db, new Dictionary<string, byte[]>
{
    { "batch-b-1", new byte[] { 1, 2, 3 }},
    { "batch-b-2", new byte[] { 4, 5, 6 } },
    { "batch-b-3", new byte[] { 7, 8, 9 } }
});
var batchedByteData = await client.Data.GetByPrefixAsync(Db, "batch-b");
var batchedByteDataDebug = string.Join(",", batchedByteData.Select(x => $"<{x.Key},{BitConverter.ToString(x.Value)}>"));

await client.Data.SetBatchAsync(Db, new Dictionary<string, string>
{
    { "batch-s-1", "batch data 1" },
    { "batch-s-2", "batch data 2" },
    { "batch-s-3", "batch data 3" }
});
var batchedStringData = await client.Data.GetStringsByPrefixAsync(Db, "batch-s");
var batchedStringDataDebug = string.Join(",", batchedStringData.Select(x => $"<{x.Key},{x.Value}>"));

Console.WriteLine($"Ping: {status}");
Console.WriteLine($"byte[]: {BitConverter.ToString(byteData)}");
Console.WriteLine($"string: {stringData}");
Console.WriteLine($"This should be null (ttl): {withTtlData ?? "<null>"}");
Console.WriteLine($"This should be null (delete): {deletedData ?? "<null>"}");
Console.WriteLine($"Prefixed data: {prefixedDataDebug}");
Console.WriteLine($"Prefixed data count (delete): {deletedPrefixedData.Count}");
Console.WriteLine($"Batched string data: {batchedStringDataDebug}");
Console.WriteLine($"Batched byte[] data: {batchedByteDataDebug}");
