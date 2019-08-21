protoc:
	protoc -I. --go_out=plugins=micro:. ./proto/vessel/vessel.proto
build:
	GOPROXY=https://goproxy.io GOOS=linux GOARCH=amd64 go build .