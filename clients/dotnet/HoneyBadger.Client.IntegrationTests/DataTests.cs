using Shouldly;

namespace HoneyBadger.Client.IntegrationTests;

public class DataTests
{
    private const string Db = "dotnet-client";
    
    private readonly IHoneyBadgerData _data;

    public DataTests()
    {
        _data = new HoneyBadgerClient("127.0.0.1:18950").Data;
    }

    [Fact]
    public async Task SetGetByteArrayData()
    {
        // Arrange
        const string key = "byte[]";
        var data = new byte[] { 1, 2, 3 };
        
        // Act
        await _data.SetAsync(Db, key, data);
        var dbData = await _data.GetAsync(Db, key);
        
        // Assert
        dbData.ShouldNotBeNull();
        dbData.ShouldBe(data);
    }
    
    [Fact]
    public async Task SetGetStringData()
    {
        // Arrange
        const string key = "string";
        var data = "string";
        
        // Act
        await _data.SetAsync(Db, key, data);
        var dbData = await _data.GetStringAsync(Db, key);
        
        // Assert
        dbData.ShouldNotBeNull();
        dbData.ShouldBe(data);
    }
    
    [Fact]
    public async Task SetWithTtl()
    {
        // Arrange
        const string key = "string";
        const string data = "with-ttl";
        
        // Act
        await _data.SetAsync(Db, key, data, TimeSpan.FromSeconds(1));
        await Task.Delay(1500);
        var dbData = await _data.GetStringAsync(Db, key);
        
        // Assert
        dbData.ShouldBeNull();
    }
    
    [Fact]
    public async Task Delete()
    {
        // Arrange
        const string key = "string";
        const string data = "will-be-deleted";
        
        // Act
        await _data.SetAsync(Db, key, data);
        await _data.DeleteAsync(Db, key);
        var dbData = await _data.GetStringAsync(Db, key);
        
        // Assert
        dbData.ShouldBeNull();
    }

    [Fact]
    public async Task GetByPrefix()
    {
        // Arrange
        const string prefix = "prefixed-";
        const int count = 3;

        for (var i = 0; i < count; i++)
        {
            await _data.SetAsync(Db, $"{prefix}{i}", $"data {i}");
        }
        
        // Act
        var data = await _data.GetStringsByPrefixAsync(Db, prefix);
        
        // Assert
        data.Count.ShouldBe(count);
        
        for (var i = 0; i < count; i++)
        {
            data[$"{prefix}{i}"].ShouldBe($"data {i}");
        }
    }
    
    [Fact]
    public async Task DeleteByPrefix()
    {
        // Arrange
        const string prefix = "prefixed-";

        for (var i = 0; i < 3; i++)
        {
            await _data.SetAsync(Db, $"{prefix}{i}", $"data {i}");
        }
        
        // Act
        await _data.DeleteByPrefixAsync(Db, prefix);
        var data = await _data.GetStringsByPrefixAsync(Db, prefix);
        
        // Assert
        data.Count.ShouldBe(0);
    }

    [Fact]
    public async Task GetSetByteArrayBatch()
    {
        // Arrange
        await _data.SetBatchAsync(Db, new Dictionary<string, byte[]>
        {
            { "batch-b-1", new byte[] { 1, 2, 3 } },
            { "batch-b-2", new byte[] { 4, 5, 6 } },
            { "batch-b-3", new byte[] { 7, 8, 9 } }
        });
        
        // Act
        var data = await _data.GetByPrefixAsync(Db, "batch-b-");
        
        // Assert
        data["batch-b-1"].ShouldBe(new byte[] { 1, 2, 3 });
        data["batch-b-2"].ShouldBe(new byte[] { 4, 5, 6 });
        data["batch-b-3"].ShouldBe(new byte[] { 7, 8, 9 });
    }
    
    [Fact]
    public async Task GetSetStringBatch()
    {
        // Arrange
        await _data.SetBatchAsync(Db, new Dictionary<string, string>
        {
            { "batch-s-1", "data1" },
            { "batch-s-2", "data2" },
            { "batch-s-3", "data3" }
        });
        
        // Act
        var data = await _data.GetStringsByPrefixAsync(Db, "batch-s-");
        
        // Assert
        data["batch-s-1"].ShouldBe("data1");
        data["batch-s-2"].ShouldBe("data2");
        data["batch-s-3"].ShouldBe("data3");
    }
}
