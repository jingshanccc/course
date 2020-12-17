package handler

import (
	"course/public"
	"github.com/gin-gonic/gin"
)

type Captcha struct {
	Id           string
	Type         string
	VerifyValue  string
	Base64String string
}

//LoginCaptcha : 图形验证码
func LoginCaptcha(ctx *gin.Context) {
	var req Captcha
	if err := ctx.Bind(&req); err == nil {
		id, base64 := public.GetCaptchaBase64(req.Id, req.Type)
		req.Id = id
		req.Base64String = base64
		public.ResponseSuccess(ctx, req)
		//captcha.WriteTo(ctx.Writer)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}
