namespace HoneyBadger.Client;

public interface IHoneyBadgerData
{
    Task<byte[]?> GetAsync(string db, string key, CancellationToken ct = default);
    
    byte[]? Get(string db, string key);

    Task<string?> GetStringAsync(string db, string key, CancellationToken ct = default);

    Task SetAsync(string db, string key, byte[] data, TimeSpan? ttl = null, CancellationToken ct = default);
    
    void Set(string db, string key, byte[] data, TimeSpan? ttl = null);
    
    Task SetAsync(string db, string key, string data, TimeSpan? ttl = null, CancellationToken ct = default);

    Task DeleteAsync(string db, string key, CancellationToken ct = default);
    
    void Delete(string db, string key);

    Task DeleteByPrefixAsync(string db, string prefix, CancellationToken ct = default);

    Task<SendStream> CreateSendStream(string db);

    IAsyncEnumerable<KeyValuePair<string, byte[]>> ReadAsync(string db, string prefix);
    
    IAsyncEnumerable<KeyValuePair<string, string>> ReadStringAsync(string db, string prefix);
}
