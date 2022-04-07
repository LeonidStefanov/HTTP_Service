package service

import (
	"log"

	"github.com/LeonidStefanov/HTTP_Service/pkg/models"
)

type Database interface {
	MakeFriends(userID1, userID2 int) error
	ChangeAge(userID int, newAge int) error
	CreateUser(user *models.User) error
	GetUsers() (models.Users, error)
	GetFriends(userID int) ([]string, error)
	DeleteUser(userID int) error
}

type Service interface {
	MakeFriends(userID1, userID2 int) error
	ChangeAge(userID int, newAge int) error
	CreateUser(user *models.User) error
	GetUsers() (models.Users, error)
	GetFriends(userID int) ([]string, error)
	DeleteUser(userID int) error
}

type service struct {
	db Database
}

func NewService(db Database) *service {
	return &service{
		db: db,
	}
}

func (s *service) GetUsers() (models.Users, error) {
	data, err := s.db.GetUsers()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data, nil
}

func (s *service) CreateUser(user *models.User) error {
	err := s.db.CreateUser(user)
	if err != nil {
		return err

	}

	return nil

}

func (s *service) GetFriends(userID int) ([]string, error) {
	data, err := s.db.GetFriends(userID)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil

}

func (s *service) ChangeAge(userID int, newAge int) error {
	err := s.db.ChangeAge(userID, newAge)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) MakeFriends(userID1, userID2 int) error {
	err := s.db.MakeFriends(userID1, userID2)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteUser(userID int) error {
	err := s.db.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}
