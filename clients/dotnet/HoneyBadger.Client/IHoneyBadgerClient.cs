namespace HoneyBadger.Client;

public interface IHoneyBadgerClient : IDisposable
{
    Task<byte[]?> GetAsync(string db, string key);

    Task<string?> GetStringAsync(string db, string key);

    Task<StatusCode> SetAsync(string db, string key, byte[] data);
    
    Task<StatusCode> SetStringAsync(string db, string key, string data);

    Task<StatusCode> PingAsync();
}
