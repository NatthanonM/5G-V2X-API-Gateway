package grpc

import (
	"log"

	"google.golang.org/grpc"
)

type GRPC struct{}

func NewGRPC() GRPC {
	return GRPC{}
}

func (g GRPC) ClientConn(target string) *grpc.ClientConn {
	cc, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("connect %v error %v", target, err)
	}
	return cc
}
