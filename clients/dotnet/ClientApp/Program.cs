// See https://aka.ms/new-console-template for more information

using HoneyBadger.Client;

const string Db = "dotnet-client";

var client = new HoneyBadgerClient("127.0.0.1:18950");

var status = await client.PingAsync();

await client.Data.SetAsync(Db, "byte[]", new byte[] { 1, 2, 3 });
var byteData = await client.Data.GetAsync(Db, "byte[]");

await client.Data.SetStringAsync(Db, "string", "string data");
var stringData = await client.Data.GetStringAsync(Db, "string");

await client.Data.SetStringAsync(Db, "with-ttl", "present", TimeSpan.FromSeconds(1));
await Task.Delay(1500);
var withTtlData = await client.Data.GetStringAsync(Db, "with-ttl");

Console.WriteLine($"Ping: {status}");
Console.WriteLine($"byte[]: {BitConverter.ToString(byteData)}");
Console.WriteLine($"string: {stringData}");
Console.WriteLine($"This should be null: {withTtlData ?? "<null>"}");
