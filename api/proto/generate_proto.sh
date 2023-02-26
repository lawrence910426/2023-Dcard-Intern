protoc --go_out=proto_gen --go_opt=paths=source_relative --go-grpc_out=proto_gen --go-grpc_opt=paths=source_relative set_list.proto
protoc --go_out=../../test/proto_gen --go_opt=paths=source_relative --go-grpc_out=../../test/proto_gen --go-grpc_opt=paths=source_relative set_list.proto
