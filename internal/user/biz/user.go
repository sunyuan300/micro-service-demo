package biz

import (
	"context"
	"time"
)

//User DO 领域对象
type User struct {
	ID			int64
	Age 		int64
	Name 		string
	Address 	string
	CreateAt 	time.Time
	UpdateAt 	time.Time
}

// UserRepo DO依赖的存储抽象
type UserRepo interface {
	//db
	CreateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, id int64,user *User) error
	GetUser(ctx context.Context, id int64) (*User,error)
	ListUser(ctx context.Context) ([]*User,error)
}

// UserCase 对上(service): 接收DO进行业务处理，对下(data): 透传DO进行数据存储。
type UserCase struct {
	repo UserRepo
}

func NewUserCase(repo UserRepo) *UserCase {
	return &UserCase{repo: repo}
}

func (uc *UserCase) Create(ctx context.Context,user *User) error {
	return uc.repo.CreateUser(ctx,user)
}

func (uc *UserCase) Delete(ctx context.Context,id int64) error {
	return uc.repo.DeleteUser(ctx,id)
}

func (uc *UserCase) Update(ctx context.Context,id int64,user *User) error {
	return uc.repo.UpdateUser(ctx,id,user)
}

func (uc *UserCase) Get(ctx context.Context,id int64) (*User, error) {
	return uc.repo.GetUser(ctx,id)
}

func (uc *UserCase) List(ctx context.Context) ([]*User, error) {
	return uc.repo.ListUser(ctx)
}