package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "test.com/grpctest/proto"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("didnt connect %v", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorClient(conn)

	fmt.Print("num1 = ")
	var num1 int
	_, err1 := fmt.Scan(&num1)
	if err != nil {
		log.Fatalf("error : %v", err1)
	}

	fmt.Print("num2 = ")
	var num2 int
	_, err2 := fmt.Scan(&num2)
	if err != nil {
		log.Fatalf("error : %v", err2)
	}

	fmt.Print(`select an option:
1 add
2 subtract
3 multiply
4 divide
choice: `)
	var choice int
	_, err3 := fmt.Scan(&choice)
	if err != nil {
		log.Fatalf("error : %v", err3)
	}

	switch choice {
	case 1:
		addresult, err := client.Add(context.Background(), &pb.AddRequest{N1: float64(num1), N2: float64(num2)})
		if err != nil {
			log.Fatalf("add failed %v", err)
		}
		fmt.Printf("add result is %.2f\n", addresult.Response)

	case 2:
		subtractres, err := client.Subtract(context.Background(), &pb.SubtractRequest{N1: float64(num1), N2: float64(num2)})
		if err != nil {
			log.Fatalf("subtract failed %v", err)
		}
		fmt.Printf("subtract result is %.2f\n", subtractres.Response)

	case 3:
		multiplyres, err := client.Multiply(context.Background(), &pb.MultiplyRequest{N1: float64(num1), N2: float64(num2)})
		if err != nil {
			log.Fatalf("multiplying failed %v", err)
		}
		fmt.Printf("multiplying result is %.2f\n", multiplyres.Response)

	case 4:
		if num2 == 0 {
			fmt.Print("num2 can not be 0, enter a new value for num2: ")
			_, err2 := fmt.Scan(&num2)
			if err != nil {
				log.Fatalf("error : %v", err2)
			}
		}
		divideres, err := client.Divide(context.Background(), &pb.DivideRequest{N1: float64(num1), N2: float64(num2)})
		if err != nil {
			log.Fatalf("division failed %v", err)
		}
		fmt.Printf("division result is %.2f\n", divideres.Response)

	}

}
