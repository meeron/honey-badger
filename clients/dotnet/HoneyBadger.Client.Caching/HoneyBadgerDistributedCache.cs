using Microsoft.Extensions.Caching.Distributed;

namespace HoneyBadger.Client.Caching;

internal class HoneyBadgerDistributedCache : IDistributedCache
{
    private readonly IHoneyBadgerClient _client;
    private readonly string _db;
    
    internal HoneyBadgerDistributedCache(string address, string db)
    {
        _db = db;
        _client = new HoneyBadgerClient(address);
    }

    public byte[]? Get(string key) => _client.Data.Get(_db, key);

    public Task<byte[]?> GetAsync(string key, CancellationToken token = new CancellationToken()) =>
        _client.Data.GetAsync(_db, key, token);

    public void Set(string key, byte[] value, DistributedCacheEntryOptions options) =>
        _client.Data.Set(_db, key, value, ToTtlTimeSpan(options));

    public Task SetAsync(string key, byte[] value, DistributedCacheEntryOptions options,
        CancellationToken token = new CancellationToken()) =>
        _client.Data.SetAsync(_db, key, value, ToTtlTimeSpan(options), token);

    public void Refresh(string key) =>
        throw new NotSupportedException("'Refresh' is not supported with HoneyBadger.Client");

    public Task RefreshAsync(string key, CancellationToken token = new CancellationToken()) =>
        throw new NotSupportedException("'Refresh' is not supported with HoneyBadger.Client");

    public void Remove(string key) => _client.Data.Delete(_db, key);

    public Task RemoveAsync(string key, CancellationToken token = new CancellationToken()) =>
        _client.Data.DeleteAsync(_db, key, token);

    private static TimeSpan? ToTtlTimeSpan(DistributedCacheEntryOptions options) =>
        options.AbsoluteExpiration.HasValue
            ? options.AbsoluteExpiration.Value - DateTimeOffset.Now
            : options.SlidingExpiration;
}
