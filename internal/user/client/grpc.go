package main

import (
	"context"
	"fmt"
	v1 "github.com/kratos-portal/api/user/v1"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	conn,err := grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := v1.NewUserServiceClient(conn)
	reply,err := client.GetUser(context.Background(),&v1.GetUserRequest{Id: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}