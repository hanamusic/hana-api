package controllers

import (
	"context"
	"hana-api/models"
	mongo_models "hana-api/models/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAlbum(ginContext *gin.Context) {
	collection := mongoClient.Database("hana-db").Collection("albums")

	s := mongo_models.Album{
		ID:        primitive.NewObjectID(),
		AlbumID:   uuid.New().String(),
		AlbumName: ginContext.Query("album_name"),
	}

	_, err := collection.InsertOne(context.TODO(), s)

	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		return
	}

	ginContext.IndentedJSON(http.StatusOK, models.Result{
		IsSuccessful: true,
		Message:      "Added a new album",
	})
}
