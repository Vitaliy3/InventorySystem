package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"myapp/app/routes"
	"net/http"
	"strings"
)

type App struct {
	*revel.Controller
}

var Session = make(map[string]string)

func (c App) InvSys() revel.Result {
	fmt.Println("in invSys")
	return c.Render()
}

func (c App) Index() revel.Result {
	fmt.Println("------------------------------------------------------------------------------------------")
	cookies, _ := c.Request.Cookie("token")
	if cookies == nil {
		return c.Redirect("/login")
	}
	splitCookie := strings.Split(cookies.GetValue(), ":")
	if Session[splitCookie[0]] != "" {
		fmt.Println("success authorize")
		return c.Redirect(routes.App.InvSys())
	} else {
		fmt.Println("fail authorize")
		return c.Redirect("/login")
	}
	return c.Render()
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) TryAuth() revel.Result{
	var login, password, base string
	err := json.Unmarshal(c.Params.JSON, &base)
	if err != nil {
		fmt.Println("err in marshall", err)
	}
	decoded, err := base64.StdEncoding.DecodeString(base)
	authData := strings.Split(string(decoded), ":")
	login = authData[0]
	password = authData[1]
	if login == "1" && password == "1" {
		newToken := GravatarMD5(login)
		Session[newToken] = "admin"
		cookie := &http.Cookie{
			Name:  "token",
			Value: newToken + ":admin",
		}
		c.SetCookie(cookie)
		fmt.Println("Success login")
		return c.Redirect("/")
	}else{
		return c.RenderJSON("")
	}
	return c.Render()
}

func (c App) Logout() revel.Result{
	for k, _ := range Session {
		delete(Session, k)
	}
	return c.RenderJSON("")
}
func GravatarMD5(login string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(login)))
	return hex.EncodeToString(h.Sum(nil))
}
