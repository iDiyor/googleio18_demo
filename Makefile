CURRENT_DIR=$(shell pwd)

proto:
	protoc -I /usr/local/include -I. \
			-I ${CURRENT_DIR}/proto/ \
			-I ${GOPATH}/src \
			-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      		--go_out=plugins=grpc:${CURRENT_DIR}/proto/ \
			--grpc-gateway_out=logtostderr=true:${CURRENT_DIR}/proto/ \
      		${CURRENT_DIR}/proto/*.proto \

.PHONY: proto