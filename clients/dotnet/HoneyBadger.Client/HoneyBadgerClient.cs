using Grpc.Net.Client;
using HoneyBadger.Client.Hb;
using HoneyBadger.Client.Internal;

namespace HoneyBadger.Client;

public class HoneyBadgerClient : IHoneyBadgerClient
{
    private readonly Sys.SysClient _sysClient;
    private readonly GrpcChannel _channel;
    
    public HoneyBadgerClient(string address)
    {
        _channel = GrpcChannel.ForAddress(NormalizeAddress(address));
        _sysClient = new Sys.SysClient(_channel);
        
        Data = new HoneyBadgerData(_channel);
        Db = new HoneyBadgerDb(_channel);
    }
    
    public IHoneyBadgerData Data { get; }
    
    public IHoneyBadgerDb Db { get; }

    public async Task<string> PingAsync()
    {
        var res = await _sysClient.PingAsync(new PingRequest());
        return res.Mesage;
    }

    public void Dispose()
    {
        _channel.Dispose();
    }

    private static string NormalizeAddress(string address)
    {
        if (!address.StartsWith("http://"))
        {
            address = $"http://{address}";
        }

        return address;
    }
}
