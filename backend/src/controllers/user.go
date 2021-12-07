package controllers

import (
	"net/http"
	"strconv"

	"github.com/dogerescat/go-app/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (uc *UserController) Create(c *gin.Context) {
	u := models.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	err := u.Create()
	if err != nil {
		c.JSON(404, gin.H{"result": "false"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "true"})
}

func (uc *UserController) Read(c *gin.Context) {
	id := c.Params.ByName("id")
	i, _ := strconv.Atoi(id)
	u, err := models.Read(i)
	if err != nil {
		c.JSON(404, gin.H{"result": "false"})
	}
	c.JSON(http.StatusOK, u)
}
