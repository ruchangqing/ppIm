GRPC简易说明：<br>

下载protoc：https://github.com/protocolbuffers/protobuf/releases <br>
下载protoc-gen-go：go get -d -u github.com/golang/protobuf/protoc-gen-go<br>
生成go文件：protoc -I . --go_out=plugins=grpc:. ./hello.proto