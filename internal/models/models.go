package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Guardian struct {
	Name             string `bson:"name"`
	Relationship     string `bson:"relationship"`
	Phone            string `bson:"phone"`
	AlternativePhone string `bson:"alternativePhone"`
	Email            string `bson:"email"`
}

type FellowshipInfo struct {
	SeniorCell             string `bson:"seniorCell"`
	FoundationSchoolStatus string `bson:"foundationSchoolStatus"`
	LeadershipRole         string `bson:"leadershipRole"` // STRING
	DesignationCell        string `bson:"designationCell"`
}

type Member struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	FullName        string             `bson:"fullName"`
	Phone           string             `bson:"phone"`
	Email           string             `bson:"email"`
	HomeAddress     string             `bson:"homeAddress"`
	DateOfBirth     time.Time          `bson:"dateOfBirth"`

	Occupation      string             `bson:"occupation"`
	CurrentEmployer string             `bson:"currentEmployer"`
	ImageURL       string              `bson:"imageUrl,omitempty" json:"imageUrl,omitempty"` 

	Guardian   Guardian       `bson:"guardian"`
	Fellowship FellowshipInfo `bson:"fellowship"`

	IsNewMember bool      `bson:"isNewMember"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}
