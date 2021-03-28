package config

// basic config
const (
	//192.168.10.130 127.0.0.1
	BasicName   = "com.chan.course.srv."
	BasicHost   = "192.168.10.130"
	ServiceHost = "127.0.0.1"
	MySQLUrl    = "root:123456@tcp(192.168.10.130:33060)/course?charset=utf8mb4&parseTime=True&loc=Local"
	RedisUrl    = "192.168.10.130:6379"
	//RegistryAddr = "192.168.10.130:8500"
	RegistryAddr = "192.168.10.130:2379"
	TracerAddr   = "192.168.10.130:6831"
)

// user-service
const (
	DefaultPassword     = "123456"
	UserServiceName     = BasicName + "user"
	UserClientName      = BasicName + "user-cli"
	UserServiceAddr     = ServiceHost + ":8081"
	JwtKey              = "micah"
	TokenExpire         = 48
	UserInfoKey         = "user_info_"
	EmailResetEmailCode = "email_reset_email_code_"
	TemplatePath        = "F:/GoWorkspace/course/public/template"
)

// file-service
const (
	FileServiceName = BasicName + "file"
	FileClientName  = "com.chan.course.file-cli"
	FileServiceAddr = ServiceHost + ":8082"
	FilePath        = "F:/GoWorkspace/course/file-srv/store/"
	FileUrl         = "http://dev.course.com:4000/api/v1/file/store/"
)

// course-service
const (
	CourseServiceName = BasicName + "course"
	CourseCliName     = BasicName + "course-cli"
	CourseServiceAddr = ServiceHost + ":8083"
)

// gateway
const (
	GatewayName = "com.chan.course.gateway"
	GatewayAddr = ServiceHost + ":4000"
)
