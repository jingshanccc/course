package config

const (
	JwtKey = "micah"
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

const (
	Email_Reset_Email_Code = "email_reset_email_code_"
)
