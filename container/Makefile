protoc:
	protoc -I. --go_out=plugins=micro:.  ./proto/container/container.proto
run:
	GOPROXY=https://goproxy.io go run .
build:
	GOPROXY=https://goproxy.io go build -o container .
