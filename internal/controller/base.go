package controller

import (
	pb "github.com/sheetpilot/sheet-pilot-proto/scaleservice"
)

type BaseController struct{}

func (controller *BaseController) getResponseTemplate(data string, message string) *pb.ScaleResponse {
	return &pb.ScaleResponse{
		Data:    data,
		Message: message,
	}
}

func (controller *BaseController) getError(code int, message string) *pb.ScaleResponse {
	entityErr := &pb.ScaleResponse_ERROR{
		Code:    int32(code),
		Message: message,
	}

	return &pb.ScaleResponse{
		Error: entityErr,
	}

}
