package dao

import (
	"course/file-srv/proto/dto"
	"course/public"
	"course/public/util"
	"time"
)

type FileDao struct {
}
type File struct {
	Id         string
	Key        string
	Path       string
	Name       string
	Suffix     string
	Size       int64
	ShardIndex int32
	ShardSize  int32
	ShardTotal int32
	CreateAt   time.Time
	UpdateAt   time.Time
}

func (File) TableName() string {
	return "file"
}

//SelectByProperty: 通过一个属性查询
func (f *FileDao) SelectByProperty(prop string, value interface{}) *File {
	var res File
	err := public.DB.Model(&File{}).Where(prop+" = ?", value).Find(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

//Save: 新增/更新文件记录
func (f *FileDao) Save(in *dto.FileDto) *public.BusinessException {
	fileEntity := f.SelectByProperty("`key`", in.Key)
	now := time.Now()
	if fileEntity != nil && fileEntity.Id != "" { // update
		fileEntity.UpdateAt = now
		fileEntity.ShardIndex += 1
		err := public.DB.Model(&File{Id: fileEntity.Id}).Updates(fileEntity).Error
		if err != nil {
			return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else {
		fileEntity = &File{}
		_ = util.CopyProperties(fileEntity, in)
		fileEntity.Id = util.GetShortUuid()
		fileEntity.CreateAt = now
		fileEntity.UpdateAt = now
		err := public.DB.Create(fileEntity).Error
		if err != nil {
			return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return nil
}

//Save: 新增文件记录
func (f *FileDao) SaveNew(in *dto.FileDto) *public.BusinessException {
	now := time.Now()
	fileEntity := &File{}
	_ = util.CopyProperties(fileEntity, in)
	fileEntity.Id = util.GetShortUuid()
	fileEntity.Path = in.Key + "." + in.Suffix
	fileEntity.CreateAt = now
	fileEntity.UpdateAt = now
	err := public.DB.Create(fileEntity).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//DeleteById: 删除文件记录
func (f *FileDao) DeleteByProperty(property, value string) {
	public.DB.Delete(File{}, property+" = ?", value)
}
