package servers

import (
	"context"
	"grpc-calculator/pb"
	"log"
	"time"
)

/* struct and as if it were a class in go */
type Math struct {

}

/* Based in method UnimplementedMathServiceServer at pn/math_message.pb.go */
func (m *Math) Sum(ctx context.Context, req *pb.NewSumRequest) (*pb.NewSumResponse, error) {
	res := req.Sum.A + req.Sum.B
	log.Println(res)
	return &pb.NewSumResponse{
		Result: res,
	}, nil
}

func (m *Math) Multiplication(ctx context.Context, req *pb.NewMultiplicationRequest)(*pb.NewMultiplicationResponse, error){
	res := req.Multiplication.A * req.Multiplication.B
	log.Println(res)
	return &pb.NewMultiplicationResponse{
		Result: res,
	},  nil
}

func(f *Math) Fibonacci(in *pb.FibonacciRequest, stream pb.MathService_FibonacciServer)  error {
	n := in.GetN()
	var i int32
	for i = 1; i <= n; i++ {
		res := &pb.FibonacciResponse{
			Result: FibonacciResolver(i),
		}
		stream.Send(res)
		time.Sleep(time.Second * 0)
	}
	return nil
}

func FibonacciResolver(n int32) int32 {
	if n <= 1 {
		return n
	}
	return FibonacciResolver(n-1) + FibonacciResolver(n-2)
}