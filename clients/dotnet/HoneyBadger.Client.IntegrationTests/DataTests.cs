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

        var data = new List<KeyValuePair<string, string>>();
        await foreach (var item in _data.ReadStringAsync(Db, prefix))
        {
            data.Add(item);
        }

        // Assert
        data.Count.ShouldBe(0);
    }

    [Fact]
    public async Task SendWithStream()
    {
        // Arrange
        const string key = "test-stream";
        const string data = "this is test data";
        
        // Act
        var stream = await _data.CreateSendStream(Db);
        await stream.Write(key, data);
        await stream.Close();

        var resultData = await _data.GetStringAsync(Db, key);
        
        // Assert
        resultData.ShouldBe(data);
    }

    [Fact]
    public async Task ReadStringAsync()
    {
        // Arrange
        var result = new Dictionary<string, string>();
        
        await _data.SetAsync(Db, "read-string-1", "test data 1");
        await _data.SetAsync(Db, "read-string-2", "test data 2");
        
        // Act
        await foreach (var item in _data.ReadStringAsync(Db, "read-string-"))
        {
            result.Add(item.Key, item.Value);
        }
        
        // Assert
        result.Count.ShouldBe(2);
        result["read-string-1"].ShouldBe("test data 1");
        result["read-string-2"].ShouldBe("test data 2");
    }
    
    [Fact]
    public async Task ReadAsync()
    {
        // Arrange
        var result = new Dictionary<string, byte[]>();
        
        await _data.SetAsync(Db, "read-bytes-1", new byte[] { 1 });
        await _data.SetAsync(Db, "read-bytes-2", new byte[] { 2 });
        
        // Act
        await foreach (var item in _data.ReadAsync(Db, "read-bytes-"))
        {
            result.Add(item.Key, item.Value);
        }
        
        // Assert
        result.Count.ShouldBe(2);
        result["read-bytes-1"].ShouldBe(new byte[] { 1 });
        result["read-bytes-2"].ShouldBe(new byte[] { 2 });
    }
}
