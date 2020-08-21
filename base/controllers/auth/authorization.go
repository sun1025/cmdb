package auth

import (
	"net/http"
	"strings"
	
	"github.com/astaxie/beego"
	
	"cmdb/base/controllers/base"
	"cmdb/models"
	"cmdb/services"

)

// AuthorizationController: All base controllers that require authentication to access
type AuthorizationController struct {
	base.BaseController
	LoginUser *models.User
}

func (c *AuthorizationController) getNav() string {
	controllerName, _ := c.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
}

// Prepare: User authentication check
func (c *AuthorizationController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	sessionValue := c.GetSession(sessionKey)
	
	c.Data["loginUser"] = nil
	c.Data["nav"] = c.getNav()
	
	if sessionValue != nil {
		if pk, ok := sessionValue.(int); ok {
			if user := services.UserService.GetByPk(pk); user != nil {
				c.Data["loginUser"] = user
				c.LoginUser = user
				return
			}
		}
	}
	
	action := beego.AppConfig.DefaultString("auth::LoginAction",
	"AuthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}