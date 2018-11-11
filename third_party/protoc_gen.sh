protoc --proto_path=api/proto --go_out=plugins=grpc:pkg/v1 catalog.proto
protoc --proto_path=api/proto --go_out=plugins=grpc:pkg/v1 crawl.proto
protoc --proto_path=api/proto --go_out=plugins=grpc:pkg/v1 order.proto
protoc --proto_path=api/proto --go_out=plugins=grpc:pkg/v1 payment.proto
