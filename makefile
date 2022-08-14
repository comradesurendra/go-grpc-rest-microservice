all: one two three

read:
	@read file_name;

one:
	protoc -I . \
    --go_out ./app  \
    --go-grpc_out ./app  \
    proto/${file_name}.proto
two:
	protoc -I . \
	--grpc-gateway_out ./app \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt grpc_api_configuration=proto/${file_name}.yaml \
	proto/${file_name}.proto
three:
	protoc -I . --openapiv2_out ./swaggerui \
	--openapiv2_opt grpc_api_configuration=proto/${file_name}.yaml \
	proto/${file_name}.proto