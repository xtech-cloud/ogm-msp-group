.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/group/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/group/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/group/collection.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/group/element.proto
