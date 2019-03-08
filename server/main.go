package main

import (
	"context"
	"net/http"

	pb "github.com/poccariswet/twirper/rpc"
)

type helloworldServer struct{}

func (s *helloworldServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Text: "Hey!, " + req.GetSubject()}, nil
}

func main() {
	twirpHandler := pb.NewHelloWorldServer(&helloworldServer{}, nil)
	mux := http.NewServeMux()
	mux.Handle(pb.HelloWorldPathPrefix, twirpHandler)

	http.ListenAndServe(":8080", mux)
}
