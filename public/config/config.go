package config

// basic config
const (
	//192.168.10.130 127.0.0.1
	BasicName   = "com.chan.course.srv."
	BasicHost   = "192.168.10.130"
	ServiceHost = "127.0.0.1"
	MySQLUrl    = "root:123456@tcp(" + BasicHost + ":33060)/course?charset=utf8mb4&parseTime=True&loc=Local"
	RedisUrl    = BasicHost + ":6379"
	//RegistryAddr = "192.168.10.130:8500"
	RegistryAddr = BasicHost + ":2379"
	TracerAddr   = BasicHost + ":6831"
)

// user-service
const (
	DefaultPassword     = "123456"
	UserServiceName     = BasicName + "user"
	UserServiceAddr     = ServiceHost + ":8081"
	UserApiName         = BasicName + "api.user"
	UserApiAddr         = ServiceHost + ":8091"
	JwtKey              = "micah"
	TokenExpire         = 48
	UserInfoKey         = "user_info_"
	EmailResetEmailCode = "email_reset_email_code_"
	TemplatePath        = "F:/GoWorkspace/course/public/template"
)

// file-service
const (
	FileServiceName = BasicName + "file"
	FileServiceAddr = ServiceHost + ":8082"
	FileApiName     = BasicName + "api.file"
	FileApiAddr     = ServiceHost + ":8092"
	FilePath        = "F:/GoWorkspace/course/file-srv/store/"
	FileUrl         = "http://dev.course.com:4000/api/v1/file/store/"
)

// course-service
const (
	CourseServiceName = BasicName + "course"
	CourseServiceAddr = ServiceHost + ":8083"
	CourseApiName     = BasicName + "api.course"
	CourseApiAddr     = ServiceHost + ":8093"
)

// gateway
const (
	GatewayName = "com.chan.course.gateway"
	GatewayAddr = ServiceHost + ":4000"
)
