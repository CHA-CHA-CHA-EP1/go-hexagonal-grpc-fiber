package repository

import (
	"context"
	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/internal/auth/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
    GetById(id uint) (*domain.User, error)
    GetByEmail(email string) (*domain.User, error)
    Create(user *domain.UserRegistration) error
    Update(user *domain.User) error
}

type userRepository struct {
    db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
    return &userRepository{
        db: db,
    }
}

func (r *userRepository) GetById(id uint) (*domain.User, error){
    var user domain.User
    ctx := context.Background()

    if err := r.db.Collection("users").FindOne(ctx, bson.M{"id": id}).Decode(&user); err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error){
    var user domain.User
    ctx := context.Background()

    if err := r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepository) Create(user *domain.UserRegistration) error {
    ctx := context.Background()

    _, err := r.db.Collection("users").InsertOne(ctx, user)
    if err != nil {
        return err
    }

    return nil
}

func (r *userRepository) Update(user *domain.User) error {
    ctx := context.Background()

    _, err := r.db.Collection("users").UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": user})
    if err != nil {
        return err
    }
    
    return nil
}
