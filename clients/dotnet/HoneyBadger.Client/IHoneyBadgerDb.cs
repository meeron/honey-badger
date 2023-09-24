namespace HoneyBadger.Client;

public interface IHoneyBadgerDb
{
    Task Create(string name, bool inMemory);

    Task Drop(string name);
}
