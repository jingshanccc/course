package handler

import (
	"context"
	"course/config"
	"course/file-srv/dao"
	"course/file-srv/proto/dto"
	"course/proto/basic"
	"course/public/util"
	"encoding/base64"
	"github.com/micro/go-micro/v2/errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var fileDao = &dao.FileDao{}

type FileServiceHandler struct {
}

//Upload: 上传文件接口
func (f *FileServiceHandler) Upload(ctx context.Context, in *dto.FileDto, out *dto.FileDto) error {
	// 获取base64字符串 转为文件
	bytes, err := base64.StdEncoding.DecodeString(strings.Split(in.Shard, ",")[1])
	log.Println(err)
	path := in.Key + "." + in.Suffix
	localPath := config.FilePath + path + "." + strconv.FormatInt(int64(in.ShardIndex), 10)
	// 暂存文件到本地
	err = ioutil.WriteFile(localPath, bytes, 0666)
	if err != nil {
		return err
	}
	_ = util.CopyProperties(out, in)
	out.Path = path
	exception := fileDao.Save(out)
	if exception != nil {
		return errors.New(config.FileServiceName, exception.Error(), exception.Code())
	}
	// 当当前分片为最后一块分片时 合并本地文件 并删除分片
	if in.ShardIndex == in.ShardTotal {
		merge(in)
	}
	out.Path = config.FileUrl + path
	return nil
}

//Check: 检查文件是否已上传
func (f *FileServiceHandler) Check(ctx context.Context, in *basic.String, out *dto.FileDto) error {
	fileDb := fileDao.SelectByProperty("`key`", in.Str)
	if fileDb != nil {
		_ = util.CopyProperties(out, fileDb)
		out.Path = config.FileUrl + out.Path
	}
	return nil
}

//merge: 合并分片
func merge(in *dto.FileDto) {
	fileName := in.Key + "." + in.Suffix
	file, _ := os.Create(config.FilePath + fileName)
	defer file.Close()
	for i := 0; i < int(in.ShardTotal); i++ {
		shard, _ := os.Open(config.FilePath + fileName + "." + strconv.Itoa(i+1))
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
	for i := 0; i < int(in.ShardTotal); i++ {
		err := os.Remove(config.FilePath + fileName + "." + strconv.Itoa(i+1))
		if err != nil {
			log.Printf("delete shard failed, file is %s, shardIndex is %v \n", in.Key, i+1)
		}
	}
}
