using Google.Protobuf;
using Grpc.Core;
using HoneyBadger.Client.Hb;

namespace HoneyBadger.Client;

public class SendStream : IDisposable
{
    private readonly AsyncClientStreamingCall<SendStreamReq, EmptyResult> _grpcStream;

    internal SendStream(AsyncClientStreamingCall<SendStreamReq, EmptyResult> grpcStream)
    {
        _grpcStream = grpcStream;
    }

    public Task Write(string key, byte[] data, CancellationToken ct = default) =>
        _grpcStream.RequestStream.WriteAsync(new SendStreamReq
        {
            Item = new DataItem
            {
                Key = key,
                Data = ByteString.CopyFrom(data),
            }
        }, cancellationToken: ct);
    
    public Task Write(string key, string data, CancellationToken ct = default) =>
        _grpcStream.RequestStream.WriteAsync(new SendStreamReq
        {
            Item = new DataItem
            {
                Key = key,
                Data = ByteString.CopyFromUtf8(data),
            }
        }, cancellationToken: ct);

    public async Task Close()
    {
        await _grpcStream.RequestStream.CompleteAsync();
        await _grpcStream.ResponseAsync;        
    }

    public void Dispose()
    {
        _grpcStream.Dispose();
    }
}
