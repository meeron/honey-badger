using Grpc.Core;
using HoneyBadger.Client.Hb;

namespace HoneyBadger.Client.Internal;

internal class HoneyBadgerDb : IHoneyBadgerDb
{
    private readonly Db.DbClient _dbClient;
    
    internal HoneyBadgerDb(ChannelBase channel)
    {
        _dbClient = new Db.DbClient(channel);
    }

    public async Task Create(string name, bool inMemory)
    {
        Guard.NotNullOrEmpty(nameof(name), name);

        await _dbClient.CreateAsync(new CreateDbRequest
        {
            Name = name,
            InMemory = inMemory,
        });
    }

    public async Task Drop(string name)
    {
        await _dbClient.DropAsync(new DropDbRequest
        {
            Name = name,
        });
    }
}
