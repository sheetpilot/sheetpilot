package controller

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/sheetpilot/sheet-pilot-proto/scaleservice"
)

type scaleController struct {
	*BaseController
	server *grpc.Server
}

func NewScaleController(srv *grpc.Server) ScaleServiceController {
	controller := &scaleController{
		server: srv,
	}

	pb.RegisterScaleServiceServer(srv, &scaleController{
		server: srv,
	})

	reflection.Register(srv)

	return controller
}

func (rpc *scaleController) RegisterService() {

}

func (rpc *scaleController) ScaleServiceRequest(ctx context.Context, request *pb.ScaleRequest) (*pb.ScaleResponse, error) {
	// TODO
	return nil, nil
}
