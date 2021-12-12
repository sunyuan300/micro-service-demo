package service

import (
	"context"
	v1 "github.com/kratos-portal/api/user/v1"
	"github.com/kratos-portal/internal/user/biz"
)

// NewUserService 实现gRPC
func NewUserService(user *biz.UserCase) *UserService {
	return &UserService{
		user: user,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply,error) {
	err :=s.user.Create(ctx,&biz.User{
		Age: req.Age,
		Name: req.Name,
		Address: req.Address,
	})
	return &v1.CreateUserReply{},err
}

func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply,error) {
	err := s.user.Delete(ctx,req.Id)
	return &v1.DeleteUserReply{},err
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserUserReply,error){
	err := s.user.Update(ctx,req.Id,&biz.User{
		Age: req.Age,
		Name: req.Name,
		Address: req.Address,
	})
	return &v1.UpdateUserUserReply{},err
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply,error){
	u,err := s.user.Get(ctx,req.Id)
	return &v1.GetUserReply{User:&v1.User{Id:u.ID,Name:u.Name,Age:u.Age,Address:u.Address}},err
}

func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest)(*v1.ListUserReply,error) {
	reply := &v1.ListUserReply{}
	us,err := s.user.List(ctx)
	for _,u := range us {
		reply.Users = append(reply.Users,&v1.User{
			Id:      u.ID,
			Name:    u.Name,
			Age:     u.Age,
			Address: u.Address,
		})
	}
	return reply,err
}