package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Link is link collection structure
type Link struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
	Slug string `bson:"slug" json:"slug"`
	Link string `bson:"link" json:"link" validate:"required"`
	ShortenedLink string `bson:"shortenedLink" json:"shortenedLink" validate:"required"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt" validate:"required"`
}

// LinkCreate is the structure for validating JSON body in ShortenedLinkCreate
type LinkCreate struct {
	Link string `bson:"link" json:"link" validate:"required"`
}

// Redirect is the structure for validating JSON body in RedirectToLink
type Redirect struct {
	Slug string `bson:"slug" json:"slug" validate:"required"`
}
