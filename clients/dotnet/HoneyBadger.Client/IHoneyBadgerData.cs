namespace HoneyBadger.Client;

public interface IHoneyBadgerData
{
    Task<byte[]?> GetAsync(string db, string key);
    
    byte[]? Get(string db, string key);

    Task<string?> GetStringAsync(string db, string key);

    Task<IReadOnlyDictionary<string, byte[]>> GetByPrefixAsync(string db, string prefix);
    
    Task<IReadOnlyDictionary<string, string>> GetStringsByPrefixAsync(string db, string prefix);

    Task<StatusCode> SetAsync(string db, string key, byte[] data, TimeSpan? ttl = null);
    
    StatusCode Set(string db, string key, byte[] data, TimeSpan? ttl = null);
    
    Task<StatusCode> SetAsync(string db, string key, string data, TimeSpan? ttl = null);

    Task<StatusCode> SetBatchAsync(string db, IReadOnlyDictionary<string, byte[]> data);
    
    Task<StatusCode> SetBatchAsync(string db, IReadOnlyDictionary<string, string> data);

    Task<StatusCode> DeleteAsync(string db, string key);
    
    StatusCode Delete(string db, string key);

    Task<StatusCode> DeleteByPrefixAsync(string db, string prefix);
}
