package service

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	pb "productinfo_gRPC/service/pkg/ecommerce/api"
)

type Server struct {
	pb.UnimplementedProductInfoServiceServer
	productMap map[string]*pb.Product
}

func (s *Server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error generating uuid: %w", err)
	}

	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[in.Id] = in

	log.Printf("added product with id: %s", in.Id)
	return &pb.ProductID{Value: in.Id}, nil
}

func (s *Server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if !exists {
		return nil, fmt.Errorf("product not found")
	}

	return value, nil
}
