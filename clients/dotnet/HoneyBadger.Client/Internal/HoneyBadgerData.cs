using Google.Protobuf;
using Grpc.Core;
using HoneyBadger.Client.Hb;

namespace HoneyBadger.Client.Internal;

internal class HoneyBadgerData : IHoneyBadgerData
{
    private readonly Data.DataClient _dataClient;

    internal HoneyBadgerData(ChannelBase channel)
    {
        _dataClient = new Data.DataClient(channel);
    }

    public Task<byte[]?> GetAsync(string db, string key)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        
        return GetAsync<byte[]>(db, key, data => data.ToByteArray());
    }

    public Task<string?> GetStringAsync(string db, string key)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        
        return GetAsync<string>(db, key, data => data.ToStringUtf8());
    }

    public Task<IReadOnlyDictionary<string, byte[]>> GetByPrefixAsync(string db, string prefix)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(prefix), prefix);
        
        return GetByPrefix(db, prefix, data => data.ToByteArray());
    }

    public Task<IReadOnlyDictionary<string, string>> GetStringsByPrefixAsync(string db, string prefix)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(prefix), prefix);
        
        return GetByPrefix(db, prefix, data => data.ToStringUtf8());
    }

    public Task<StatusCode> SetAsync(string db, string key, byte[] data, TimeSpan? ttl = null)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        Guard.NotNull(nameof(data), data);
        
        return SetAsync(db, key, ByteString.CopyFrom(data), ttl);
    }

    public Task<StatusCode> SetAsync(string db, string key, string data, TimeSpan? ttl = null)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        Guard.NotNullOrEmpty(nameof(data), data);
        
        return SetAsync(db, key, ByteString.CopyFromUtf8(data), ttl);
    }

    public Task<StatusCode> SetBatchAsync(string db, IReadOnlyDictionary<string, byte[]> data)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNull(nameof(data), data);
        
        return SetBatchAsync(db, data, ByteString.CopyFrom);
    }

    public Task<StatusCode> SetBatchAsync(string db, IReadOnlyDictionary<string, string> data)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNull(nameof(data), data);
        
        return SetBatchAsync(db, data, ByteString.CopyFromUtf8);
    }

    public async Task<StatusCode> DeleteAsync(string db, string key)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        
        var res = await _dataClient.DeleteAsync(new KeyRequest
        {
            Db = db,
            Key = key,
        });
        return res.Code.ToStatusCode();
    }

    public async Task<StatusCode> DeleteByPrefixAsync(string db, string prefix)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(prefix), prefix);
        
        var res = await _dataClient.DeleteByPrefixAsync(new PrefixRequest
        {
            Db = db,
            Prefix = prefix
        });
        return res.Code.ToStatusCode();
    }

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
    
    private async Task<IReadOnlyDictionary<string, T>> GetByPrefix<T>(string db, string prefix, Func<ByteString, T> converter)
    {
        var res = await _dataClient.GetByPrefixAsync(new PrefixRequest
        {
            Db = db,
            Prefix = prefix,
        });

        return res.Data.ToDictionary(k => k.Key, v => converter(v.Value));
    }
    
    private async Task<StatusCode> SetBatchAsync<T>(
        string db,
        IReadOnlyDictionary<string, T> data,
        Func<T, ByteString> converter)
    {
        var req = new SetBatchRequest
        {
            Db = db,
        };
        req.Data.MergeFrom(data.ToDictionary(k => k.Key, v => converter(v.Value)));
        
        var res = await _dataClient.SetBatchAsync(req);
        return res.Code.ToStatusCode();
    }
}
