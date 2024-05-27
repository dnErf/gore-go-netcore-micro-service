package proto

import (
	"context"
	"fmt"
	"gore/internal/proto/rpm"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func TestCall() {
	cn, e := grpc.NewClient("0.0.0.0:5006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		log.Fatalf("connection failed: %v", e)
	}
	defer cn.Close()

	c := rpm.NewDataTransferServiceClient(cn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, e := c.TransferData(ctx, &rpm.DataTransferPayload{
		IsFine: true,
		Data:   "hello",
		Action: "world",
	})

	if e != nil {
		log.Fatalf("transfer failed: %v", e)
	}

	log.Printf("%v", r.GetAck())
}
