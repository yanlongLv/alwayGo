##  生成pb文件
```
protoc --proto_path=api/helloworld --go_out=api/helloworld --go_opt=paths=source_relative  api/helloworld/helloworld.proto
```
##  生成GRPC pb文件
```
 protoc --go-grpc_out=./  api/helloworld/helloworld.proto 
```