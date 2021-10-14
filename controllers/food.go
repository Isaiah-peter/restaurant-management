package controllers

import (
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()


func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
     var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
	 foodId := c.Param("food_id")
	 var food models.Food

	 err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)

	 defer cancel()
	 if err != nil {
		 c.JSON(http.StatusInternalServerError, gin.H{"error":"error occurred while fetching food items"})
	 }

	 c.JSON(http.StatusOK, food)

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		validateErr := validate.Struct(food)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":validateErr.Error()})
			return
		}

		foodCollection.FindOne(ctx, bson.M{"menu_id":food.MenuId}).Decode(&menu)

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {
	return 0
}

func toFixed(num float64, pricision int) float64 {
	return 0
}
