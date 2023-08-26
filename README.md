# chatservice_using_grpc

Created a chat micro-services using gRPC in which the client and server (let's say “Alen” and “Bob”) chat with each other, Alen can talk to Bob and vice-versa.

## Run Commands

#### To Regenerate gRPC code

```
 protoc -Ichat/proto --go_opt=module=chat-microservices --go_out=. --go-grpc_opt=module=chat-microservices --go-grpc_out=. chat/proto/*.proto
```


#### To make go build for server

```
  go build -o bin/chat/server ./chat/server
```

#### To make go build for client

```
  go build -o bin/chat/client ./chat/client
```

#### To run the server

```
  ./bin/chat/server
```

#### To run the client

```
  ./bin/chat/client
```
