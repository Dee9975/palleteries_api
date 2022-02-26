package main

import (
	"context"
	"net/http"
	"os"
	"palleteries_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	initDb()
	r := gin.Default()
	r.GET("/employees", getEmployees)
	r.POST("/employees", addEmployee)
	r.DELETE("/employees/:id", removeEmployee)

	r.Run()
}

func initDb() {
	err := mgm.SetDefaultConfig(nil, "palleteries", options.Client().ApplyURI("mongodb+srv://matiss:esme9975@traveldatabase.ejnqn.mongodb.net/palleteries?retryWrites=true&w=majority"))

	if err != nil {
		os.Exit(0)
	}
}

func getEmployees(c *gin.Context) {
	result := []models.Employee{}
	err := mgm.Coll(&models.Employee{}).SimpleFind(&result, bson.M{})

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, result)
}

func addEmployee(c *gin.Context) {
	employee := models.Employee{}

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := mgm.Coll(&employee).Create(&employee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &employee)
}

func removeEmployee(c *gin.Context) {
	id := c.Param("id")

	mId, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := mgm.Coll(&models.Employee{}).DeleteOne(context.Background(), bson.M{"id": mId})

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &result)
}
