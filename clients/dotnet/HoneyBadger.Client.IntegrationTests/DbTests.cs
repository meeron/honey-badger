using Grpc.Core;
using Shouldly;

namespace HoneyBadger.Client.IntegrationTests;

public class DbTests
{
    private readonly IHoneyBadgerDb _db;

    public DbTests()
    {
        _db = new HoneyBadgerClient("127.0.0.1:18950").Db;
    }

    [Fact]
    public async Task CreateInMemoryDb()
    {
        // Act
        await _db.Create("in-memory-db", true);
        await _db.Drop("in-memory-db");
    }
    
    [Fact]
    public async Task CreateOnDiskDb()
    {
        // Act
        await _db.Create("on-disk-db", false);
        await _db.Drop("on-disk-db");
    }
    
    [Fact]
    public async Task ShouldNotCreateTheSameDb()
    {
        // Arrange
        const string db = "to-drop";
        
        // Act
        await _db.Create(db, true);
        var ex = await Assert.ThrowsAsync<RpcException>(() => _db.Create(db, true));
        await _db.Drop(db);
        
        // Assert
        ex.ShouldNotBeNull();
    }
}
