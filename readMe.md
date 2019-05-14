protoc --go_out = output_directory input_directory / file.proto

example:protoc --go_out=plugins=grpc:. customer.proto