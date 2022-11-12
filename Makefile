.PHONY: gen-pb
gen-pb:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/rock-paper-scissors.proto