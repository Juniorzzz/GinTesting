package api

import (
	"GinWeb/serializer"
	"GinWeb/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := e.Field
			tag := e.Tag
			return serializer.Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}

func Register(ctx *gin.Context) {

	var service service.UserRegisterService

	if err := ctx.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err == nil {
			res := serializer.BuildUserResponse(user)
			ctx.JSON(200, res)
		} else {
			ctx.JSON(200, "message: register failed")
		}
	} else {
		ctx.JSON(200, ErrorResponse(err))
	}

}

func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "login success"})
}
