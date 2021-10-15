package controllers

import (
	"context"
	"golang-restaurant-management/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
     result, err := menuCollection.Find(context.TODO(), bson.M{})
		 defer cancel()
		 if err != nil {
			 c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetch data"})
		 }
		 var allmenus = []bson.M
		 if err := result.All(ctx, &allmenu); err != nil {
			 log.Fatal(err)
		 }

		 c.JSON(http.StatusOK, allmenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
