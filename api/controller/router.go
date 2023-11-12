package controller

import (
	"backend/api/model"
	"log"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"DELETE",
			"PUT",
		},
		AllowHeaders: []string{
			"Authorization",
		},
	}))

	v1 := router.Group("uttcforum")
	{
		v1.GET("/main", contentsGET)
		v1.POST("/create", contentPOST)
		v1.PATCH("/update", contentPATCH)
		v1.DELETE("/delete", contentDELETE)
	}
	router.Run(":8080")
}

func contentsGET(c *gin.Context) {
	contents, err := model.GetContents()
	if err != nil {
		log.Fatalln("err")
	}
	c.JSON(http.StatusOK, gin.H{"contents": contents})
}

func contentPOST(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var typeofContents model.Contents
	if err := decoder.Decode(&typeofContents); err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := model.CreateContent(typeofContents.Title,typeofContents.Category,typeofContents.Curriculum,typeofContents.Content)
	if err != nil {
		log.Fatalln(err)
	}
}

func contentPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	content, err := model.GetContent(uint(id))
	if err != nil {
		log.Fatalln(err)
	}

	title := c.PostForm("title")
	category := c.PostForm("Category")
	curriculum := c.PostForm("Curriculum")
	item := c.PostForm("Content")
	now := time.Now()
	content.Title = title
	content.Category = category
	content.Curriculum = curriculum
	content.Content = item
	content.UpdatedAt = now
	content.UpdateContent()
	c.JSON(http.StatusOK, gin.H{"content": content})
}

func contentDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	todo, err := model.GetContent(uint(id))
	if err != nil {
		log.Fatalln(err)
	}
	err = todo.DeleteContent()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, "Deleted")
}
