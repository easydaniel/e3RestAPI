package controllers

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/revel/revel"
)

type Api struct {
	*revel.Controller
}

type User struct {
	LoginTicket string
	AccountId   string
	DepartId    string
	Account     string
	Name        string
	Password    string
	OfficeTitle string
	Role        string
}

const (
	baseUrl  = "http://140.113.8.133"
	loginUrl = baseUrl + "/mService/service.asmx/Login"
)

func (c Api) Login() revel.Result {
	return c.Render()
}

func (c Api) Docs() revel.Result {
	return c.Render()
}

func (c Api) GetUserInfo() revel.Result {

	account := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	resp, err := http.Get(loginUrl + "?account=" + account + "&password=" + password)
	if err != nil {
		revel.WARN.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		revel.WARN.Println(err)
	}

	var user User
	xml.Unmarshal(body, &user)
	if user.AccountId != "" {
		user.OfficeTitle = strings.TrimSpace(user.OfficeTitle)
		if user.OfficeTitle == "學生" {
			user.Role = "stu"
		}
	}

	return c.RenderJson(user)
}
