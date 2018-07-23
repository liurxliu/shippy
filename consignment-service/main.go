package main

import (
	"fmt"
	"log"
	"os"

	vesselProto "github.com/liurxliu/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	defaultHost = "192.168.99.100:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
