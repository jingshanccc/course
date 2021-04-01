package dao

import (
	"gitee.com/jingshanccc/course/file/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/util"
)

type FileShard struct {
	Id    int32
	Key   string
	Index int32
	Size  int32
	Total int32
}

func (FileShard) TableName() string {
	return "file_shard"
}

type FileShardDao struct {
}

//GetUploadedShards: 获取传入文件的所有已上传分片索引
func (f *FileShardDao) GetUploadedShards(key string) []int32 {
	var res []int32
	public.DB.Raw("select `index` from file_shard where `key` = ?", key).Find(&res)
	return res
}

//Save: 保存文件分片记录
func (f *FileShardDao) Save(in *dto.FileShardDto) *public.BusinessException {
	var fileShard FileShard
	_ = util.CopyProperties(&fileShard, in)
	err := public.DB.Model(&FileShard{}).Create(&fileShard).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//DeleteByKey: 上传完成后删除分片记录
func (f *FileShardDao) DeleteByKey(key string) {
	public.DB.Delete(FileShard{}, "`key` = ?", key)
}
