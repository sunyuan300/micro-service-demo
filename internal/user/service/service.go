package service

import (
	v1 "github.com/kratos-portal/api/user/v1"
	"github.com/kratos-portal/internal/user/biz"
)

// UserService 实现gRPC
type UserService struct {
	v1.UnimplementedUserServiceServer
	user *biz.UserCase
}
