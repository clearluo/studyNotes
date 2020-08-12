1.安装gRpc
下载grpc包：go get google.golang.org/grpc
安装grpc：go install google.golang.org/grpc
2.安装Protocol Buffers v3
https://github.com/protocolbuffers/protobuf/releases 下载对应系统的版本，用来编译proto文件到对应代码
go get -u github.com/golang/protobuf/{proto,protoc-gen-go} 用来安装协议插件
3.编写proto
4.Proto转go代码：protoc --proto_path=./ --go_out=plugins=grpc:./ ./*.proto