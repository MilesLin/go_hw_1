package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8080")
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	for _, role := range Data {
		if role.ID == uint(id) {
			c.JSON(http.StatusOK, role)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func Post(c *gin.Context) {
	var r Role
	if err := c.ShouldBindJSON(&r); err != nil {
		c.Error(err)
		return
	}
	latest := Data[len(Data)-1]
	nextId := latest.ID + 1
	r.ID = nextId
	Data = append(Data, r)
	c.JSON(http.StatusOK, r)
}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

func Put(c *gin.Context) {
	var r RoleVM

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	for i := 0; i < len(Data); i++ {
		if Data[i].ID == uint(id) {
			Data[i].Name = r.Name
			Data[i].Summary = r.Summary
			c.JSON(http.StatusOK, Data[i])
			return
		}
	}

	c.Status(http.StatusNotFound)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := 0; i < len(Data); i++ {
		if Data[i].ID == uint(id) {
			Data = append(Data[0:i], Data[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.Status(http.StatusNotFound)
}
