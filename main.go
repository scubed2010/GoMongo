package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Request Do This
type Request struct {
	Collection  string      `uri:"collection" binding:"required"`
	ID          string      `uri:"id"`
	Coach       Coach       `json:"Coach"`
	ContentArea ContentArea `json:"ContentArea"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(c)
		c.JSON(200, gin.H{
			"message": healthCheck(),
		})
	})

	r.GET("/read/:collection", func(c *gin.Context) {
		var r Request
		if err := c.ShouldBindUri(&r); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		switch r.Collection {
		case "coaches":
			m := Coach.getAll(Coach{}, r.Collection)
			c.JSON(200, m)
		case "content-areas":
			m := ContentArea.getAll(ContentArea{}, r.Collection)
			c.JSON(200, m)
		default:
			c.JSON(400, gin.H{"msg": r.Collection + " does not exist"})
		}
	})

	r.GET("/read/:collection/:id", func(c *gin.Context) {
		var r Request
		if err := c.ShouldBindUri(&r); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		m := Coach.getByID(Coach{}, r.Collection, r.ID)

		c.JSON(200, m)
	})

	r.POST("/create/:collection/", func(c *gin.Context) {
		var r Request
		if err := c.ShouldBindUri(&r); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		if err := c.BindJSON(&r); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		m := Coach.create(Coach{}, r.Collection, r.Coach)

		c.JSON(200, m)
	})

	r.PUT("/update/:collection", func(c *gin.Context) {
		var r Request
		if err := c.ShouldBindUri(&r); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		if err := c.BindJSON(&r); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		m := Coach.update(Coach{}, r.Collection, r.Coach)

		c.JSON(200, m)
	})

	r.DELETE("/delete/:collection/:id", func(c *gin.Context) {
		var r Request
		if err := c.ShouldBindUri(&r); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		m := Coach.delete(Coach{}, r.Collection, r.ID)

		c.JSON(200, m)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
