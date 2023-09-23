namespace HoneyBadger.Client;

public interface IHoneyBadgerData
{
    Task<byte[]?> GetAsync(string db, string key);

    Task<string?> GetStringAsync(string db, string key);

    Task<StatusCode> SetAsync(string db, string key, byte[] data, TimeSpan? ttl = null);
    
    Task<StatusCode> SetStringAsync(string db, string key, string data, TimeSpan? ttl = null);
}
