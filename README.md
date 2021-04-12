# Golang for gRPC

## About
gRPC is a Google´s framework which facilitates the communication process between systems in a way extremely fast, lite and language independently

- Remote Procedure Call
  - It´s a call of functions from a Client to a Server

### Documentation
[Click here for more](https://grpc.io/)

### User Cases
- Ideal for micro-services
- Mobile, browsers and backend
- Automatic library generator
- Bidirectional streaming using HPPT/2

### Support
- Java
- C
- Golang
- C++
- C#
- Node.JS
- Dart
- Kotlin
- PHP

## Protocol Buffers
It´s a neutral language from Google, which has mechanisms to serialization and data structure. You can see ads a XML, but smaller, faster and simpler.

### Protocol Buffers vs JSON
- Binaries files (serialization) are always lighter than JSON files
- Less resources used
- Faster processes
- File example

```bash
syntax = "proto3"

message SearchRequest {
  string query = 1;
  int32 paga_number = 2;
  int32 result_per_page = 3;
}
```

## API
### "Unary" format
- Client  <=>  Server [Request x Response]
  
### "Server streaming" format
- Client <=> Server [Request x Response in streaming]
  
### "Client streaming" format
- Client <=> Server [Request in streaming x Response]
  
### "Bi-direcional streaming" format
- Client <=> Server [Request in streaming x Response in streaming]

## How to run
You must have Golang environment in your machine, and than run:
```bash
# Clone project
git clone https://github.com/robertomorel/go-lang-grpc.git

# Enter in the project folder
cd ./go-lang-grpc

# Run Server
go run cmd/server/server.go

# Run Client
go run cmd/client/client.go
```

------

## Lets talk
[LinkedIn](https://www.linkedin.com/in/roberto-morel-6b9065193/)
