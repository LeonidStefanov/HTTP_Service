package transport

import (
	"bytes"
	"fmt"
	"home/leonid/Git/Pract/network/pkg/database"
	"home/leonid/Git/Pract/network/pkg/option"
	"home/leonid/Git/Pract/network/pkg/service"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	port = "8080"

	deleteUserJSON = `{"target_id":40}`

	newAgeJSON = `{"new_age":45}`

	makeFriendsJSON = `{"source_id" :6,"target_id":40}`

	userJSON = `{"id":40,"name":"Min","age":27,"friends":["Jon"]}`
)

func TestCreateUser(t *testing.T) {
	var cfg option.Options
	err := envconfig.Process("service", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg)

	db, err := database.NewDB(cfg.DBHost, cfg.DBPort, "skillbox")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()
	svc := service.NewService(db)

	h := NewTransport(port, svc)
	h.InitEndpoints()
	go h.Start()

	req, _ := http.NewRequest("POST", "http://localhost:8080/create/user", bytes.NewReader([]byte(userJSON)))
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected %v get %v", http.StatusCreated, resp.StatusCode)
	} else {
		buf, _ := ioutil.ReadAll(resp.Body)
		t.Log("OK", string(buf))
	}

}

func TestGetFriends(t *testing.T) {
	var cfg option.Options
	err := envconfig.Process("service", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg)

	db, err := database.NewDB(cfg.DBHost, cfg.DBPort, "skillbox")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	svc := service.NewService(db)

	h := NewTransport(port, svc)
	h.InitEndpoints()
	go h.Start()

	req, _ := http.NewRequest("GET", "http://localhost:8080/friends/5", nil)
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %v get %v", http.StatusOK, resp.StatusCode)
	} else {
		buf, _ := ioutil.ReadAll(resp.Body)
		t.Log("OK", string(buf))
	}

}

func TestGetUsers(t *testing.T) {
	var cfg option.Options
	err := envconfig.Process("service", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg)

	db, err := database.NewDB(cfg.DBHost, cfg.DBPort, "skillbox")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	svc := service.NewService(db)

	h := NewTransport(port, svc)
	h.InitEndpoints()
	go h.Start()

	req, _ := http.NewRequest("GET", "http://localhost:8080/users", nil)
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %v get %v", http.StatusOK, resp.StatusCode)
	} else {
		buf, _ := ioutil.ReadAll(resp.Body)
		t.Log("OK", string(buf))
	}

}

func TestMakeFriends(t *testing.T) {
	var cfg option.Options
	err := envconfig.Process("service", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg)

	db, err := database.NewDB(cfg.DBHost, cfg.DBPort, "skillbox")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	svc := service.NewService(db)

	h := NewTransport(port, svc)
	h.InitEndpoints()
	go h.Start()

	req, _ := http.NewRequest("POST", "http://localhost:8080/make/friends", bytes.NewReader([]byte(makeFriendsJSON)))
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %v get %v", http.StatusOK, resp.StatusCode)
	} else {
		buf, _ := ioutil.ReadAll(resp.Body)
		t.Log("OK", string(buf))
	}

}

func TestChangeAge(t *testing.T) {
	var cfg option.Options
	err := envconfig.Process("service", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg)

	db, err := database.NewDB(cfg.DBHost, cfg.DBPort, "skillbox")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	svc := service.NewService(db)

	h := NewTransport(port, svc)
	h.InitEndpoints()
	go h.Start()

	req, _ := http.NewRequest("PUT", "http://localhost:8080/change/5", bytes.NewReader([]byte(newAgeJSON)))
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %v get %v", http.StatusOK, resp.StatusCode)
	} else {
		buf, _ := ioutil.ReadAll(resp.Body)
		t.Log("OK", string(buf))
	}
}

func TestDeleteUser(t *testing.T) {
	var cfg option.Options
	err := envconfig.Process("service", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg)

	db, err := database.NewDB(cfg.DBHost, cfg.DBPort, "skillbox")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	svc := service.NewService(db)

	h := NewTransport(port, svc)
	h.InitEndpoints()
	go h.Start()

	req, _ := http.NewRequest("DELETE", "http://localhost:8080/delete/user", bytes.NewReader([]byte(deleteUserJSON)))
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %v get %v", http.StatusOK, resp.StatusCode)
	} else {
		buf, _ := ioutil.ReadAll(resp.Body)
		t.Log("OK", string(buf))
	}
}
