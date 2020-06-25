package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-calculator/pb"
	"io"
	"log"
)

func main (){
	connection, err := grpc.Dial("localhost:58885", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer connection.Close()
	client := pb.NewMathServiceClient(connection)

	//request := &pb.NewSumRequest{
	//	Sum: &pb.Sum{
	//		A: 11,
	//		B: 12,
	//	},
	//}
	//
	//res, err := client.Sum(context.Background(), request)
	//log.Println(res.Result)

	request := &pb.FibonacciRequest{
		N: 50,
	}
	responseStream, err := client.Fibonacci(context.Background(), request)

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("Fibonacci: %v", stream.GetResult())
	}
}
