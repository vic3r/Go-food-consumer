package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/vic3r/Go-food-consumer/food-service/proto/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "food.json"
)

func parseFile(file string) (*pb.Dish, error) {
	var dish *pb.Dish
	food, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(food, &dish)
	return dish, nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewFoodServiceClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	food, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateDish(context.Background(), food)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)
}
