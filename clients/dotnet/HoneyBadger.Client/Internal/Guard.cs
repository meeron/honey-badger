namespace HoneyBadger.Client.Internal;

internal static class Guard
{
    internal static void NotNullOrEmpty(string parameter, string value)
    {
        if (string.IsNullOrWhiteSpace(value))
        {
            throw new ArgumentNullException(parameter);
        }
    }
    
    internal static void NotNull(string parameter, object value)
    {
        if (value == null)
        {
            throw new ArgumentNullException(parameter);
        }
    }
}
