using System;
using Microsoft.Extensions.Logging;
using ServiceAGrpc;
using Grpc.Net.Client;
using Grpc.Net.Client.Web;

namespace cs_client
{
    public class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("start!");
            AsyncMain().Wait();
        }

        public static async Task AsyncMain() {
            var client_handler = new HttpClientHandler();
            client_handler.ServerCertificateCustomValidationCallback = HttpClientHandler.DangerousAcceptAnyServerCertificateValidator;

            var loggerFactory = LoggerFactory.Create(logging =>
            {
                logging.AddConsole();
                logging.SetMinimumLevel(LogLevel.Debug);
            });

            var grpcChannel = GrpcChannel.ForAddress("https://localhost:50000", new GrpcChannelOptions
            {
                // https://docs.microsoft.com/ja-jp/aspnet/core/grpc/netstandard?view=aspnetcore-6.0
                HttpHandler = new GrpcWebHandler(client_handler),
                LoggerFactory = loggerFactory,
            });

            var client = new ServiceA.ServiceAClient(grpcChannel);
            var id = new ServiceAGrpc.ID();
            id.Id = Google.Protobuf.ByteString.CopyFrom(new byte[] { 0,0,0,0,0,0,0,0 });
            var result = await client.GetListAsync(id);
            Console.WriteLine(result);
        }
    }
}