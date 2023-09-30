namespace HoneyBadger.Client;

public interface IHoneyBadgerData
{
    Task<byte[]?> GetAsync(string db, string key, CancellationToken ct = default);
    
    byte[]? Get(string db, string key);

    Task<string?> GetStringAsync(string db, string key, CancellationToken ct = default);

    Task<IReadOnlyDictionary<string, byte[]>> GetByPrefixAsync(string db, string prefix);
    
    Task<IReadOnlyDictionary<string, string>> GetStringsByPrefixAsync(string db, string prefix);

    Task SetAsync(string db, string key, byte[] data, TimeSpan? ttl = null, CancellationToken ct = default);
    
    void Set(string db, string key, byte[] data, TimeSpan? ttl = null);
    
    Task SetAsync(string db, string key, string data, TimeSpan? ttl = null, CancellationToken ct = default);

    Task SetBatchAsync(string db, IReadOnlyDictionary<string, byte[]> data);
    
    Task SetBatchAsync(string db, IReadOnlyDictionary<string, string> data);

    Task DeleteAsync(string db, string key, CancellationToken ct = default);
    
    void Delete(string db, string key);

    Task DeleteByPrefixAsync(string db, string prefix, CancellationToken ct = default);
}
