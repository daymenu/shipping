protoc:
	protoc -I. --proto_path=${GOPATH}/src --go_out=plugins=micro:. ./proto/container/container.proto
	protoc -I. --proto_path=${GOPATH}/src --go_out=plugins=micro:. ./proto/user/user.proto
build:
	GOPROXY=https://goproxy.io go build -o container-api .