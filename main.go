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

	r.GET("/brigades")
	r.POST("/brigades")
	r.DELETE("/brigades/:id")

	r.POST("/send_day", sendDay)

	r.GET("/history")

	r.PUT("/team")

	r.GET("/settings")
	r.POST("/settings")

	r.Run()
}

func initDb() {
	err := mgm.SetDefaultConfig(nil, "palleteries", options.Client().ApplyURI("mongodb+srv://matiss:esme9975@traveldatabase.ejnqn.mongodb.net/palleteries?retryWrites=true&w=majority"))

	if err != nil {
		os.Exit(0)
	}
}

func sendDay(c *gin.Context) {
	team := models.Team{}

	if err := c.BindJSON(&team); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	finalTeam := calculateSalaries(team)

	if err := mgm.Coll(&finalTeam).Create(&finalTeam); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, finalTeam)
}

func calculatePayout(p models.Plank) float64 {
	if p.Zkv {
		if p.D9 {
			return 7.2
		}
		return 8.2
	}

	if p.Type == 0 {
		if p.D9 {
			return 10.8
		}
		return 14.5
	}
	return 12.15
}

func calculateTotal(tara []models.Plank) float64 {
	total := 0.0

	for _, p := range tara {
		payout := calculatePayout(p)
		total += payout * p.Volume
	}

	return total
}

func calculateDayPay(total float64, members []models.TeamMember) float64 {
	workingMembers := []models.TeamMember{}
	totalHours := 0

	for _, m := range members {
		if !m.Forklift {
			totalHours += m.Hours
			workingMembers = append(workingMembers, m)
		}
	}

	return total / float64((len(workingMembers) + totalHours/8))
}

func calculateSalaries(team models.Team) models.Team {
	finalMembers := []models.TeamMember{}

	total := calculateTotal(team.Planks)
	dayPay := calculateDayPay(total, team.Members)

	for _, m := range team.Members {
		salary := dayPay
		if m.Forklift {
			salary += (0.2 * dayPay)
		}

		if m.Kalts {
			for _, p := range team.Planks {
				if p.Kalts {
					salary += p.Volume * 0.22
				}
			}
		}

		if m.ExtraHours > 0 {
			salary += (4.65 * float64(m.ExtraHours))
		}

		if m.Hours < 8 {
			splitMember := m
			splitMember.Salary = (dayPay / 8.0) * float64(m.Hours)
			finalMembers = append(finalMembers, splitMember)
			continue
		}
		finalMember := m
		finalMember.Salary = salary
		finalMembers = append(finalMembers, finalMember)
	}

	finalTeam := team
	finalTeam.Members = finalMembers

	return finalTeam
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
