package service

import (
    "errors"
	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/internal/auth/domain"
	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/internal/auth/repository"
)

type UserService interface {
    GetUserById(id uint) (*domain.User, error)
    GetUserByEmail(email string) (*domain.User, error)
    RegisterUser(user domain.UserRegistration) (*domain.User, error)
}

type userService struct {
    userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
    return &userService{
        userRepository: ur,
    }
}

func (us *userService) GetUserById(id uint) (*domain.User, error) {
    if id == 0 {
        return nil, errors.New("invalid id")
    }

    user, err := us.userRepository.GetById(id)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (us *userService) GetUserByEmail(email string) (*domain.User, error) {
    if email == "" {
        return nil, errors.New("invalid email")
    }

    user, err := us.userRepository.GetByEmail(email)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (us *userService) RegisterUser(user domain.UserRegistration) (*domain.User, error) {
    err :=  us.userRepository.Create(&user)
    if err != nil {
        return nil, err
    }

    return &domain.User{
        FirstName: user.FirstName,
        LastName: user.LastName,
        Email: user.Email,
    }, nil
}

