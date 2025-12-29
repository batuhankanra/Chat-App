package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type User struct {
	BaseModel `bson:",inline"`

	Username     string   `bson:"username" json:"username"`
	Email        string   `bson:"email" json:"email"`
	PasswordHash string   `bson:"passwordHash" json:"-"`
	Teams        []string `bson:"teams" json:"teams"`
	IsActive     bool     `bson:"isActive" json:"isActive"`
}

type Team struct {
	BaseModel `bson:",inline"`

	Name    string             `bson:"name" json:"name"`
	OwnerID primitive.ObjectID `bson:"ownerId" json:"ownerId"`
	Members []string           `bson:"members" json:"members"`
}

type Channel struct {
	BaseModel `bson:",inline"`

	TeamID    string   `bson:"teamId" json:"teamId"`
	Name      string   `bson:"name" json:"name"`
	IsPrivate bool     `bson:"isPrivate" json:"isPrivate"`
	Members   []string `bson:"members,omitempty" json:"members,omitempty"`
}
type Message struct {
	BaseModel `bson:",inline"`

	ChannelID string `bson:"channelId" json:"channelId"`
	SenderID  string `bson:"senderId" json:"senderId"`
	Content   string `bson:"content" json:"content"`
	Type      string `bson:"type" json:"type"`
}

type Notification struct {
	BaseModel `bson:",inline"`

	UserID string `bson:"userId" json:"userId"`
	Title  string `bson:"title" json:"title"`
	Body   string `bson:"body" json:"body"`
	IsRead bool   `bson:"isRead" json:"isRead"`
}

const (
	UserCollection         = "users"
	TeamCollection         = "teams"
	ChannelCollection      = "channels"
	MessageCollection      = "message"
	NotificationCollection = "notification"
)

type RegisterRequest struct {
	USername string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type CreateTeamRequest struct {
	Name string `json:"name" binding:"required"`
}
type AddTeamMemberRequest struct {
	UserID string `json:"userId" binding:"required"`
}

type SendMessageRequest struct {
	ChannelID string `json:"channelId" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Type      string `json:"type"`
}
