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

    public Task<byte[]?> GetAsync(string db, string key, CancellationToken ct = default)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        
        return GetAsync<byte[]>(db, key, data => data.ToByteArray(), ct);
    }

    public byte[]? Get(string db, string key)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        
        var res = _dataClient.Get(new KeyRequest
        {
            Db = db,
            Key = key,
        });

        return res.Hit ? res.Data.ToByteArray() : null;
    }

    public Task<string?> GetStringAsync(string db, string key, CancellationToken ct = default)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        
        return GetAsync<string>(db, key, data => data.ToStringUtf8(), ct);
    }

    public Task<IReadOnlyDictionary<string, byte[]>> GetByPrefixAsync(string db, string prefix)
    {
        throw new NotImplementedException();
    }

    public Task<IReadOnlyDictionary<string, string>> GetStringsByPrefixAsync(string db, string prefix)
    {
        throw new NotImplementedException();
    }

    public Task SetAsync(string db, string key, byte[] data, TimeSpan? ttl = null, CancellationToken ct = default)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        Guard.NotNull(nameof(data), data);
        
        return SetAsync(db, key, ByteString.CopyFrom(data), ct, ttl);
    }

    public void Set(string db, string key, byte[] data, TimeSpan? ttl = null)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        Guard.NotNull(nameof(data), data);

        _dataClient.Set(new SetRequest
        {
            Db = db,
            Key = key,
            Data = ByteString.CopyFrom(data),
            Ttl = (int)Math.Round(ttl?.TotalSeconds ?? 0),
        });
    }

    public Task SetAsync(string db, string key, string data, TimeSpan? ttl = null, CancellationToken ct = default)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        Guard.NotNullOrEmpty(nameof(data), data);
        
        return SetAsync(db, key, ByteString.CopyFromUtf8(data), ct, ttl);
    }

    public Task SetBatchAsync(string db, IReadOnlyDictionary<string, byte[]> data)
    {
        throw new NotImplementedException();
    }

    public Task SetBatchAsync(string db, IReadOnlyDictionary<string, string> data)
    {
        throw new NotImplementedException();
    }

    public async Task DeleteAsync(string db, string key, CancellationToken ct = default)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);
        
        await _dataClient.DeleteAsync(new KeyRequest
        {
            Db = db,
            Key = key,
        }, cancellationToken: ct);
    }

    public void Delete(string db, string key)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(key), key);

        _dataClient.Delete(new KeyRequest
        {
            Db = db,
            Key = key,
        });
    }

    public async Task DeleteByPrefixAsync(string db, string prefix, CancellationToken ct = default)
    {
        Guard.NotNullOrEmpty(nameof(db), db);
        Guard.NotNullOrEmpty(nameof(prefix), prefix);
        
        await _dataClient.DeleteByPrefixAsync(new PrefixRequest
        {
            Db = db,
            Prefix = prefix
        }, cancellationToken: ct);
    }

    private async Task SetAsync(string db, string key, ByteString data, CancellationToken ct, TimeSpan? ttl = null)
    {
        await _dataClient.SetAsync(new SetRequest
        {
            Db = db,
            Key = key,
            Data = data,
            Ttl = (int)Math.Round(ttl?.TotalSeconds ?? 0),
        }, cancellationToken: ct);
    }
    
    private async Task<T?> GetAsync<T>(string db, string key, Func<ByteString, T> converter, CancellationToken ct)
        where T: class
    {
        var res = await _dataClient.GetAsync(new KeyRequest
        {
            Db = db,
            Key = key,
        }, cancellationToken: ct);

        return res.Hit
            ? converter(res.Data)
            : null;
    }
    
    private async Task<IReadOnlyDictionary<string, T>> GetByPrefix<T>(string db, string prefix, Func<ByteString, T> converter)
    {
        throw new NotImplementedException();
    }
    
    private async Task SetBatchAsync<T>(
        string db,
        IReadOnlyDictionary<string, T> data,
        Func<T, ByteString> converter)
    {
        throw new NotImplementedException();
    }
}
