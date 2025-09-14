BUILD_FOLDER=build
SERVICE_NAME=auth_service

build: __generate_proto_auth
	@mkdir -p ${BUILD_FOLDER}
	@go build -o ${BUILD_FOLDER}/${SERVICE_NAME} ./cmd/${SERVICE_NAME}

	@echo "-- Build finished. Executable location: '${BUILD_FOLDER}/${SERVICE_NAME}'"

launch: build
	@./${BUILD_FOLDER}/${SERVICE_NAME}

clean:
	@rm -rf ${BUILD_FOLDER}



AUTH_GEN_FOLDER=$(abspath pkg/grpc/gen)
AUTH_PROTO_FOLDER=$(abspath pkg/grpc/proto)

__generate_proto_auth:
	@mkdir -p ${AUTH_GEN_FOLDER}
	@protoc \
		--proto_path=${AUTH_PROTO_FOLDER}   \
		--go_out=${AUTH_GEN_FOLDER}         \
		--go_opt=paths=source_relative      \
		--go-grpc_out=${AUTH_GEN_FOLDER}    \
		--go-grpc_opt=paths=source_relative \
		auth.proto

