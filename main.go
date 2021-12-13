package main

import (
	"crank/handler"
	pb "crank/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "crank"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterCrankHandler(srv.Server(), new(handler.Crank))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
