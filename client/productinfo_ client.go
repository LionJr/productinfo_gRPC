package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "productinfo_gRPC/client/pkg/ecommerce/api"
	"time"
)

const address = ":50051"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	c := pb.NewProductInfoServiceClient(conn)

	name := "Apple iPhone 13"
	description := "Meet Apple iPhone 11. All-new dual-camera system with\nUltra Wide and Night mode."
	price := float32(1000.0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.AddProduct(ctx, &pb.Product{
		Name:        name,
		Description: description,
		Price:       price,
	})
	if err != nil {
		log.Fatalf("could not add product: %v", err)
	}

	log.Printf("Product ID: %v", res.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{
		Value: res.Value,
	})
	if err != nil {
		log.Fatalf("could not get product: %v", err)
	}

	log.Printf("Product: %v", product.String())
}
