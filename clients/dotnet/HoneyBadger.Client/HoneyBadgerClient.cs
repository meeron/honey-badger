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

        Data = new HoneyBadgerData(_channel);
        _sysClient = new Sys.SysClient(_channel);
    }
    
    public IHoneyBadgerData Data { get; }

    public async Task<StatusCode> PingAsync()
    {
        var res = await _sysClient.PingAsync(new PingRequest());
        return res.Code.ToStatusCode();
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
