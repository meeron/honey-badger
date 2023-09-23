using Google.Protobuf;
using Grpc.Net.Client;
using Hb;
using HoneyBadger.Client.Internal;

namespace HoneyBadger.Client;

public class HoneyBadgerClient : IHoneyBadgerClient
{
    private readonly Data.DataClient _dataClient;
    private readonly Sys.SysClient _sysClient;
    private readonly GrpcChannel _channel;
    
    public HoneyBadgerClient(string address)
    {
        _channel = GrpcChannel.ForAddress(NormalizeAddress(address));

        _dataClient = new Data.DataClient(_channel);
        _sysClient = new Sys.SysClient(_channel);
    }
    
    public async Task<byte[]?> GetAsync(string db, string key)
    {
        var res = await _dataClient.GetAsync(new KeyRequest
        {
            Db = db,
            Key = key,
        });

        return res.Hit
            ? res.Data.ToByteArray()
            : null;
    }

    public async Task<string?> GetStringAsync(string db, string key)
    {
        var res = await _dataClient.GetAsync(new KeyRequest
        {
            Db = db,
            Key = key,
        });

        return res.Hit
            ? res.Data.ToStringUtf8()
            : null;
    }

    public async Task<StatusCode> SetAsync(string db, string key, byte[] data)
    {
        var res = await _dataClient.SetAsync(new SetRequest
        {
            Db = db,
            Key = key,
            Data = ByteString.CopyFrom(data),
        });

        return res.Code.ToStatusCode();
    }

    public async Task<StatusCode> SetStringAsync(string db, string key, string data)
    {
        var res = await _dataClient.SetAsync(new SetRequest
        {
            Db = db,
            Key = key,
            Data = ByteString.CopyFromUtf8(data),
        });

        return res.Code.ToStatusCode();
    }

    public async Task<StatusCode> PingAsync()
    {
        var res = await _sysClient.PingAsync(new PingRequest());
        return res.Code.ToStatusCode();
    }

    public void Dispose()
    {
        _channel.Dispose();
    }

    private static string NormalizeAddress(string address)
    {
        if (!address.StartsWith("http://"))
        {
            address = $"http://{address}";
        }

        return address;
    }
}
