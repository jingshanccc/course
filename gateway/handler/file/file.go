package file

import (
	"context"
	"course/config"
	"course/file-srv/proto/dto"
	"course/file-srv/proto/file"
	"course/proto/basic"
	"course/public"
	"github.com/gin-gonic/gin"
)

//Upload: 上传文件接口
func Upload(ctx *gin.Context) {
	var req dto.FileDto
	if err := ctx.Bind(&req); err == nil {
		fileService := ctx.Keys[config.FileServiceName].(file.FileService)
		res, err := fileService.Upload(context.Background(), &req)
		public.ResponseAny(ctx, err, res)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Check: 预检文件接口
func Check(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		fileService := ctx.Keys[config.FileServiceName].(file.FileService)
		res, err := fileService.Check(context.Background(), &req)
		public.ResponseAny(ctx, err, res)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}
