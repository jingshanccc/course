package config

const (
	DefaultPassword = "123456"
	//192.168.10.130 127.0.0.1
	MySQLUrl          = "root:123456@tcp(192.168.10.130:33060)/course?charset=utf8mb4&parseTime=True&loc=Local"
	RedisUrl          = "192.168.10.130:6379"
	RegistryAddr      = "192.168.10.130:8500"
	UserServiceName   = "com.chan.course.user"
	UserClientName    = "com.chan.course.user-cli"
	GatewayName       = "com.chan.course.gateway"
	CourseServiceName = "com.chan.course.course"
	CourseCliName     = "com.chan.course.course-cli"
)

// file-service
const (
	FileServiceName = "com.chan.course.file"
	FileClientName  = "com.chan.course.file-cli"
	FilePath        = "F:/GoWorkspace/course/file-srv/store/"
	FileUrl         = "http://dev.course.com:4000/api/v1/file/store/"
)

const (
	JwtKey      = "micah"
	TokenExpire = 48
	UserInfoKey = "user_info_"
)

const (
	EmailResetEmailCode = "email_reset_email_code_"
)

const (
	TemplatePath = "F:/GoWorkspace/course/public/template"
)
