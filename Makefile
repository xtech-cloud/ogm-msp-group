.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/group/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/group/collection.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/group/member.proto
