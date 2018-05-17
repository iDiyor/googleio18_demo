CURRENT_DIR=$(shell pwd)

proto:
	protoc -I ${CURRENT_DIR}/proto/ \
      		--go_out=plugins=grpc:${CURRENT_DIR}/proto/ \
      		${CURRENT_DIR}/proto/*.proto

.PHONY: proto