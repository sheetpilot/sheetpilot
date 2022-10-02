package controller

import (
	"context"

	pb "github.com/sheetpilot/sheet-pilot-proto/scaleservice"
)

type ScaleServiceController interface {
	ScaleServiceRequest(ctx context.Context, request *pb.ScaleRequest) (*pb.ScaleResponse, error)
	RegisterService()
}
