package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fuhiz/docker-go-sample/app/pkg/connecter"
	"github.com/fuhiz/docker-go-sample/app/pkg/model"
)

type UserController struct{}

type UserParam struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
	Age  int    `json:"age" binding:"required,number"`
}

// ユーザー取得
func (self *UserController) GetUser(c *gin.Context) {
	ID := c.Params.ByName("id")
	userID, _ := strconv.Atoi(ID)
	user, err := model.GetUserById(connecter.DB(), userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// ユーザー一覧
func (self *UserController) Index(c *gin.Context) {
	users, err := model.GetUsers(connecter.DB())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user search failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// ユーザー作成
func (self *UserController) CreateUser(c *gin.Context) {
	var param UserParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := model.NewUser(param.Name, param.Age)
	user, err := model.CreateUser(connecter.DB(), newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user create failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// ユーザー更新
func (self *UserController) UpdateUser(c *gin.Context) {
	ID := c.Params.ByName("id")
	userID, _ := strconv.Atoi(ID)
	user, err := model.GetUserById(connecter.DB(), userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var param UserParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateParam := map[string]interface{}{
		"name": param.Name,
		"age":  param.Age,
	}

	_, err = user.Update(connecter.DB(), updateParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// ユーザー削除
func (self *UserController) DeleteUser(c *gin.Context) {
	ID := c.Params.ByName("id")
	userID, _ := strconv.Atoi(ID)
	user, err := model.GetUserById(connecter.DB(), userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	_, err = user.Delete(connecter.DB())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
