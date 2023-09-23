using Google.Protobuf;
using Grpc.Core;
using Hb;

namespace HoneyBadger.Client.Internal;

internal class HoneyBadgerData : IHoneyBadgerData
{
    private readonly Data.DataClient _dataClient;

    internal HoneyBadgerData(ChannelBase channel)
    {
        _dataClient = new Data.DataClient(channel);
    }

    public Task<byte[]?> GetAsync(string db, string key) =>
        GetAsync<byte[]>(db, key, data => data.ToByteArray());

    public Task<string?> GetStringAsync(string db, string key) =>
        GetAsync<string>(db, key, data => data.ToStringUtf8());

    public Task<StatusCode> SetAsync(string db, string key, byte[] data, TimeSpan? ttl = null) =>
        SetAsync(db, key, ByteString.CopyFrom(data), ttl);

    public Task<StatusCode> SetStringAsync(string db, string key, string data, TimeSpan? ttl = null) =>
        SetAsync(db, key, ByteString.CopyFromUtf8(data), ttl);

    private async Task<StatusCode> SetAsync(string db, string key, ByteString data, TimeSpan? ttl = null)
    {
        var res = await _dataClient.SetAsync(new SetRequest
        {
            Db = db,
            Key = key,
            Data = data,
            Ttl = (int)Math.Round(ttl?.TotalSeconds ?? 0),
        });

        return res.Code.ToStatusCode();
    }
    
    private async Task<T?> GetAsync<T>(string db, string key, Func<ByteString, T> converter)
        where T: class
    {
        var res = await _dataClient.GetAsync(new KeyRequest
        {
            Db = db,
            Key = key,
        });

        return res.Hit
            ? converter(res.Data)
            : null;
    }
}
