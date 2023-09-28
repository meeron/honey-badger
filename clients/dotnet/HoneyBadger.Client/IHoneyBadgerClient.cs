namespace HoneyBadger.Client;

public interface IHoneyBadgerClient : IDisposable
{
    IHoneyBadgerData Data { get; }
    
    IHoneyBadgerDb Db { get; }
    
    Task<string> PingAsync();
}
