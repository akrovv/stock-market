package main

import (
	"fmt"

	"github.com/akrovv/exchange/internal/models/protos/gen/go/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	client := exchange.NewExchangeClient(conn)
}
