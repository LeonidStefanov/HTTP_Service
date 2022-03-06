package service

import (
	"home/leonid/Git/Pract/network/pkg/models"
	"log"
)

type Database interface {
	MakeFfriends(userID1, userID2 int) error
	ChangeAge(userID int, age int) error
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	GetFiends(id int) ([]string, error)
	DeleteUser(userID int) error
}

type Service interface {
	MakeFfriends(userID1, userID2 int) error
	ChangeAge(userID int, age int) error
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	GetFiends(id int) ([]string, error)
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

func (s *service) GetUsers() ([]*models.User, error) {
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

func (s *service) GetFiends(id int) ([]string, error) {
	data, err := s.db.GetFiends(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil

}

func (s *service) ChangeAge(userID int, age int) error {
	err := s.db.ChangeAge(userID, age)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) MakeFfriends(userID1, userID2 int) error {
	err := s.db.MakeFfriends(userID1, userID2)
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
