package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

type Headers struct {
	UserId string
}

func (h *TestHandler) TestUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": id,
	})
}

func (h *TestHandler) TestUsers(c *gin.Context) {
	id := c.Query("id")
	names := c.QueryArray("names")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "empty id",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": id,
		"names":  names,
	})
}
func (h *TestHandler) TestHeader(c *gin.Context) {
	userId := c.GetHeader("UserId")
	fmt.Printf("c.Request.Header: %v\n", c.Request.Header)
	c.JSON(http.StatusOK, gin.H{
		"result": c.Request.Header,
		"userId": userId,
	})
}

type PersonData struct {
	Name string `json:"name" binding:"required,alpha,min=4,max=10"`
	Age  int    `json:"age" binding:"required,isAdults"`
	Sex  string `json:"sex"`
}

func (h *TestHandler) TestBody(c *gin.Context) {
	p := PersonData{}
	err := c.Bind(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"body": &p,
	})
}

func (h *TestHandler) TestFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"file": file.Filename,
	})
}
