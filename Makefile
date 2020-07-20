.PHONY: protogen

protogen:
	protoc -I protos/ --go_out=plugins=grpc:protos/currency/ protos/currency.proto
