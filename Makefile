build:
	swagger generate server -f swagger.yml
	dep init
	docker build -t dataparser .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns dataparser
