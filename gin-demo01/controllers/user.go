package controllers

import (
	_ "os/user"
	"strconv"

	"github.com/28267/goDemo/gin-demo01/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// func (u UserController) GetUserInfo(c *gin.Context) {
// 	idStr := c.Param("id")
// 	// name := c.Param("name")

// 	id, _ := strconv.Atoi(idStr)
// 	user, _ := models.GetUserTest(id)
// 	ReturnSuccess(c, 0, "success", user, 1)

// 	// defer func() {
// 	// 	if err := recover(); err != nil {
// 	// 		fmt.Println(err, "--catch error")
// 	// 	}
// 	// }()
// 	// var a = 5
// 	// var b = 0
// 	// var p = a / b
// 	// fmt.Println(p)
// }

func (u UserController) GetList(c *gin.Context) {
	// logger.Write("日志信息", "user")
	ReturnError(c, 404, "none info List")
}

func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "save error")
		return
	}
	ReturnSuccess(c, 0, "save success", id, 1)

}

func (u UserController) UpdateUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(id, username)
	ReturnSuccess(c, 0, "update success", true, 1)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4002, "delete error")
		return
	}
	ReturnSuccess(c, 0, "delete success", id, 1)
}

func (u UserController) GetUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	user, err := models.GetUser(id)
	if err != nil {
		ReturnError(c, 4002, "none info List")
		return
	}
	ReturnSuccess(c, 0, "success", user, 1)
}
