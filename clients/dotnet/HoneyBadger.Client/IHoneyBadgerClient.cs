namespace HoneyBadger.Client;

public interface IHoneyBadgerClient : IDisposable
{
    IHoneyBadgerData Data { get; }
    
    Task<StatusCode> PingAsync();
}
