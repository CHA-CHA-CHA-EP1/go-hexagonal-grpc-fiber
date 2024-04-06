package domain

import "github.com/google/uuid"

type User struct {
    ID         uuid.UUID 
    FirstName  string   `json:"first_name" bson:"first_name"`
    LastName   string  `json:"last_name" bson:"last_name"`
    Email      string   `json:"email" bson:"email"`  
    Password   string   `json:"password" bson:"password"`
}

type UserRegistration struct {
    FirstName string `json:"first_name" validate:"required,min=3,max=30"`
    LastName string `json:"last_name" validate:"required,min=3,max=30"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}


type UserLogin struct {
    Email    string `json:"email" validate:"required,email"`
    Password   string   `json:"password" bson:"password"`
}
