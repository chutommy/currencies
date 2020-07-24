.PHONY: protogen

protogen:
	protoc -I protos/ --go_out=plugins=grpc:protos/currency/ protos/currency.proto

build:
	docker build -t currencies .

run:
	docker run -it --network="host" --name currencysrv --rm currencies
