package public

type ExceptionEnum int32

const (
	//not business exception : dao,valid,etc.
	VALID_PARM_ERROR        ExceptionEnum = 10000
	PREPARE_SQL_ERROR       ExceptionEnum = 10001
	EXECUTE_SQL_ERROR       ExceptionEnum = 10002
	ROW_SCAN_ERROR          ExceptionEnum = 10003
	BEGIN_TRANSACTION_ERROR ExceptionEnum = 10004

	OK                      ExceptionEnum = 20000
	MEMBER_NOT_EXIST        ExceptionEnum = 20001
	USER_LOGIN_NAME_EXIST   ExceptionEnum = 20002
	LOGIN_USER_ERROR        ExceptionEnum = 20003
	LOGIN_MEMBER_ERROR      ExceptionEnum = 20004
	IMAGE_CODE_TOO_FREQUENT ExceptionEnum = 20005
	IMAGE_CODE_ERROR        ExceptionEnum = 20006
	IMAGE_CODE_EXPIRED      ExceptionEnum = 20007
	ERROR_PASSWORD          ExceptionEnum = 20008
	SAME_PASSWORD           ExceptionEnum = 20009
)

func (code ExceptionEnum) getCode() int32 {
	return int32(code)
}

func (code ExceptionEnum) getDesc() string {
	switch code {
	case OK:
		return "成功"
	case VALID_PARM_ERROR:
		return "请求参数不合法"
	case PREPARE_SQL_ERROR:
		return "准备sql语句失败"
	case EXECUTE_SQL_ERROR:
		return "执行sql语句失败"
	case ROW_SCAN_ERROR:
		return "查询结果赋值失败"
	case BEGIN_TRANSACTION_ERROR:
		return "开始数据库事务失败"
	case MEMBER_NOT_EXIST:
		return "会员不存在"
	case USER_LOGIN_NAME_EXIST:
		return "登录名已存在"
	case LOGIN_USER_ERROR:
		return "用户名不存在"
	case LOGIN_MEMBER_ERROR:
		return "手机号不存在"
	case IMAGE_CODE_TOO_FREQUENT:
		return "获取验证码过于频繁"
	case IMAGE_CODE_ERROR:
		return "验证码不正确"
	case IMAGE_CODE_EXPIRED:
		return "验证码不存在或已过期，请重新发送"
	case ERROR_PASSWORD:
		return "密码错误"
	case SAME_PASSWORD:
		return "新密码不能与旧密码相同"
	default:
		return "未知错误"
	}
}
