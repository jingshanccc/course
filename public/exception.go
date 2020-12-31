package public

import "fmt"

type BusinessException struct {
	code        int32
	description string
}

func (c BusinessException) Error() string {
	return c.description
}
func (c BusinessException) Code() int32 {
	return c.code
}

func NewBusinessException(enum ExceptionEnum) *BusinessException {
	return &BusinessException{enum.getCode(), enum.getDesc()}
}
func NoException(desc string) *BusinessException {
	return &BusinessException{int32(OK), desc}
}

//BadRequestException: 错误的客户端请求，自定义错误信息
func BadRequestException(desc string) *BusinessException {
	if desc == "" {
		desc = BAD_REQUEST.getDesc()
	}
	return &BusinessException{int32(BAD_REQUEST), desc}
}

//EntityExistException: 存在属性相同的实体
func EntityExistException(entity, property string, value interface{}) *BusinessException {
	return &BusinessException{int32(BAD_REQUEST), fmt.Sprintf("已存在%v为%v的%v，请修改！", property, value, entity)}
}
