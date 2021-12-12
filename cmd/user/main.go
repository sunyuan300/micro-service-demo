package main

import (
	"github.com/kratos-portal/internal/user/biz"
	"github.com/kratos-portal/internal/user/data"
	"github.com/kratos-portal/internal/user/server"
	"github.com/kratos-portal/internal/user/service"
	"log"
)

func main()  {
	dataData,err :=data.NewData()
	if err != nil {
		panic(err)
	}
	userRepo :=data.NewUserRepo(dataData)
	userCase := biz.NewUserCase(userRepo)
	userService := service.NewUserService(userCase)
	srv := server.NewGRPCServer("tcp","127.0.0.1:8080",userService)
	log.Fatalln(srv.Start(srv.GetContext()))
}
