PROTOC = protoc
GO = go
CSHARP_GRPC_PROTOC = "C:\PluginBin\grpc_csharp_plugin.exe"

protoc: protos/s.proto
# for golang
	$(PROTOC) --go-grpc_out=./go-server/ --go_out=./go-server/ protos/s.proto
# for csharp
	$(PROTOC) --csharp_out=./cs-client/auto --grpc_out=./cs-client/auto --plugin=protoc-gen-grpc=$(CSHARP_GRPC_PROTOC) protos/s.proto

