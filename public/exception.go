package public

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

func NewBusinessException(enum ExceptionEnum) BusinessException {
	return BusinessException{enum.getCode(), enum.getDesc()}
}
func NoException(desc string) BusinessException {
	return BusinessException{int32(OK), desc}
}
