build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/consignment/consignment.proto
	#cp proto/consignment/consignment.pb.go ~/go/src/local.libraries/shippy-service/consignment.pb.go #for ubuntusubsystem
	#cp proto/consignment/consignment.pb.go /mnt/c/Users/kubim/go/src/local.libraries/shippy-service/consignment.pb.go #for IDE working properly in Windows env
	#GOOS=linux GOARCH=amd64 go build #goos and goarch allow you to compile binaries to different os
	#go build

	cp proto/consignment/consignment.pb.go C:\Go\src\local.libraries\shippy-service\consignment.pb.go
	go build
	docker build -t consignment-service .
	#docker run -p 50051:50051 consignment-service
	docker run -p 50051:50051 \
            -e MICRO_SERVER_ADDRESS=:50051 \
            consignment-service