package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "time"
)

// func run(message string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(message)
// 	}
// }

// func channel(ch chan string, message string) {
// 	time.Sleep(3 * time.Second)
// 	var data = fmt.Sprintf(message)
// 	ch <- data
// }
// func main() {
// 	var message = make(chan string)
// 	go channel(message, "channel 1")
// 	go channel(message, "channel 2")
// 	go channel(message, "channel 3")

// 	var message1 = <-message
// 	fmt.Println(message1)

// 	var message2 = <-message
// 	fmt.Println(message2)

// 	var message3 = <-message
// 	fmt.Println(message3)
// }

type car struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Type  string `json:"car_type"`
}

var cars = []car{
	{ID: "1", Brand: "Honda", Type: "City"},
	{ID: "2", Brand: "Toyota", Type: "Avanza"},
}

func main() {
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	// GET /cars - list cars
	r.GET("/cars", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cars)
	})
	// POST /cars - create car
	r.POST("/cars", func(ctx *gin.Context) {
		var car car
		if err := ctx.ShouldBindJSON(&car); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		cars = append(cars, car)
		ctx.JSON(http.StatusCreated, car)
	})
	// PUT /cars/:id - update car
	r.PUT("/cars/:id", func(ctx *gin.Context) {
		id := ctx.Param("car_id")
		for i, car := range cars {
			if car.ID == id {
				cars = append(cars[:i], cars[i+1:]...)
				break
			}
		}
		ctx.Status(http.StatusCreated)
	})
	// DELETE /cars/;id - delete car
	r.DELETE("/cars/:id", func(ctx *gin.Context) {
		id := ctx.Param("car_id")
		for i, car := range cars {
			if car.ID == id {
				cars = append(cars[:i], cars[i+1:]...)
				break
			}
		}
		ctx.Status(http.StatusNoContent)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
