package dbmock

import (
	"errors"
	"fmt"

	"sync"

	"github.com/LeonidStefanov/HTTP_Service/pkg/models"
)

type database struct {
	sync.RWMutex
	users map[int]*models.User
}

func NewMock() *database {
	user1 := models.User{ID: 40, Name: "Ivan", Age: 23, Friends: []string{"Ran"}}
	user2 := models.User{ID: 20, Name: "Han", Age: 23, Friends: []string{"Ran"}}
	user3 := models.User{ID: 66, Name: "Tan", Age: 23, Friends: []string{"Ran"}}

	m := map[int]*models.User{40: &user1, 20: &user2, 66: &user3}
	return &database{
		users: m,
	}
}

func (d *database) GetUsers() (models.Users, error) {
	d.RLock()
	defer d.RUnlock()

	var allUsers models.Users

	for _, v := range d.users {

		allUsers = append(allUsers, *v)

	}

	fmt.Println(allUsers)

	return allUsers, nil
}

func (d *database) CreateUser(user *models.User) error {
	d.Lock()
	defer d.Unlock()
	d.users[user.ID] = user

	return nil
}

func (d *database) GetFriends(userID int) ([]string, error) {
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

func (d *database) ChangeAge(userID int, newAge int) error {
	d.Lock()
	defer d.Unlock()

	fmt.Println(d.users, userID)

	u, ok := d.users[userID]
	if !ok {
		return errors.New("no such user")
	}
	u.Age = newAge
	d.users[userID] = u

	return nil

}

func (d *database) MakeFriends(userID1, userID2 int) error {

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

	for _, v := range d.users {

		fmt.Println(v)

	}

	return nil
}
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
