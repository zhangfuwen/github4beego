# github4beego
A ready made github controller for login with github

# 1. Usage Example

## 1.1 Installation
```go
go get github.com/bradrydzewski/go.auth
go get github.com/zhangfuwen/github4beego
```
## 1.2 Usage
```go
githubController = github4beego.NewGithubController("clientId","secretKey","/after_login")
beego.Router("/github_login", githubController, "get:FirstLogin")
beego.Router("/after_login", githubController,"get:SuccessRedirected")
```
`clientId` and `secretKey` is what you got when you apply a application at github.

When a user wants to login, you gave them a link first for her/him to click and login with github. The link could be `/github_login` or anything you like, as long as it's served with `githubController`'s FirstLogin method, the user will be redirected to github and finish the authentication process. When everything is finished, the user's User info will be retrieved from github and stored in user's session store and the user will then be redirected to `/after_login` or anything you like. You can then access user's User info. A example is given in githubController's SuccessRedirected method which prints out User to responseWriter. You can use it as a test. And then copy the code to serve your own page.

For your reference, SucessRedirected method is listed below:
```go
func (this *GithubHandler) SucessRedirected() {
	this.StartSession()
	user := this.GetSession("GitHubUser").(auth.GitHubUser)
	io.WriteString(this.Ctx.ResponseWriter, fmt.Sprintf("%v", user))

	return
}
```
And User is like:
```go
type GitHubUser struct {
	UserEmail    interface{} `json:"email"`
	UserName     interface{} `json:"name"`
	UserGravatar interface{} `json:"gravatar_id"`
	UserCompany  interface{} `json:"company"`
	UserLink     interface{} `json:"html_url"`
	UserLogin    string      `json:"login"`
}
```
