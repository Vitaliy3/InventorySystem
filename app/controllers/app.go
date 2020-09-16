package controllers

import (
	"encoding/base64"
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/models"
	"net/http"
	"strconv"
	"strings"
)

var Session = make(map[string]string)

type App struct {
	*revel.Controller
}

func CheckPerm(c *revel.Controller, role string) bool {
	cookies, _ := c.Request.Cookie("token")
if cookies==nil{
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

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) Auth() revel.Result {
	renderInterface := app.RenderInterface{}
	employeeMod := models.Employee{}
	var login, password string
	cookies, _ := c.Request.Cookie("auth")
	splitCookie := strings.Split(cookies.GetValue(), ":")
	decoded, _ := base64.StdEncoding.DecodeString(splitCookie[0])
	authData := strings.Split(string(decoded), ":")
	login = authData[0]
	password = authData[1]
	result, err := employeeMod.Auth(app.DB, login, password)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		cookie := &http.Cookie{
			Name:  "token",
			Value: result.Token + ":" + result.Role,
		}
		Session[result.Token] = strconv.Itoa(result.Id) + ":" + result.Role
		c.SetCookie(cookie)
		renderInterface.Data = "/"
	}
	return c.RenderJSON(renderInterface)
}

func (c App) Logout() revel.Result {
	cookies, _ := c.Request.Cookie("token")
	splitCookie := strings.Split(cookies.GetValue(), ":")
	delete(Session, splitCookie[0])
	renderInterface := app.RenderInterface{}
	renderInterface.Data = "/"
	return c.RenderJSON(renderInterface)
}
