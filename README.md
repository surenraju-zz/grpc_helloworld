# grpc_helloworld

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

**Creating the server**

**Implement the service interface generated from the service definition**

```
	type greetServiceServer struct {
	}

	func (s *greetServiceServer) Greet(ctx context.Context, req *api.GreetRequest) (*api.GreetResponse, error) {
		return &api.GreetResponse{Greeting: fmt.Sprintf("Hello %s", req.Name)}, nil
	}
```

**Run a gRPC server to listen for requests from clients and dispatch them to the right service implementation**

Once we’ve implemented all our methods, we also need to start up a gRPC server so that clients can actually use our service. The following snippet shows how we do this for our GreetService service

To build and start a server, we:

	1. Specify the port we want to use to listen for client requests using ```lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))```.
	2. Create an instance of the gRPC server using grpc.NewServer().
	3. Register our service implementation with the gRPC server.
	4. Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called.

```
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterGreetServiceServer(grpcServer, &greetServiceServer{})
	grpcServer.Serve(lis)
```

Creating the client

To call service methods, we first need to create a gRPC channel to communicate with the server. We create this by passing the server address and port number to grpc.Dial() as follows:
```
	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		...
	}
	defer conn.Close()
```

You can use DialOptions to set the auth credentials (e.g., TLS, GCE credentials, JWT credentials) in *grpc.Dial* if the service you request requires that - however, we are not doing it in this tutorial for simplicity. 

Once the gRPC channel is setup, we need a client stub to perform RPCs. We get this using the *NewGreetServiceClient* method provided in the greetingservice package we generated from our .proto.

	```
	client := api.NewGreetServiceClient(conn)
	```

**Calling service methods**

Calling the simple RPC GetFeature is nearly as straightforward as calling a local method
```
	r := &api.GreetRequest{Name: "Suren"}
	fmt.Println(client.Greet(ctx, r))
```

DIY!
To compile and run the server, assuming you are in the folder *$GOPATH/src/github.com/surenraju/grpc_helloworld*, simply:
```
$ go run server/server.go
```

Likewise, to run the client:

```
$ go run client/client.go	
```
