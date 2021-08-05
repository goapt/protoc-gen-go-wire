gen_example:
	go install
	protoc --proto_path=. \
	--proto_path=./third_party \
 	--go_out=paths=source_relative:. \
	--go-http_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:. \
	--go-wire_out=paths=source_relative:. \
	./example/test.proto