package dao

import (
	"course/course-srv/proto/dto"
	"course/public"
	"course/public/util"
	"log"
)

type TeacherDao struct {
}
type Teacher struct {
	Id       string
	Name     string
	Nickname string
	Image    string
	Position string
	Motto    string
	Intro    string
}

func (Teacher) TableName() string {
	return "teacher"
}

//All: 获取全部讲师
func (c *TeacherDao) All() ([]*dto.TeacherDto, *public.BusinessException) {
	var res []*dto.TeacherDto
	err := public.DB.Model(&Teacher{}).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

//List : get Teacher page
func (c *TeacherDao) List(in *dto.TeacherPageDto) (int64, []*dto.TeacherDto, *public.BusinessException) {
	forCount, forPage := util.GeneratePageSql(nil, in.Blurry, in.Sort, []string{"name", "nickname"}, "")
	var count int64
	err := public.DB.Model(&Course{}).Raw("select count(1) from teacher x " + forCount).Find(&count).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res []*dto.TeacherDto
	err = public.DB.Model(&Teacher{}).Raw("select x.* from teacher x "+forPage, (in.Page-1)*in.Size, in.Size).Find(&res).Error
	return count, res, nil
}

//Save: 保存/更新讲师
func (c *TeacherDao) Save(cd *dto.TeacherDto) (*dto.TeacherDto, *public.BusinessException) {
	TeacherEntity := &Teacher{}
	_ = util.CopyProperties(TeacherEntity, cd)
	if cd.Id != "" { //update
		TeacherEntity.Id = ""
		err := public.DB.Model(&Teacher{Id: cd.Id}).Updates(TeacherEntity).Error
		if err != nil {
			return &dto.TeacherDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		cd.Id = util.GetShortUuid()
		TeacherEntity.Id = cd.Id
		err := public.DB.Create(TeacherEntity).Error
		if err != nil {
			return &dto.TeacherDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return cd, nil
}

// Delete: 删除讲师
func (c *TeacherDao) Delete(ids []string) *public.BusinessException {
	err := public.DB.Model(&Teacher{}).Where("id in ?", ids).Delete(&Teacher{}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}
