package main

import (
	"./data"
	"./models"

	"github.com/gin-gonic/gin"
)

// Request - Object that all incoming requests are mapped to
type Request struct {
	Collection  string             `uri:"collection" binding:"required"`
	ID          string             `uri:"id"`
	Coach       models.Coach       `json:"Coach"`
	ContentArea models.ContentArea `json:"ContentArea"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", healthCheck)
	r.GET("/read/:collection", readAll)
	r.GET("/read/:collection/:id", readOne)
	r.POST("/create/:collection/", create)
	r.PUT("/update/:collection", update)
	r.DELETE("/delete/:collection/:id", delete)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": data.HealthCheck(),
	})
}

func readAll(c *gin.Context) {
	var r Request
	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	switch r.Collection {
	case "coaches":
		m := models.Coach.GetAll(models.Coach{}, r.Collection)
		c.JSON(200, m)
	case "content-areas":
		m := models.ContentArea.GetAll(models.ContentArea{}, r.Collection)
		c.JSON(200, m)
	default:
		c.JSON(400, gin.H{"msg": r.Collection + " does not exist"})
	}
}

func readOne(c *gin.Context) {
	var r Request
	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	m := models.Coach.GetByID(models.Coach{}, r.Collection, r.ID)

	c.JSON(200, m)
}

func create(c *gin.Context) {
	var r Request
	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if err := c.BindJSON(&r); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	m := models.Coach.Create(models.Coach{}, r.Collection, r.Coach)

	c.JSON(200, m)
}

func update(c *gin.Context) {
	var r Request
	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if err := c.BindJSON(&r); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	m := models.Coach.Update(models.Coach{}, r.Collection, r.Coach)

	c.JSON(200, m)
}

func delete(c *gin.Context) {
	var r Request
	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	m := models.Coach.Delete(models.Coach{}, r.Collection, r.ID)

	c.JSON(200, m)
}
