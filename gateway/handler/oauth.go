package handler

import (
	"context"
	"course/gateway/middleware"
	"course/public"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Authorize: oauth获取授权码
func Authorize(ctx *gin.Context) {
	err := middleware.AuthServer.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}
}

//Token: oauth获取令牌
func Token(ctx *gin.Context) {
	err := middleware.AuthServer.HandleTokenRequest(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}
}

//Redirect: oauth跳转地址 完成将授权码和client信息重定向到获取令牌接口的工作
func Redirect(ctx *gin.Context) {
	client, _ := middleware.AuthServer.Manager.GetClient(context.Background(), ctx.Request.FormValue("state"))

	postValue := url.Values{
		"code":          {ctx.Request.FormValue("code")},
		"client_id":     {client.GetID()},
		"client_secret": {client.GetSecret()},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {"http://" + ctx.Request.Host + strings.Split(ctx.Request.RequestURI, "?")[0]},
	}

	resp, err := http.PostForm("http://"+ctx.Request.Host+"/api/v1/oauth/token", postValue)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}
	content := string(body)
	public.ResponseSuccess(ctx, content)
}
