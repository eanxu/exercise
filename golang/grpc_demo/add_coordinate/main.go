package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	pb "grpc_demo/add_coordinate/proto"
)

func main() {
	// 连接服务器
	//conn, err := grpc.Dial("192.168.0.216:29091", grpc.WithInsecure())
	conn, err := grpc.Dial(":9091", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	c := pb.NewAddCoordinateServiceClient(conn)
	// 调用服务端的SayHello
	r, err := c.AddCoordinate(ctx,
		&pb.AddCoordinateRequest{
			Tiff: "/home/marshmallow/Desktop/vm/vm/temp/GF2_PMS1_E112.6_N33.3_20200919_L1A0005072297-MSS1.tiff",
			Xml:  "/home/marshmallow/Desktop/vm/vm/temp/GF2_PMS1_E112.6_N33.3_20200919_L1A0005072297-MSS1.xml",
		},
	)
	if err != nil {
		fmt.Printf("rpc error: %v", err)
	}
	fmt.Println(r)
}