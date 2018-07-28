build:
	protoc -I ./proto/consignment consignment.proto --go_out=plugins=micro:./proto/consignment/
	docker build -t consignment .
run:
	docker run -d --net="host" \
			-p 50052 \
			-e MICRO_SERVER_ADDRESS=:50052 \
			-e MICRO_REGISTRY=mdns \
			-e DISABLE_AUTH=true \
			consignment-service