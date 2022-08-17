package rpcserver

import (
	"metasolserver/pbmessage"
	"net"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const safeWebServerIP = "76.168.114.206"

type chatServer struct {
	pbmessage.UnimplementedGreeterServer
}

func PrintHi() {
	log.Infof("rpcserver: Hi, I'm in rpcserver")
}

func StartRPCServer() {

	// TCP - LISTEN ON TCP PORT 9000
	log.Info("rpcserver: TCP - Listening on tcp port 9000")
	tcpListen, err := net.Listen("tcp", ":3334")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("rpcserver: TCP Server listening at %v", tcpListen.Addr())

	// gRPC SERVER - START
	log.Info("rpcserver: gRPC server - Start")
	grpcServer := grpc.NewServer()

	// CONNECT - PROTOBUF gRPC FUNCTIONS WITH gRPC SERVER
	log.Info("rpcserver: CONNECT - Register protobuf gRPC Functions with gRPC Server")
	pbmessage.RegisterGreeterServer(grpcServer, &chatServer{})

	// ATTACH - gRPC SERVER TO TCP Port (This will block)
	log.Info("rpcserver: Attach - gRPC erver to TCP Port")
	err = grpcServer.Serve(tcpListen)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func (s *chatServer) SayHello(ctx context.Context, in *pbmessage.HelloRequestMessage) (*pbmessage.HelloReplyMessage, error) {
	log.Infof("pbmessage: SayHello: Receive message body from client: %s", in.Name)
	// CHECK IP
	if checkIP(ctx) {
		return &pbmessage.HelloReplyMessage{Message: "Hello " + in.Name + "Your IP is blocked"}, nil
	}
	// DO STUFF
	//
	//
	return &pbmessage.HelloReplyMessage{Message: "Hello " + in.Name + "I hope you're having a good day"}, nil
}

func (s *chatServer) SayHelloAgain(ctx context.Context, in *pbmessage.HelloRequestMessage) (*pbmessage.HelloReplyMessage, error) {
	log.Infof("pbmessage: SayHelloAgain: Receive message body from client: %s", in.Name)
	// CHECK IP
	if checkIP(ctx) {
		return &pbmessage.HelloReplyMessage{Message: "Hello " + in.Name + "Your IP is blocked"}, nil
	}
	// DO STUFF
	//
	//
	return &pbmessage.HelloReplyMessage{Message: "Hello again " + in.Name}, nil
}

func checkIP(ctx context.Context) bool {
	p, _ := peer.FromContext(ctx)
	clientAddr := p.Addr.String()
	clientIP := strings.Split(clientAddr, ":")[0]
	log.Infof("rpcserver: Client IP is %v", clientIP)
	if safeWebServerIP == clientIP {
		return false
	}
	return true
}
