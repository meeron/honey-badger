using Shouldly;

namespace HoneyBadger.Client.IntegrationTests;

public class ClientTests
{
    private readonly IHoneyBadgerClient _client;

    public ClientTests()
    {
        _client = new HoneyBadgerClient("127.0.0.1:18950");
    }
    
    [Fact]
    public async Task Ping()
    {
        // Act
        var result = await _client.PingAsync();
        
        result.ShouldBe("pong");
    }
}
