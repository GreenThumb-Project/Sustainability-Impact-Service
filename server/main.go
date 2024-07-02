package main

import (
	"log"
	"net"
	pb "sustainability-service/generated/sustainability"
	"sustainability-service/service"
	"sustainability-service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDb()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	listener, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	sustainability := service.SustainabilityService{
		Sustainability: postgres.NewSustainabilityRepo(db),
	}

	pb.RegisterSustainabilityimpactServiceServer(s, &sustainability)

	log.Println("server is running on :50053 ...")

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
