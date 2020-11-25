package route

import (
	"course/gateway/handler"
	"course/gateway/handler/course"
	"course/gateway/handler/user"
	"course/public"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(SaveServices(service))
	v1 := ginRouter.Group("/api/v1")

	admin := v1.Group("/admin")
	{
		userGroup := admin.Group("/user")
		{
			userGroup.GET("/list", user.GetUserList)
			userGroup.POST("/login", user.Login)
			userGroup.POST("/save-password", user.SavePassword)
			userGroup.POST("/save", user.Save)
			userGroup.POST("/delete", user.DeleteUser)
			userGroup.GET("/logout", user.Logout)
		}
		admin.GET("/captcha/image-code", handler.GetCaptcha)
		resource := admin.Group("/resource")
		{
			resource.GET("/load-tree", user.LoadTree)
			resource.POST("/save", user.SaveJson)
			resource.DELETE("/delete", user.Delete)
		}
		role := admin.Group("/role")
		{
			role.POST("/list", user.RoleList)
			role.POST("/save", user.SaveRole)
			role.DELETE("/delete", user.DeleteRole)
			role.POST("/save-resource", user.SaveRoleResource)
			role.GET("/list-resource", user.ListRoleResource)
			role.POST("/save-user", user.SaveRoleUser)
			role.GET("/list-user", user.ListRoleUser)
		}

		courseGroup := admin.Group("/courseGroup")
		{
			courseGroup.POST("/list", course.ListCourse)
			courseGroup.POST("/save", course.SaveCourse)
			courseGroup.DELETE("/delete", course.DelCourse)
			courseGroup.POST("/save-content", course.SaveCourseContent)
			courseGroup.GET("/list-category", course.ListCourseCategory)
			courseGroup.POST("/sort", course.SortCourse)
			courseGroup.GET("/find-content", course.FindCourseContent)
		}
		category := admin.Group("/category")
		{
			category.GET("/all", course.AllCategory)
			category.POST("/save", course.SaveCategory)
			category.DELETE("/delete", course.DeleteCategory)
		}
		chapter := admin.Group("/chapter")
		{
			chapter.POST("/list", course.ListChapter)
			chapter.POST("/save", course.SaveChapter)
			chapter.DELETE("/delete", course.DelChapter)
		}
		section := admin.Group("/section")
		{
			section.POST("/list", course.ListSection)
			section.POST("/save", course.SaveSection)
			section.DELETE("/delete", course.DelSection)
		}
	}

	return ginRouter
}

//SaveServices : 将服务实例存放到gin中
func SaveServices(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		//将实例存到gin.Keys里
		context.Keys = make(map[string]interface{})
		context.Keys[public.UserServiceName] = service[0]
		context.Keys[public.CourseServiceName] = service[1]
		context.Next()
	}
}

// Cors : 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Authorization"}
	config.AllowOrigins = []string{"*"} //http://localhost:8081
	config.AllowCredentials = true
	return cors.New(config)
}
