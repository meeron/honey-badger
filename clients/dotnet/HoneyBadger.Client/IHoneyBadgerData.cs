namespace HoneyBadger.Client;

public interface IHoneyBadgerData
{
    Task<byte[]?> GetAsync(string db, string key);
    
    byte[]? Get(string db, string key);

    Task<string?> GetStringAsync(string db, string key);

    Task<IReadOnlyDictionary<string, byte[]>> GetByPrefixAsync(string db, string prefix);
    
    Task<IReadOnlyDictionary<string, string>> GetStringsByPrefixAsync(string db, string prefix);

    Task SetAsync(string db, string key, byte[] data, TimeSpan? ttl = null);
    
    void Set(string db, string key, byte[] data, TimeSpan? ttl = null);
    
    Task SetAsync(string db, string key, string data, TimeSpan? ttl = null);

    Task SetBatchAsync(string db, IReadOnlyDictionary<string, byte[]> data);
    
    Task SetBatchAsync(string db, IReadOnlyDictionary<string, string> data);

    Task DeleteAsync(string db, string key);
    
    void Delete(string db, string key);

    Task DeleteByPrefixAsync(string db, string prefix);
}
