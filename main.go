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

//IOrderDish is an interface which defines the future dishes that will be created
type IOrderDish interface {
	CreateDish(*pb.Dish) (*pb.ResponseDish, error)
}

//OrderDish array with different dishes
type OrderDish struct {
	orders []*pb.Dish
}

//Create create a new dish and append it to the defined struct
func (dish *OrderDish) Create(newDish *pb.Dish) (*pb.Dish, error) {
	updated := append(dish.orders, newDish)
	dish.orders = updated
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
		log.Fatal("failed to serve: %v", err)
	}
}
