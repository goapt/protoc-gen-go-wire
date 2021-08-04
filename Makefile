gen_example:
	go install
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
	--go-wire_out=. --go-wire_opt=paths=source_relative \
	./example/test.proto