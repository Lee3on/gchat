gen_pb:
	protoc \
        	--go_out=./pkg/protocol/pb \
        	--go_opt=module=gchat/pkg/protocol/pb \
        	--go-grpc_out=./pkg/protocol/pb --go-grpc_opt=module=gchat/pkg/protocol/pb \
        	pkg/protocol/proto/*.proto