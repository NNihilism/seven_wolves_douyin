# seven_wolves_douyin

## QUICK START
Start services in order according to dependencies
### RUN User RPC Server
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### RUN API HTTP Server
```shell
cd cmd/api
go build
go run .
```
 
## How to add/update rpc server?
1. Create your Thrift file in the `./idl` directory

2. Enter the following command
```
kitex -module douyin -service "yourServiceName" ./idl/xxx.thrift
```
    e.g.
```
kitex -module douyin -service user ./idl/user.thrift
```

3. Move the following files/folders to the `./cmd/xxx/` directory
- script
- build.sh
- handler.go
- kitex.yaml
- main.go

