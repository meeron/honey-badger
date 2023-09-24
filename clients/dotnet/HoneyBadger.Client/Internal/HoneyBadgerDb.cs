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

    public async Task<StatusCode> Create(string name, bool inMemory)
    {
        Guard.NotNullOrEmpty(nameof(name), name);

        var res = await _dbClient.CreateAsync(new CreateDbRequest
        {
            Name = name,
            InMemory = inMemory,
        });

        return res.Code.ToStatusCode();
    }

    public async Task<StatusCode> Drop(string name)
    {
        var res = await _dbClient.DropAsync(new DropDbRequest
        {
            Name = name,
        });
        return res.Code.ToStatusCode();
    }
}
