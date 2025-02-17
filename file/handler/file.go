package handler

import (
	"context"
	"encoding/base64"
	"gitee.com/jingshanccc/course/file/dao"
	"gitee.com/jingshanccc/course/file/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"github.com/micro/go-micro/v2/errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	fileDao      = &dao.FileDao{}
	fileShardDao = &dao.FileShardDao{}
)

type FileServiceHandler struct {
}

//Upload: 上传文件接口
func (f *FileServiceHandler) Upload(ctx context.Context, in *dto.FileDto, out *dto.FileDto) error {
	// 获取base64字符串 转为文件
	bytes, err := base64.StdEncoding.DecodeString(strings.Split(in.Shard, ",")[1])
	log.Println(err)
	path := in.Key + "." + in.Suffix
	localPath := config.Conf.Services["file"].Others["filePath"].(string) + path + "." + strconv.FormatInt(int64(in.ShardIndex), 10)
	// 暂存文件到本地
	err = ioutil.WriteFile(localPath, bytes, 0666)
	if err != nil {
		return err
	}
	_ = util.CopyProperties(out, in)
	out.Path = path
	exception := fileDao.Save(out)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["file"].Name, exception.Error(), exception.Code())
	}
	// 当当前分片为最后一块分片时 合并本地文件 并删除分片
	if in.ShardIndex == in.ShardTotal {
		merge(in)
	}
	out.Path = config.Conf.Services["file"].Others["fileUrl"].(string) + path
	return nil
}

//Check: 检查文件是否已上传
func (f *FileServiceHandler) Check(ctx context.Context, in *basic.String, out *dto.FileDto) error {
	fileDb := fileDao.SelectByProperty("`key`", in.Str)
	if fileDb != nil {
		_ = util.CopyProperties(out, fileDb)
		out.Path = config.Conf.Services["file"].Others["fileUrl"].(string) + out.Path
	}
	return nil
}

//merge: 合并分片
func merge(in *dto.FileDto) {
	fileName := in.Key + "." + in.Suffix
	file, _ := os.Create(config.Conf.Services["file"].Others["filePath"].(string) + fileName)
	defer file.Close()
	for i := 1; i <= int(in.ShardTotal); i++ {
		shard, _ := os.Open(config.Conf.Services["file"].Others["filePath"].(string) + fileName + "." + strconv.Itoa(i))
		buf := make([]byte, 10*1024*1024)
		for {
			n, err := shard.Read(buf)
			if err == io.EOF {
				shard.Close()
				break
			}
			//写入文件
			file.Write(buf[:n])
		}
	}
	time.Sleep(100 * time.Millisecond)
	for i := 1; i <= int(in.ShardTotal); i++ {
		err := os.Remove(config.Conf.Services["file"].Others["filePath"].(string) + fileName + "." + strconv.Itoa(i))
		if err != nil {
			log.Printf("delete shard failed, file is %s, shardIndex is %v \n", in.Key, i)
		}
	}
}

//VerifyUpload: 检查文件是否已上传 以及已上传的分片索引列表
func (f *FileServiceHandler) VerifyUpload(ctx context.Context, in *basic.String, out *dto.VerifyRes) error {
	// 首先从file表查询是否有该文件
	fileDb := fileDao.SelectByProperty("`key`", in.Str)
	if fileDb.Id != "" { // 若有 直接返回 shouldUpload=false
		out.ShouldUpload = false
		var fileDto dto.FileDto
		_ = util.CopyProperties(&fileDto, fileDb)
		fileDto.Path = config.Conf.Services["file"].Others["fileUrl"].(string) + fileDto.Path
		out.File = &fileDto
	} else { // 若没有 查询file_shard表该文件的所有记录的index字段
		uploadedShards := fileShardDao.GetUploadedShards(in.Str)
		out.ShouldUpload = true
		out.UploadedList = uploadedShards
	}
	return nil
}

//UploadShard: 上传文件分片
func (f *FileServiceHandler) UploadShard(ctx context.Context, in *dto.FileShardDto, out *basic.Boolean) error {
	localPath := config.Conf.Services["file"].Others["filePath"].(string) + in.Key + "." + strconv.FormatInt(int64(in.Index), 10)
	// 暂存文件到本地
	err := ioutil.WriteFile(localPath, in.Blob, 0666)
	if err != nil {
		exception := public.IntervalException("文件写出错误")
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["file"].Name, exception.Error(), exception.Code())
	}
	exception := fileShardDao.Save(in)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["file"].Name, exception.Error(), exception.Code())
	}
	if shards := fileShardDao.GetUploadedShards(in.Key); len(shards) == int(in.Total) {
		// 所有分片上传完成 返回的布尔值作为前端发起合并请求的标识
		out.Is = true
	} else {
		out.Is = false
	}
	return nil
}

//Merge: 合并文件
func (f *FileServiceHandler) Merge(ctx context.Context, in *dto.FileDto, out *dto.FileDto) error {
	file, _ := os.Create(config.Conf.Services["file"].Others["filePath"].(string) + in.Key + "." + in.Suffix)
	defer file.Close()
	for i := 0; i < int(in.ShardTotal); i++ {
		shard, _ := os.Open(config.Conf.Services["file"].Others["filePath"].(string) + in.Key + "." + strconv.Itoa(i))
		buf := make([]byte, 10*1024*1024)
		for {
			n, err := shard.Read(buf)
			if err == io.EOF {
				shard.Close()
				break
			}
			//写入文件
			_, err = file.Write(buf[:n])
			if err != nil {
				log.Printf("merge file failed, file is %s, shardIndex is %v \n", in.Key, i)
				exception := public.IntervalException("合并文件错误")
				return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["file"].Name, exception.Error(), exception.Code())
			}
		}
	}
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < int(in.ShardTotal); i++ {
		err := os.Remove(config.Conf.Services["file"].Others["filePath"].(string) + in.Key + "." + strconv.Itoa(i))
		if err != nil {
			log.Printf("delete shard failed, file is %s, shardIndex is %v \n", in.Key, i)
		}
	}
	fileShardDao.DeleteByKey(in.Key)
	exception := fileDao.SaveNew(in)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["file"].Name, exception.Error(), exception.Code())
	}
	out.Path = config.Conf.Services["file"].Others["fileUrl"].(string) + in.Key + "." + in.Suffix
	return nil
}

//Cancel: 取消上传文件/删除文件接口
func (f *FileServiceHandler) Cancel(ctx context.Context, in *basic.String, out *basic.String) error {
	// 先判断是否已完成上传
	fileDb := fileDao.SelectByProperty("`key`", in.Str)
	if fileDb.Id != "" { // 已完成 直接删除文件和对应数据库记录
		_ = os.Remove(config.Conf.Services["file"].Others["filePath"].(string) + fileDb.Key + "." + fileDb.Suffix)
		fileDao.DeleteByProperty("`key`", fileDb.Id)
	} else { // 未完成 删除已上传的文件分片和对应数据库记录
		shards := fileShardDao.GetUploadedShards(in.Str)
		for _, index := range shards {
			err := os.Remove(config.Conf.Services["file"].Others["filePath"].(string) + in.Str + "." + strconv.Itoa(int(index)))
			if err != nil {
				log.Printf("delete shard failed, file is %s, shardIndex is %v \n", in.Str, index)
			}
		}
		fileShardDao.DeleteByKey(in.Str)
	}
	return nil
}
