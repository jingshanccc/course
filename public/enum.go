package public

type ExceptionEnum int32

const (
	//not business exception : dao,valid,etc.
	VALID_PARM_ERROR        ExceptionEnum = 10000
	PREPARE_SQL_ERROR       ExceptionEnum = 10001
	EXECUTE_SQL_ERROR       ExceptionEnum = 10002
	ROW_SCAN_ERROR          ExceptionEnum = 10003
	BEGIN_TRANSACTION_ERROR ExceptionEnum = 10004

	OK                       ExceptionEnum = 20000
	MEMBER_NOT_EXIST         ExceptionEnum = 20001
	USER_LOGIN_NAME_EXIST    ExceptionEnum = 20002
	USER_NOT_EXIST           ExceptionEnum = 20003
	LOGIN_MEMBER_ERROR       ExceptionEnum = 20004
	VERIFY_CODE_TOO_FREQUENT ExceptionEnum = 20005
	VERIFY_CODE_ERROR        ExceptionEnum = 20006
	VERIFY_CODE_EXPIRED      ExceptionEnum = 20007
	ERROR_PASSWORD           ExceptionEnum = 20008
	SAME_PASSWORD            ExceptionEnum = 20009
	SEND_EMAIL_CODE_ERROR    ExceptionEnum = 20010
	USER_EMAIL_EXIST         ExceptionEnum = 20011
	USER_PHONE_EXIST         ExceptionEnum = 20012
	ROLE_USER_EXIST          ExceptionEnum = 20013

	BAD_REQUEST    ExceptionEnum = 40000
	INTERVAL_ERROR ExceptionEnum = 50000
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
	case USER_NOT_EXIST:
		return "用户不存在"
	case LOGIN_MEMBER_ERROR:
		return "手机号不存在"
	case VERIFY_CODE_TOO_FREQUENT:
		return "获取验证码过于频繁"
	case VERIFY_CODE_ERROR:
		return "验证码不正确"
	case VERIFY_CODE_EXPIRED:
		return "验证码不存在或已过期，请重新发送"
	case ERROR_PASSWORD:
		return "密码错误"
	case SAME_PASSWORD:
		return "新密码不能与旧密码相同"
	case SEND_EMAIL_CODE_ERROR:
		return "发送邮件验证码失败，请检查邮箱有效性或联系管理员！"
	case USER_EMAIL_EXIST:
		return "用户邮箱已存在"
	case USER_PHONE_EXIST:
		return "用户手机号码已存在"
	case ROLE_USER_EXIST:
		return "所选角色存在用户关联，请解除关联再试！"
	case BAD_REQUEST:
		return "错误的客户端请求"
	case INTERVAL_ERROR:
		return "服务器内部错误"
	default:
		return "未知错误"
	}
}
