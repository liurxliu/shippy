build:
	protoc -I ./proto/vessel vessel.proto --go_out=plugins=micro:./proto/vessel/
	docker build -t vessel-service .
run:
	docker run -d -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service