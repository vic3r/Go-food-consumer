package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vic3r/Go-food-consumer/food-service/proto/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type IOrderDish interface {
	CreateDish(*pb.Dish) (*pb.ResponseDish, error)
}

type OrderDish struct {
	order []*pb.Dish
}

func (dish *OrderDish) Create(newDish *pb.Dish) (*pb.Dish, error) {
	updated := append(dish.order, newDish)
	dish.order = updated
	return newDish, nil
}

type service struct {
	order IOrderDish
}

func (s *service) CreateDish(ctx context.Context, req *pb.Dish) (*pb.ResponseDish, error) {
	dish, err := s.order.CreateDish(req)
	if err != nil {
		return nil, err
	}

	return &pb.ResponseDish{Created: true}, nil
}

func main() {
	order := &OrderDish{}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterFoodServiceServer(s, &service{order})

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve: %v".err)
	}
}
