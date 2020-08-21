package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"

	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"
)

// UserController 用户管理控制器
type UserController struct {
	auth.LayoutController
}

// Add 添加用户
func (c *UserController) Add() {
	form := &forms.UserForm{}
	if c.Ctx.Input.IsPost() {
		c.ParseForm(form)
		if err := services.UserService.Add(form); err != nil {
			beego.Error(err.Error())
		} else {
			beego.Informational("添加新用户成功")
			// action := beego.AppConfig.DefaultString("auth::UserAction", "UserController.Query")
			action := "UserController.Query"
			c.Redirect(beego.URLFor(action), http.StatusFound)
		}
	}
	c.Data["user"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "user/add.html"
}

// Query 查询用户
func (c *UserController) Query() {
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data)

	q := c.GetString("q")

	c.Data["users"] = services.UserService.Query(q)
	c.Data["q"] = q
	c.TplName = "user/query.html"
}

// Modify 修改用户
func (c *UserController) Modify() {
	// 假设当前用户不能修改其他人的信息

	form := &forms.UserForm{}
	// GET 获取数据
	// POST 修改用户
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			//验证数据
			services.UserService.Modify(form)

			//存储消息
			flash := beego.NewFlash()
			flash.Set("notice", "修改用户信息成功")
			flash.Error("error")
			flash.Success("success")
			flash.Warning("warning")
			flash.Store(&c.Controller)

			c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}
	} else if pk, err := c.GetInt("pk"); err == nil {
		if user := services.UserService.GetByPk(pk); user != nil {
			form.ID = user.ID
			form.Name = user.Name
			form.Nickname = user.Nickname
			form.Department = user.Department
			form.Gender = user.Gender
			form.Tel = user.Tel
			form.Email = user.Email
			form.Addr = user.Addr
			form.Status = user.Status
		}
	}

	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "user/modify.html"
}

// Delete user data
func (c *UserController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil && c.LoginUser.ID != pk {
		services.UserService.Delete(pk)
	}
	c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}
