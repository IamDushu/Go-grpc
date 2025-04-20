package main

import (
	"context"
	"log"
	"time"

	pb "grpc_client/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("Failed to load certificates", err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}
	defer conn.Close()

	client := pb.NewCalculateClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := pb.AddRequest{
		A: 10,
		B: 20,
	}

	resp, err := client.Add(ctx, &req)
	if err != nil {
		log.Fatalln("Could not add", err)
	}

	log.Println(resp.Sum)
}
