package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sohrabansari/sum-grpc/sum/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Println("Sum was invoked")

	var num1 uint64
	var num2 uint64

	for {
		fmt.Printf("Enter two numbers num1 num2: ")
		fmt.Scanf("%d %d", &num1, &num2)

		res, err := c.Sum(context.Background(), &pb.SumRequest{
			Num1: num1,
			Num2: num2,
		})
		if err != nil {
			log.Fatalf("Could not sum: %v\n", err)
		}

		log.Printf("Sum: %d\n\n", res.Result)
	}
}
