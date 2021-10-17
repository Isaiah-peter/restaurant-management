package controllers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"log"
	"net/http"
	"time"

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
		 if err := result.All(ctx, &allmenus); err != nil {
			 log.Fatal(err)
		 }

		 c.JSON(http.StatusOK, allmenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
     var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
	 menuId := c.Param("menu_id")
	 var menu models.Menu

	 err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)

	 defer cancel()
	 if err != nil {
		 c.JSON(http.StatusInternalServerError, gin.H{"error":"error occurred while fetching menu items"})
	 }

	 c.JSON(http.StatusOK, menu)

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		 var menu models.Menu
     var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)

		 if err := c.BindJSON(&menu); err != nil {
			 c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			 return
		 }

		 validateErr := validate.Struct(menu)
 		if validateErr != nil {
 			c.JSON(http.StatusBadRequest, gin.H{"error":validateErr.Error()})
 			return
 		}

		menu.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.DeletedAt,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()

		menu.MenuId = menu.ID.Hex()

		result, inserterr := menuCollection.InsertOne(ctx, menu)
		if inserterr != nil {
			msg := fmt.Sprintf("menu is not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
    defer cancel()

		c.JSON(http.StatusOK, result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
