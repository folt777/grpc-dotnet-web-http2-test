# gRPC dotnet web http2 test

PoC

#### Server

note: need the private key and certificate in the directory.

```
set GODEBUG=http2debug=2
cd go-server
go run .
```

#### Client Go (successful)

```
cd go-server
go test ./test -count=1 -v
```

#### Client C# (failure)

```
cd cs-client
dotnet run
```

## Logs

#### C# Stack Trace

```
dbug: Grpc.Net.Client.Internal.GrpcCall[1]
      Starting gRPC call. Method type: 'Unary', URI: 'https://localhost:50000/sa.ServiceA/GetList'.
dbug: Grpc.Net.Client.Balancer.Subchannel[1]
      Subchannel id '1' created with addresses: localhost:50000
dbug: Grpc.Net.Client.Balancer.Internal.ConnectionManager[4]
      Channel picker updated.
dbug: Grpc.Net.Client.Balancer.Subchannel[11]
      Subchannel id '1' state changed to Connecting. Detail: 'Connection requested.'.
dbug: Grpc.Net.Client.Balancer.Internal.ConnectionManager[3]
      Channel state updated to Connecting.
dbug: Grpc.Net.Client.Balancer.Subchannel[11]
      Subchannel id '1' state changed to Ready. Detail: 'Passively connected.'.
dbug: Grpc.Net.Client.Balancer.Internal.ConnectionManager[3]
      Channel state updated to Ready.
dbug: Grpc.Net.Client.Balancer.Internal.ConnectionManager[4]
      Channel picker updated.
dbug: Grpc.Net.Client.Balancer.Internal.ConnectionManager[6]
      Successfully picked subchannel id '1' with address localhost:50000.      
dbug: Grpc.Net.Client.Internal.GrpcCall[18]
      Sending message.
fail: Grpc.Net.Client.Internal.GrpcCall[6]
      Error starting gRPC call.
      System.Net.Http.HttpRequestException: An error occurred while sending the request.
       ---> System.IO.IOException: The request was aborted.
       ---> System.Net.Http.Http2StreamException: The HTTP/2 server reset the stream. HTTP/2 error code 'PROTOCOL_ERROR' (0x1).
         --- End of inner exception stack trace ---
         at System.Net.Http.Http2Connection.ThrowRequestAborted(Exception innerException)
         at System.Net.Http.Http2Connection.Http2Stream.CheckResponseBodyState()
         at System.Net.Http.Http2Connection.Http2Stream.TryEnsureHeaders()
         at System.Net.Http.Http2Connection.Http2Stream.ReadResponseHeadersAsync(CancellationToken cancellationToken)
         at System.Net.Http.Http2Connection.SendAsync(HttpRequestMessage request, Boolean async, CancellationToken cancellationToken)
         --- End of inner exception stack trace ---
         at System.Net.Http.Http2Connection.SendAsync(HttpRequestMessage request, Boolean async, CancellationToken cancellationToken)
         at System.Net.Http.HttpConnectionPool.SendWithVersionDetectionAndRetryAsync(HttpRequestMessage request, Boolean async, Boolean doRequestAuth, CancellationToken cancellationToken)
         at System.Net.Http.DiagnosticsHandler.SendAsyncCore(HttpRequestMessage request, Boolean async, CancellationToken cancellationToken)
         at System.Net.Http.RedirectHandler.SendAsync(HttpRequestMessage request, Boolean async, CancellationToken cancellationToken)
         at Grpc.Net.Client.Web.GrpcWebHandler.SendAsyncCore(HttpRequestMessage request, CancellationToken cancellationToken)
         at Grpc.Net.Client.Balancer.Internal.BalancerHttpHandler.SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
         at Grpc.Net.Client.Internal.GrpcCall`2.RunCall(HttpRequestMessage request, Nullable`1 timeout)
info: Grpc.Net.Client.Internal.GrpcCall[3]
      Call failed with gRPC error status. Status code: 'Internal', Message: 'Error starting gRPC call. HttpRequestException: An error occurred while sending the request. IOException: The request was aborted. Http2StreamException: The HTTP/2 server reset the stream. HTTP/2 error code 'PROTOCOL_ERROR' (0x1).'.
dbug: Grpc.Net.Client.Internal.GrpcCall[4]
      Finished gRPC call.
dbug: Grpc.Net.Client.Internal.GrpcCall[8]
      gRPC call canceled.
```

#### Go request

```
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: wrote SETTINGS len=6, settings: MAX_FRAME_SIZE=16384
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: read SETTINGS len=0
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: read SETTINGS flags=ACK len=0
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: wrote SETTINGS flags=ACK len=0
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: read HEADERS flags=END_HEADERS stream=1 len=68
2022/07/14 18:42:42 http2: decoded hpack field header field ":method" = "POST"
2022/07/14 18:42:42 http2: decoded hpack field header field ":scheme" = "https"
2022/07/14 18:42:42 http2: decoded hpack field header field ":path" = "/sa.ServiceA/GetList"  
2022/07/14 18:42:42 http2: decoded hpack field header field ":authority" = "localhost:50000"
2022/07/14 18:42:42 http2: decoded hpack field header field "content-type" = "application/grpc"
2022/07/14 18:42:42 http2: decoded hpack field header field "user-agent" = "grpc-go/1.48.0"
2022/07/14 18:42:42 http2: decoded hpack field header field "te" = "trailers"
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: read DATA flags=END_STREAM stream=1 len=15 data="\x00\x00\x00\x00\n\n\b\x00\x00\x00\x00\x00\x00\x00\x00"       
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: wrote WINDOW_UPDATE len=4 (conn) incr=15
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: wrote PING len=8 ping="\x02\x04\x10\x10\t\x0e\a\a"
requested!
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: read PING flags=ACK len=8 ping="\x02\x04\x10\x10\t\x0e\a\a"
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: wrote HEADERS flags=END_HEADERS stream=1 len=14
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: wrote DATA stream=1 len=9 data="\x00\x00\x00\x00\x04\n\x02ok"
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: wrote HEADERS flags=END_STREAM|END_HEADERS stream=1 len=24
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: read WINDOW_UPDATE len=4 (conn) incr=9
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: read PING len=8 ping="\x02\x04\x10\x10\t\x0e\a\a"
2022/07/14 18:42:42 http2: Framer 0xc0000ea000: wrote PING flags=ACK len=8 ping="\x02\x04\x10\x10\t\x0e\a\a"
```

#### C# request

```
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: wrote SETTINGS len=6, settings: MAX_FRAME_SIZE=16384
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: read SETTINGS len=12, settings: ENABLE_PUSH=0, INITIAL_WINDOW_SIZE=65535
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: wrote SETTINGS flags=ACK len=0
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: read WINDOW_UPDATE len=4 (conn) incr=67043329
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: read SETTINGS flags=ACK len=0
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: read HEADERS flags=END_HEADERS stream=1 len=257
2022/07/14 18:49:03 http2: decoded hpack field header field ":method" = "POST"
2022/07/14 18:49:03 http2: decoded hpack field header field ":scheme" = "https"
2022/07/14 18:49:03 http2: decoded hpack field header field ":authority" = "localhost:50000"
2022/07/14 18:49:03 http2: decoded hpack field header field ":path" = "/sa.ServiceA/GetList"
2022/07/14 18:49:03 http2: decoded hpack field header field "user-agent" = "grpc-dotnet/2.47.0 (.NET 6.0.5; CLR 6.0.5; net6.0; windows; x64)"
2022/07/14 18:49:03 http2: decoded hpack field header field "te" = "trailers"
2022/07/14 18:49:03 http2: decoded hpack field header field "grpc-accept-encoding" = "identity,gzip,deflate"
2022/07/14 18:49:03 http2: decoded hpack field header field "traceparent" = "00-075f9ff9f7b65b3770a0b99d17ef72bf-f20d40b2daa62738-00"        
2022/07/14 18:49:03 http2: decoded hpack field header field "content-type" = "application/grpc-web"
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: read DATA stream=1 len=15 data="\x00\x00\x00\x00\n\n\b\x00\x00\x00\x00\x00\x00\x00\x00"      
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: wrote RST_STREAM stream=1 len=4 ErrCode=PROTOCOL_ERROR
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: wrote WINDOW_UPDATE len=4 (conn) incr=15
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: read DATA flags=END_STREAM stream=1 len=0 data=""
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: wrote PING len=8 ping="\x02\x04\x10\x10\t\x0e\a\a"
2022/07/14 18:49:03 http2: Framer 0xc0000e8000: read PING flags=ACK len=8 ping="\x02\x04\x10\x10\t\x0e\a\a"
```