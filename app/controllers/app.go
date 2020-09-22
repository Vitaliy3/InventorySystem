package controllers

import (
	"encoding/base64"
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/entity"
	"myapp/app/providers"
	"net/http"
	"strconv"
	"strings"
)

var Session = make(map[string]string) //хранит все пользовательские сессиии

type App struct {
	*revel.Controller
	renderInterface app.RenderInterface
}

//функция проверки прав пользователя
func CheckPerm(c *revel.Controller, role string) bool {
	cookies, _ := c.Request.Cookie("token")
	if cookies == nil {
		return false
	}

	splitCookie := strings.Split(cookies.GetValue(), ":")
	session := Session[splitCookie[0]]
	if session != "" {
		splitSession := strings.Split(session, ":")
		if splitSession[1] == role {
			return true
		}
	}
	return false
}

//отображает главную страницу или отправляет пользователя на авторизацию
func (c App) Index() revel.Result {
	cookies, _ := c.Request.Cookie("token")
	if cookies == nil {
		return c.Redirect("/login")
	}

	splitCookie := strings.Split(cookies.GetValue(), ":")
	if Session[splitCookie[0]] != "" {
		return c.Render()
	} else {
		return c.Redirect("/login")
	}
}

//отображение страницы с авторизацией
func (c App) Login() revel.Result {
	return c.Render()
}

//проверяет данные для авторизации и создает куки для пользователя,если авторизация прошла успешно
func (c App) Auth() revel.Result {
	var login, password string
	employeeModel := providers.Employee{}

	cookies, _ := c.Request.Cookie("auth")
	splitCookie := strings.Split(cookies.GetValue(), ":")
	decoded, _ := base64.StdEncoding.DecodeString(splitCookie[0])
	authData := strings.Split(string(decoded), ":") //получение логина и пароля
	login = authData[0]
	password = authData[1]

	result, err := employeeModel.Auth(app.Db, entity.Authorization{Login: login, Password: password})
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		cookie := &http.Cookie{
			Name:  "token",
			Value: result.Token + ":" + result.Role,
		}

		Session[result.Token] = strconv.Itoa(result.Id) + ":" + result.Role
		c.SetCookie(cookie)
		c.renderInterface.Data = "/"
	}
	return c.RenderJSON(c.renderInterface)
}

//выход из аккаунта
func (c App) Logout() revel.Result {
	cookies, _ := c.Request.Cookie("token") //получение токена
	splitCookie := strings.Split(cookies.GetValue(), ":")
	delete(Session, splitCookie[0])                       //удаление сесиии пользователя
	c.renderInterface.Data = "/"
	return c.RenderJSON(c.renderInterface)
}
