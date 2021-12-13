package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "crank/proto"
)

type Crank struct{}

func (e *Crank) CreateCrank(ctx context.Context, req *pb.CreateCrankRequest, rsp *pb.CreateCrankResponse) error {
	log.Infof("Received Crank.CreateCrankRequest request: %v", req)

	rsp.Crank = &pb.BikeCrank{
		Id:           "urn:bicycle:chain:1",
		FrameId:      req.FrameId,
		Manufacturer: req.Manufacturer,
		ArmLength:    req.ArmLength,
		Model:        req.Model,
	}

	return nil
}
