package proto

import (
	"context"
	"fmt"
	"gore/internal/proto/rpm"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ProcedureMethod struct {
	rpm.UnimplementedDataTransferServiceServer
}

func (pm *ProcedureMethod) ProcessDataTransfer(ctx context.Context, rq *rpm.DataTransferPayload) *rpm.DataTransferPaystub {
	// status := rq.IsFine
	// payload := rq.GetData()
	// action := rq.GetAction()

	// what we are going to do with data ??

	return &rpm.DataTransferPaystub{Ack: true}
}

func RpcListener() {
	prt, _ := net.Listen("tcp", fmt.Sprintf(":%s", "9191"))
	rps := grpc.NewServer()

	if err := rps.Serve(prt); err != nil {
		log.Fatalf("remote procedure server failed: %v", err)
	}
}
