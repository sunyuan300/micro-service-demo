// Code generated by github.com/sunyuan300/protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	http "net/http"
)

// context.gin.http.

type UserServiceHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)

	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)

	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserUserReply, error)

	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)

	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
}

func RegisterUserServiceHTTPServer(group *gin.RouterGroup, srv UserServiceHTTPServer) {

	group.Handle("POST", "/v1/user", _UserService_CreateUser_HTTP_Handler(srv))

	group.Handle("DELETE", "/v1/user/{id}", _UserService_DeleteUser_HTTP_Handler(srv))

	group.Handle("PUT", "/v1/user/{id}", _UserService_UpdateUser_HTTP_Handler(srv))

	group.Handle("GET", "/v1/user/{id}", _UserService_GetUser_HTTP_Handler(srv))

	group.Handle("GET", "/v1/user", _UserService_ListUser_HTTP_Handler(srv))

}

func _UserService_CreateUser_HTTP_Handler(srv UserServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in CreateUserRequest

		if err := ctx.ShouldBindJSON(&in); err != nil {
			return
		}

		out, err := srv.CreateUser(ctx, &in)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func _UserService_DeleteUser_HTTP_Handler(srv UserServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in DeleteUserRequest

		if err := ctx.ShouldBindUri(&in); err != nil {
			return
		}

		if err := ctx.ShouldBindQuery(&in); err != nil {
			return
		}

		out, err := srv.DeleteUser(ctx, &in)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func _UserService_UpdateUser_HTTP_Handler(srv UserServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in UpdateUserRequest

		if err := ctx.ShouldBindUri(&in); err != nil {
			return
		}

		if err := ctx.ShouldBindJSON(&in); err != nil {
			return
		}

		out, err := srv.UpdateUser(ctx, &in)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func _UserService_GetUser_HTTP_Handler(srv UserServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in GetUserRequest

		if err := ctx.ShouldBindUri(&in); err != nil {
			return
		}

		if err := ctx.ShouldBindQuery(&in); err != nil {
			return
		}

		out, err := srv.GetUser(ctx, &in)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func _UserService_ListUser_HTTP_Handler(srv UserServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in ListUserRequest

		if err := ctx.ShouldBindQuery(&in); err != nil {
			return
		}

		out, err := srv.ListUser(ctx, &in)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, out)
	}
}
