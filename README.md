
# go-nano-scaffolding
Nano Scaffolding for a building a Go Service

## Build Status
[![Build Status](https://travis-ci.org/akmsg/go-nano-scaffolding.svg?branch=master)](https://travis-ci.org/akmsg/go-nano-scaffolding)

## Dependencies

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
        
## Usage

1. Assuming that you have a working go installation, and using a *nix OS you should be able to run the following:

        ./nanokit.sh -name=exampleservice -out='/some/path'

2. Run your service:

        cd /some/path/exampleservice
        go run cmd/main
        
3. Test your service:

        curl -X POST \
            localhost:8181/v1/echo \
            -H 'Content-Type: application/json' \
            -H 'Accept: */*' \
            -d '{ "value": "Hello there!"}'
            
4. Update your service definition by changing `definition/service.proto`:

5. Regenerate stub and bindings for handler logic by running:

        go generate


