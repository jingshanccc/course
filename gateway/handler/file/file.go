package file

import (
	"context"
	"course/config"
	"course/file-srv/proto/dto"
	"course/file-srv/proto/file"
	"course/proto/basic"
	"course/public"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
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

//UploadShard: 上传分片
func UploadShard(ctx *gin.Context) {
	blob, err := ctx.FormFile("blob")
	if err != nil {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
		return
	}
	content, err := blob.Open()
	if err != nil {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
		return
	}
	bytes, err := ioutil.ReadAll(content)
	if err != nil {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
		return
	}
	var req dto.FileShardDto
	req.Blob = bytes
	req.Key = ctx.PostForm("key")
	i, _ := strconv.ParseInt(ctx.PostForm("index"), 10, 32)
	req.Index = int32(i)
	i, _ = strconv.ParseInt(ctx.PostForm("size"), 10, 32)
	req.Size = int32(i)
	i, _ = strconv.ParseInt(ctx.PostForm("total"), 10, 32)
	req.Total = int32(i)
	fileService := ctx.Keys[config.FileServiceName].(file.FileService)
	res, err := fileService.UploadShard(context.Background(), &req)
	public.ResponseAny(ctx, err, res)
}

//VerifyUpload: 预检文件接口
func VerifyUpload(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		fileService := ctx.Keys[config.FileServiceName].(file.FileService)
		res, err := fileService.VerifyUpload(context.Background(), &req)
		public.ResponseAny(ctx, err, res)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Merge: 合并文件接口
func Merge(ctx *gin.Context) {
	var req dto.FileDto
	if err := ctx.Bind(&req); err == nil {
		// 由于分片可能数量较大 设置较长的超时时间防止由于读写文件超时而无法正常响应
		con, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		fileService := ctx.Keys[config.FileServiceName].(file.FileService)
		res, err := fileService.Merge(con, &req)
		public.ResponseAny(ctx, err, res)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Cancel: 取消上传文件接口
func Cancel(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		fileService := ctx.Keys[config.FileServiceName].(file.FileService)
		res, err := fileService.Cancel(context.Background(), &req)
		public.ResponseAny(ctx, err, res)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}
