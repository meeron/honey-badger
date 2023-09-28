using System.Text;
using HoneyBadger.Client.Caching;
using Microsoft.Extensions.Caching.Distributed;
using Microsoft.Extensions.DependencyInjection;
using Shouldly;

namespace HoneyBadger.Client.IntegrationTests;

public class DistributedCacheTests
{
    private readonly IServiceProvider _sp;

    public DistributedCacheTests()
    {
        var services = new ServiceCollection();
        services.AddHoneyBadgerDistributedCache("127.0.0.1:18950", nameof(DistributedCacheTests));

        _sp = services.BuildServiceProvider();
    }

    [Fact]
    public void GetSetData()
    {
        // Arrange
        const string data = "test data";
        
        var cache = _sp.GetService<IDistributedCache>()!;
        
        // Act
        cache.Set(nameof(GetSetData), Encoding.UTF8.GetBytes(data));
        var result = cache.Get(nameof(GetSetData))!;
        var resultString = Encoding.UTF8.GetString(result);

        // Assert
        resultString.ShouldBe(data);
    }
    
    [Fact]
    public async Task GetSetAsyncData()
    {
        // Arrange
        const string data = "test data";
        
        var cache = _sp.GetService<IDistributedCache>()!;
        
        // Act
        await cache.SetAsync(nameof(GetSetAsyncData), Encoding.UTF8.GetBytes(data));
        var result = await cache.GetAsync(nameof(GetSetAsyncData));
        var resultString = Encoding.UTF8.GetString(result!);

        // Assert
        resultString.ShouldBe(data);
    }
    
    [Fact]
    public async Task GetSetDataWithAbsoluteExpiration()
    {
        // Arrange
        const string data = "test data";
        
        var cache = _sp.GetService<IDistributedCache>()!;
        
        // Act
        await cache.SetAsync(nameof(GetSetDataWithAbsoluteExpiration), Encoding.UTF8.GetBytes(data), new DistributedCacheEntryOptions
        {
            AbsoluteExpiration = DateTimeOffset.Now.AddSeconds(1),
        });
        await Task.Delay(1500);
        var result = await cache.GetAsync(nameof(GetSetDataWithAbsoluteExpiration));

        // Assert
        result.ShouldBeNull();
    }
    
    [Fact]
    public async Task GetSetDataWithSlidingExpiration()
    {
        // Arrange
        const string data = "test data";
        
        var cache = _sp.GetService<IDistributedCache>()!;
        
        // Act
        await cache.SetAsync(nameof(GetSetDataWithSlidingExpiration), Encoding.UTF8.GetBytes(data), new DistributedCacheEntryOptions
        {
            SlidingExpiration = TimeSpan.FromSeconds(1),
        });
        await Task.Delay(1500);
        var result = await cache.GetAsync(nameof(GetSetDataWithSlidingExpiration));

        // Assert
        result.ShouldBeNull();
    }

    [Fact]
    public void Remove()
    {
        // Arrange
        const string data = "test data";
        const string key = nameof(Remove);
        
        var cache = _sp.GetService<IDistributedCache>()!;
        
        // Act
        cache.Set(key, Encoding.UTF8.GetBytes(data));
        cache.Remove(key);
        var result = cache.Get(key);
        
        // Assert
        result.ShouldBeNull();
    }
    
    [Fact]
    public async Task RemoveAsync()
    {
        // Arrange
        const string data = "test data";
        const string key = nameof(RemoveAsync);
        
        var cache = _sp.GetService<IDistributedCache>()!;
        
        // Act
        await cache.SetAsync(key, Encoding.UTF8.GetBytes(data));
        await cache.RemoveAsync(key);
        var result = await cache.GetAsync(key);
        
        // Assert
        result.ShouldBeNull();
    }
}
