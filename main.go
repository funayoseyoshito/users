package main

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Users struct {
	Id        int `xorm:"pk"`
	Name      string
	Email     string
	UpdatedAt *time.Time
	CreatedAt *time.Time `xorm:"updated"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:33306)/user_info")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Create
	r.POST("/users", func(c *gin.Context) {
		u := &Users{
			Name:  c.PostForm("name"),
			Email: c.PostForm("email"),
		}
		engine.Insert(u)
		c.JSON(200, u)
	})
	//Read
	r.GET("/users/:id", func(c *gin.Context) {
		u := &Users{}
		engine.Where("id = ?", c.Param("id")).Get(u)
		c.JSON(200, u)
	})
	//Update
	r.PATCH("/users/:id", func(c *gin.Context) {
		u := &Users{
			Name:  c.PostForm("name"),
			Email: c.PostForm("email"),
		}
		affected, _ := engine.Id(c.Param("id")).Update(u)
		c.JSON(200, affected)
	})
	//Delete
	r.DELETE("/users/:id", func(c *gin.Context) {
		effected, _ := engine.Id(c.Param("id")).Delete(&Users{})
		c.JSON(200, effected)
	})
	//List
	r.GET("/users", func(c *gin.Context) {
		u := &[]Users{}
		engine.Find(u)
		c.JSON(200, u)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
