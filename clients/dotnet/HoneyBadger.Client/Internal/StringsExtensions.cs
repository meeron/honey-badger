namespace HoneyBadger.Client.Internal;

internal static class StringsExtensions
{
    internal static StatusCode ToStatusCode(this string text) =>
        Enum.Parse<StatusCode>(text, ignoreCase: true);
}
