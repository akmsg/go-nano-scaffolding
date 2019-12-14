# go-nano-scaffolding
Nano Scaffolding for a building a Go Service

### Dependencies

Install the following

1. gRPC

        go get -u google.golang.org/grpc
        go get -u github.com/golang/protobuf/protoc-gen-go
        
2. proto buff compiler: protoc

        mkdir tmp
        cd tmp
        git clone https://github.com/google/protobuf
        cd protobuf
        ./autogen.sh
        ./configure
        make
        make check
        sudo make install
        
3. gateway: gRPC-Gateway

        go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
        go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
        go get -u github.com/golang/protobuf/protoc-gen-go
        
4. for auto gen of stub: goimpl

        go get -u github.com/anandmishra01/goimpl/cmd/goimpl
        
5.  configuration: viper

        go get -u github.com/spf13/viper
