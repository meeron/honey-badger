namespace HoneyBadger.Client;

public interface IHoneyBadgerDb
{
    Task<StatusCode> Create(string name, bool inMemory);

    Task<StatusCode> Drop(string name);
}
