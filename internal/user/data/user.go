package data

import (
	"context"
	"github.com/kratos-portal/internal/user/biz"
)

// userRepo 存储抽象的实现
type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{data:data}
}

func (ur *userRepo) CreateUser(ctx context.Context,user *biz.User) error {
	return ur.data.db.Create(user).Error
}

func (ur *userRepo) DeleteUser(ctx context.Context,id int64) error {
	return ur.data.db.Delete(biz.User{},"id=?",id).Error
}

func (ur *userRepo) UpdateUser(ctx context.Context,id int64,user *biz.User) error {
	return ur.data.db.Updates(user).Where("id=?",id).Error
}

func (ur *userRepo) GetUser(ctx context.Context,id int64) (u *biz.User,err error) {
	//err = ur.data.db.Take(u,"id=?",id).Error
	u = &biz.User{ID: id,Name: "hello"}
	return
}

func (ur *userRepo) ListUser(ctx context.Context) (us []*biz.User,err error){
	err = ur.data.db.Take(us).Error
	return
}