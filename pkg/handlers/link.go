package handlers

import (
	"context"
	"fmt"
	"github.com/MojixCoder/awesomeProject/pkg/config"
	"github.com/MojixCoder/awesomeProject/pkg/db"
	"github.com/MojixCoder/awesomeProject/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hash/adler32"
	"net/http"
	"strconv"
	"time"
)

// Repo is the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	AppConfig *config.Config
	Client *mongo.Client
}

// NewRepo creates a new repository
func NewRepo(a *config.Config, client *mongo.Client) *Repository {
	return &Repository{
		AppConfig: a,
		Client: client,
	}
}

// SetRepo sets the repository for the handlers
func SetRepo(r *Repository) {
	Repo = r
}

/* ----- Handlers ----- */

// ShortenedLinkCreate creates a shortened link
func (repo *Repository) ShortenedLinkCreate(c *gin.Context) {
	var linkBody models.LinkCreate
	var link models.Link

	if err := c.ShouldBindJSON(&linkBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&linkBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	client := repo.Client
	linkCollection := db.GetCollection(client, "link")

	count, err := linkCollection.CountDocuments(context.TODO(), bson.M{"link": linkBody.Link})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "couldn't get collection link",
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": fmt.Sprintf("an object with this link(%s) already exists.", linkBody.Link),
		})
		return
	}

	link.ID = primitive.NewObjectID()
	finalHash := strconv.FormatUint(uint64(adler32.Checksum([]byte((link.ID.Hex())))), 32)
	link.Slug = finalHash
	link.Link = linkBody.Link
	link.ShortenedLink = repo.AppConfig.BaseURL + "/" + finalHash
	link.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	result, err := linkCollection.InsertOne(context.TODO(), link)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "unable to insert object.",
		})
		return
	}

	c.JSON(http.StatusCreated, result)

}

func (repo *Repository) RedirectToLink(c *gin.Context) {
	slug := c.Param("slug")
	var result models.Link
	client := repo.Client
	linkCollection := db.GetCollection(client, "link")
	err := linkCollection.FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": fmt.Sprintf("object with slug(%s) not found.", slug),
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, result.Link)
}
