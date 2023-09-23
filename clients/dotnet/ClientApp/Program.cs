// See https://aka.ms/new-console-template for more information

using HoneyBadger.Client;

var client = new HoneyBadgerClient("127.0.0.1:18950");

var status = await client.PingAsync();

Console.WriteLine($"Status: {status}");
