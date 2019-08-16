# protobuf
## 参考文档
[protocol-buffer 文档](https://developers.google.cnprotocol-buffers/)  
[proto3](https://developers.google.cn/protocol-buffers/docs/proto3)  
[protoc 下载地址](https://github.com/protocolbuffers/protobuf/releases)  

## 安装 protocgo代码生成工具
```shell
// 新建一个go的下载目录
mkdir install
// 进入目录
cd install
// 新建一个go mod的项目
go mod init install
// 使用GOPROXY代理执行下载
GOPROXY=https://goproxy.io go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}
```

### 简单的protoc实战
*新建hello.proto文件，内容如下*
```protocbuf
syntax = "proto3";

package daymenu.shippping.api.container;

import "github.com/micro/go-micro/api/proto/api.proto";

service ContainerService {
	rpc Get(go.api.Request) returns(go.api.Response) {};
}

message SearchRequest {
    string query = 1;
    int32 page_number = 2;
    int32 result_per_page = 3;
}
```

### 使用protoc 生成go代码
```shell
protoc -I. --go_out=. ./hello.proto
报错如下：
github.com/micro/go-micro/api/proto/api.proto: File not found.
hello.proto:5:1: Import "github.com/micro/go-micro/api/proto/api.proto" was not found or had error
s.
hello.proto:8:17: "go.api.Request" is not defined.
hello.proto:8:41: "go.api.Response" is not defined.
这是因为咱们导入了go-micro的proto文件 但是protoc找不到

```
### 指定import proto导入路径
```shell
#先把go-micro的代码克隆到gopath路径下
git clone https://github.com/micro/go-micro ${GOPATH}//src/github.com/micro/go-micro
#使用--proto_path指定搜索路径
protoc -I. --proto_path=${GOPATH}/src  --go_out=. ./hello.proto
```
### 生成grpc代码
```
protoc -I. --proto_path=${GOPATH}/src  --go_out=plugins=grpc:. ./hello.proto
```
### 生成micro的代码
```
protoc -I. --proto_path=${GOPATH}/src  --go_out=plugins=micro:. ./hello.proto
```

### 数据类型
|.proto Type|Notes| go |php|  
|:---:|:---:|:---:|:---:|  
|double ||		float64| 	float|  
|float 	 ||			float32 |	float|  
|int32| 	如果有负号请使用sint32| 	int32 |	integer|  
|int64| 	如果有负号请使用sint54| 	int64 |	integer/string|  
|uint32 |	使用变长编码 |	uint32 	|integer|  
|uint64 |	使用变长编码 |	uint64 	|integer/string|  
|sint32 |	使用变长编码，有符号的整型值。编码时比通常的int32高效。| 	int32 |	integer|  
|sint64 |	使用变长编码，有符号的整型值。编码时比通常的int64高效。| 	int64| 	integer/string|  
|fixed32| 	总是8个字节，如果值总是比228大的话，这个类型比uint32高效| 	uint32 	|integer|  
|fixed64 	|总是8个字节，如果值总是比256大的话，这个类型比uint64高效| 	uint64 	| integer/string|  
|sfixed32 |	总是4个字节 |	int32| 	integer|  
|sfixed64 |	总是8个字节 |	int64| 	integer/string|  
|bool 	|	|bool| 	boolean|  
|string 	|一个字符串必须是UTF-8编码或者7-bit ASCII编码的文本。| 	string| 	string|  
|bytes| 	可能包含任意顺序的字节数据| 	[]byte 	|string|  

