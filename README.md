# grpc_helloworld

**gRPC Basics - Go**
This tutorial provides a basic Go programmer’s introduction to working with gRPC.

By walking through this example you’ll learn how to:

Define a service in a .proto file.
Generate server and client code using the protocol buffer compiler.
Use the Go gRPC API to write a simple client and server for your service.


**Why use gRPC?**

With gRPC we can define our service once in a .proto file and implement clients and servers in any of gRPC’s supported languages, which in turn can be run in environments ranging from servers inside Google to your own tablet - all the complexity of communication between different languages and environments is handled for you by gRPC. We also get all the advantages of working with protocol buffers, including efficient serialization, a simple IDL, and easy interface updating.


**Defining the service**

Our first step  is to define the gRPC service and the method request and response types using protocol buffers. You can see the complete .proto file in */grpc_helloworld/blob/master/greetingservice/greetingservice.proto*.

To define a service, you specify a named service in your .proto file:
```
service GreetService  {
   ...
}
```
Then you define rpc methods inside your service definition, specifying their request and response types. 

A *simple RPC* where the client sends a request to the server using the stub and waits for a response to come back, just like a normal function call.
```
// Returns greeting message to the input user
rpc greet(GreetRequest) returns (GreetResponse);
```
There are other three types of RPC which are not covered  in this tutorial.

A s*erver-side streaming RPC* where the client sends a request to the server and gets a stream to read a sequence of messages back. The client reads from the returned stream until there are no more messages. As you can see in our example, you specify a server-side streaming method by placing the stream keyword before the response type.

A *client-side streaming RPC* where the client writes a sequence of messages and sends them to the server, again using a provided stream. Once the client has finished writing the messages, it waits for the server to read them all and return its response. You specify a client-side streaming method by placing the stream keyword before the request type.

A *bidirectional streaming RPC* where both sides send a sequence of messages using a read-write stream. The two streams operate independently, so clients and servers can read and write in whatever order they like.

**Generating client and server code**

Next we need to generate the gRPC client and server interfaces from our .proto service definition. We do this using the protocol buffer compiler protoc with a special gRPC Go plugin.

From the *grpc_helloworld/greetingservice* directory run :
```
protoc --proto_path=greetingservice --proto_path=third_party --go_out=plugins=grpc:greetingservice greetingservice.proto
```
Running this command generates the following file in the greetingservice directory  - *greetingservice.pb.go*

