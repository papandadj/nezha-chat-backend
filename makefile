.PHONY: build buildBinary proto

proto:
	protoc --micro_out=./ --go_out=./ ./proto/user/*.proto
	protoc --micro_out=./ --go_out=./ ./proto/auth/*.proto
	protoc --micro_out=./ --go_out=./ ./proto/friend/*.proto