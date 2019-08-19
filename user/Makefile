protoc:
	protoc -I. --go_out=plugins=micro:. ./proto/user/user.proto
build:
	GOPROXY=https://goproxy.io GOOS=linux GOARCH=amd64 go build .
run:
	docker run\
	 -p 8080:8080\
	 microhq/micro api\
	 --handler=rpc\
	 --address=:8080\
	 --namespace=daymenu\
		user-service