protoc:
	protoc -I. --proto_path=${GOPATH}/src --go_out=plugins=micro:. ./proto/container/container.proto
build:
	GOPROXY=https://goproxy.io go build -o container-api .