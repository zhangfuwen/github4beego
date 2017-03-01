package github4beego

import (
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bradrydzewski/go.auth"
	"io"
	"net/http"
)

func NewGithubController(clientId, secretKey, successRedirect string) GithubController {
	gob.Register(auth.GitHubUser{})
	// Set the default authentication configuration parameters
	auth.Config.CookieSecret = []byte("_github_user")
	auth.Config.LoginRedirect = "/auth/login"          // send user here to login
	auth.Config.LoginSuccessRedirect = successRedirect // send user here post-login
	auth.Config.CookieSecure = false                   // for local-testing only
	return GithubController{
		authHandler: auth.Github(clientID, secretKey, "user,user:email"),
	}
}

type GithubController struct {
	beego.Controller
	authHandler *auth.AuthHandler
}

func (this *GithubController) FirstLogin() {
	this.StartSession()
	this.authHandler.Success = func(w http.ResponseWriter, r *http.Request, u auth.User, t auth.Token) {
		this.SetSession("GitHubUser", *(u.(*auth.GitHubUser)))
		http.Redirect(w, r, auth.Config.LoginSuccessRedirect, http.StatusSeeOther)
	}
	this.authHandler.Failure = func(w http.ResponseWriter, r *http.Request, err error) {
		beego.Error(err)
	}

	this.authHandler.ServeHTTP(this.Ctx.ResponseWriter, this.Ctx.Request)

}

func (this *GithubController) SucessRedirected() {
	this.StartSession()
	user := this.GetSession("GitHubUser").(auth.GitHubUser)
	io.WriteString(this.Ctx.ResponseWriter, fmt.Sprintf("%v", user))
	return
}
