package handler

import (
	user "apiInterview/pkg/type"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Error string `json:"error"`
}

// Error code 400 (BadRequest) - binding error
// Error code 403 (Forbidden) - input data error

func ErrHandler(ctx *gin.Context, err error, errCode int, usr any, succesAnswer bool) bool {
	if err != nil {
		logrus.Error(err)
		fmt.Print(err.Error())
		ctx.JSON(errCode, Response{
			Error: err.Error(),
		})
		return false
	}
	if succesAnswer {
		ctx.JSON(http.StatusOK, usr)
	}
	return true
}

func (handler *Handler) CreateUser(ctx *gin.Context) {
	var user user.User

	err := ctx.BindJSON(&user)
	if !ErrHandler(ctx, err, http.StatusBadRequest, &user, false) {
		return
	}
	ErrHandler(ctx, handler.srvc.CreateUser(user.Username, user.Password, user.Description), http.StatusForbidden, &user, true)
}

func (handler *Handler) UpdateUser(ctx *gin.Context) {
	var user user.User
	usrname := ctx.Param("username")
	err := ctx.BindJSON(&user)
	if !ErrHandler(ctx, err, http.StatusBadRequest, &user, false) {
		return
	}
	user.Username = usrname
	ErrHandler(ctx, handler.srvc.UpdateUser(usrname, user.Password, user.Description), http.StatusForbidden, &user, true)
}

func (handler *Handler) GetUser(ctx *gin.Context) {
	res := &user.User{}
	usrname := ctx.Param("username")
	res, err := handler.srvc.GetUser(usrname)
	ErrHandler(ctx, err, http.StatusForbidden, res, true)

}

func (handler *Handler) DeleteUser(ctx *gin.Context) {
	usrname := ctx.Param("username")
	ErrHandler(ctx, handler.srvc.DeleteUser(usrname), http.StatusForbidden, nil, true)
}

func (handler *Handler) GetAllUsers(ctx *gin.Context) {
	users, err := handler.srvc.GetAllUsers()
	ErrHandler(ctx, err, http.StatusForbidden, users, true)
}
