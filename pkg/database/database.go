package database

import (
	"errors"
	"fmt"
	"home/leonid/Git/Pract/network/pkg/models"
	"sync"
)

type database struct {
	sync.RWMutex
	users map[int]*models.User
}

func NewDB() *database {
	return &database{
		users: make(map[int]*models.User),
	}
}

func (d *database) GetUsers() ([]*models.User, error) {
	d.RLock()
	defer d.RUnlock()

	users := []*models.User{}

	for _, v := range d.users {
		users = append(users, v)
	}
	for _, v := range users {
		fmt.Println(v)
	}

	return users, nil
}

func (d *database) CreateUser(user *models.User) error {
	d.Lock()
	defer d.Unlock()
	d.users[user.ID] = user

	return nil
}

func (d *database) GetFiends(userID int) ([]string, error) {
	d.RLock()
	defer d.RUnlock()
	data := []string{}
	u, ok := d.users[userID]
	if !ok {
		return nil, errors.New("no such user")
	}
	data = append(data, u.Friends...)

	return data, nil
}

func (d *database) ChangeAge(userID int, age int) error {
	d.Lock()
	defer d.Unlock()

	fmt.Println(d.users, userID)

	u, ok := d.users[userID]
	if !ok {
		return errors.New("no such user")
	}
	u.Age = age
	d.users[userID] = u

	return nil

}

func (d *database) MakeFfriends(userID1, userID2 int) error {
	fmt.Println(d.users)
	fmt.Println(userID1, userID2)
	u, ok := d.users[userID1]
	if !ok {
		return errors.New("no such user")
	}
	u2, ok := d.users[userID2]
	if !ok {
		return errors.New("no such user")
	}

	u.Friends = append(u.Friends, u2.Name)
	u2.Friends = append(u2.Friends, u.Name)

	d.users[userID2] = u
	d.users[userID2] = u2

	return nil

}

func (d *database) DeleteUser(userID int) error {
	u, ok := d.users[userID]
	if !ok {
		return errors.New("no such user")
	}
	for _, v := range d.users {

		for i, j := range v.Friends {

			if j == u.Name {

				v.Friends = RemoveIndex(u.Friends, i)

			}

		}

	}

	delete(d.users, userID)

	return nil
}
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
