package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"home/leonid/Git/Pract/network/pkg/models"
	"home/leonid/Git/Pract/network/pkg/service"
	"strings"

	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type transport struct {
	echo *echo.Echo
	port string
	svc  service.Service
}

func NewTransport(port string, svc service.Service) *transport {

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	return &transport{
		echo: e,
		port: port,
		svc:  svc,
	}
}

func (t *transport) Start() error {

	return t.echo.Start(":" + t.port)

}

func (t *transport) InitEndpoints() {
	t.echo.PUT("/change/:id", t.changeAge)
	t.echo.POST("/create/user", t.creatUser)
	t.echo.POST("/make/friends", t.makeFriends)
	t.echo.DELETE("/delete/user", t.deleteUser)
	t.echo.GET("/friends/:id", t.getFriends)
	t.echo.GET("/users", t.getUsers)

}

func (t *transport) creatUser(c echo.Context) error {

	body := c.Request().Body
	defer body.Close()

	buf, err := ioutil.ReadAll(body)
	if err != nil {
		err = errors.New("не могу записать в буфер, проверьте правильность ReadAll(body)")
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}

	u := &models.User{}

	err = json.Unmarshal(buf, u)
	if err != nil {
		err = errors.New("не могу сериализовать json, проверьте правильность структуры")
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}
	err = t.svc.CreateUser(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}

	return c.JSON(http.StatusCreated, CreateResponse(fmt.Sprintf("Пользователь  %v добавлен, Id: %v  ", u.Name, u.ID), http.StatusText(http.StatusCreated)))
}

func (t *transport) makeFriends(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	buf, err := ioutil.ReadAll(body)
	if err != nil {
		err = errors.New("не могу записать в буфер, проверьте правильность ReadAll(body)")
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}

	m := &models.MakeFfriends{}

	err = json.Unmarshal(buf, m)
	if err != nil {
		err = errors.New("не могу сериализовать json, проверьте правильность структуры")
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}

	err = t.svc.MakeFriends(m.SourceID, m.TargetID)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
		}

		return c.JSON(http.StatusInternalServerError, CreateError(err, http.StatusText(http.StatusInternalServerError)))
	}

	return c.JSON(http.StatusOK, CreateResponse(fmt.Sprintf("Пользователь %v и пользователь %v теперь друзья", m.SourceID, m.TargetID), http.StatusText(http.StatusOK)))

}

func (t *transport) deleteUser(c echo.Context) error {
	body := c.Request().Body

	buf, err := ioutil.ReadAll(body)
	if err != nil {
		err = errors.New("не могу записать в буфер, проверьте правильность структуры")
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}

	del := &models.DeleteUser{}

	err = json.Unmarshal(buf, del)
	if err != nil {
		err = errors.New("не могу сериализовать json, проверьте правильность структуры")
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}
	err = t.svc.DeleteUser(del.TargetID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}
	return c.JSON(http.StatusOK, CreateResponse(fmt.Sprintf("Пользователь c ID: %v удален", del.TargetID), http.StatusText(http.StatusOK)))

}

func (t *transport) getFriends(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}

	data, err := t.svc.GetFriends(ID)

	fmt.Println(data)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
		}

		return c.JSON(http.StatusInternalServerError, CreateError(err, http.StatusText(http.StatusInternalServerError)))
	}

	return c.JSON(http.StatusOK, CreateResponse(fmt.Sprintf("Друзья пользователя с ID %v: %v", id, data), http.StatusText(http.StatusOK)))
}

func (t *transport) getUsers(c echo.Context) error {
	data, err := t.svc.GetUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (t *transport) changeAge(c echo.Context) error {

	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return err

	}

	body := c.Request().Body

	buf, err := ioutil.ReadAll(body)
	if err != nil {
		err = errors.New("не могу записать в буфер, проверьте правильность структуры")
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}

	ch := &models.ChangeAge{}

	err = json.Unmarshal(buf, ch)
	if err != nil {
		err = errors.New("не могу сериализовать json, проверьте правильность структуры")
		return c.JSON(http.StatusBadRequest, CreateError(err, http.StatusText(http.StatusBadRequest)))
	}
	err = t.svc.ChangeAge(ID, ch.NewAge)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, CreateResponse(fmt.Sprintf("Возраст пользователя успешно обновлён на %v", ch.NewAge), http.StatusText(http.StatusOK)))
}
