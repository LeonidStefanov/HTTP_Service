package database

import (
	"fmt"
	"home/leonid/Git/Pract/network/pkg/models"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type database struct {
	db *mgo.Database
	s  *mgo.Session
}

func NewDB(host string, port string, dbName string) (*database, error) {

	conn := fmt.Sprintf("mongodb://%v:%v/skillbox", host, port)
	fmt.Println(conn)

	session, err := mgo.DialWithTimeout(conn, time.Second*5)
	if err != nil {
		log.Println("DialWithTimeout: ", err)
		return nil, err

	}
	db := session.DB(dbName)
	session.Ping()

	fmt.Println("Connected to MongoDB!")

	return &database{
		db: db,
		s:  session,
	}, nil
}

func (base *database) Close() {
	base.db.Session.Close()
}

func (base *database) GetUsers() (models.Users, error) {

	var users models.Users

	err := base.db.C("skillbox").Find(bson.M{}).All(&users)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (base *database) CreateUser(user *models.User) error {

	err := base.db.C("skillbox").Insert(user)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (base *database) GetFriends(userID int) ([]string, error) {
	var user models.User
	err := base.db.C("skillbox").Find(bson.M{"id": userID}).One(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user.Friends, nil
}

func (base *database) ChangeAge(userID int, newAge int) error {

	err := base.db.C("skillbox").Update(bson.M{"id": userID}, bson.M{"$set": bson.M{"age": newAge}})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func (base *database) MakeFriends(userID1, userID2 int) error {
	var user1 models.User
	var user2 models.User

	err := base.db.C("skillbox").Find(bson.M{"id": userID1}).One(&user1)
	if err != nil {
		log.Println(err)
		return err
	}
	err = base.db.C("skillbox").Find(bson.M{"id": userID2}).One(&user2)
	if err != nil {
		log.Println(err)
		return err
	}

	user1.Friends = append(user1.Friends, user2.Name)
	user2.Friends = append(user2.Friends, user1.Name)

	err = base.db.C("skillbox").Update(bson.M{"id": userID1}, user1)
	if err != nil {
		log.Println(err)
		return err
	}
	err = base.db.C("skillbox").Update(bson.M{"id": userID2}, user2)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func (base *database) DeleteUser(userID int) error {

	allUsers, err := base.GetUsers()
	if err != nil {
		log.Println(err)
		return err
	}

	var deleteUsers models.User

	for i := 0; i < len(allUsers); i++ {
		if allUsers[i].ID == userID {
			deleteUsers = allUsers[i]
		}
	}

	for i := 0; i < len(allUsers); i++ {
		for j := 0; j < len(allUsers[i].Friends); j++ {
			if allUsers[i].Friends[j] == deleteUsers.Name {
				allUsers[i].Friends = RemoveIndex(allUsers[i].Friends, j)
				base.db.C("skillbox").Update(bson.M{}, allUsers[i])
			}
		}
	}

	err = base.db.C("skillbox").Remove(bson.M{"id": userID})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
