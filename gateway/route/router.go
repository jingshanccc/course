package route

import (
	"course/config"
	"course/gateway/handler"
	"course/gateway/handler/course"
	"course/gateway/handler/file"
	"course/gateway/handler/user"
	"course/gateway/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.SaveServices(service), middleware.Cors())
	v1 := ginRouter.Group("/api/v1")
	auth := v1.Group("/oauth")
	{
		auth.GET("/captcha/image-code", handler.LoginCaptcha)
		auth.GET("/authorize", handler.Authorize)
		auth.GET("/redirect", handler.Redirect)
		auth.POST("/token", handler.Token)
		auth.GET("/logout", user.Logout)
	}
	admin := v1.Group("/admin")
	admin.Use(middleware.JWT())
	{
		userGroup := admin.Group("/user")
		{
			userGroup.POST("/list", user.GetUserList)
			userGroup.POST("/info", user.UserInfo)
			userGroup.POST("/save-info", user.SaveUserInfo)
			userGroup.POST("/save-password", user.SavePassword)
			userGroup.POST("/save", user.Save)
			userGroup.POST("/delete", user.DeleteUser)
			userGroup.GET("/email-code", user.SendEmailCode)
			userGroup.POST("/update-email", user.UpdateEmail)
		}
		//admin.GET("/captcha/image-code", handler.LoginCaptcha)
		resource := admin.Group("/resource")
		{
			resource.POST("/list", user.MenuList)
			resource.GET("/load-menus", user.LoadMenus)
			resource.GET("/load-tree", user.LoadTree)
			resource.GET("/child", user.MenuChild)
			resource.POST("/parents", user.MenuParent)
			resource.POST("", user.SaveResource)
			resource.PUT("", user.SaveResource)
			resource.DELETE("", user.Delete)
		}
		role := admin.Group("/role")
		{
			role.GET("/all", user.AllRole)
			role.GET("/find", user.GetRole)
			role.GET("/level", user.RoleLevel)
			role.POST("/list", user.RoleList)
			role.POST("/save", user.SaveRole)
			role.DELETE("/delete", user.DeleteRole)
			role.POST("/save-resource", user.SaveRoleResource)
			role.GET("/list-resource", user.ListRoleResource)
			role.POST("/save-user", user.SaveRoleUser)
			role.GET("/list-user", user.ListRoleUser)
		}

		courseGroup := admin.Group("/course")
		{
			courseGroup.POST("/list", course.ListCourse)
			courseGroup.POST("", course.SaveCourse)
			courseGroup.DELETE("", course.DelCourse)
			courseGroup.PUT("", course.SaveCourse)
			courseGroup.POST("/save-content", course.SaveCourseContent)
			courseGroup.GET("/list-category", course.ListCourseCategory)
			courseGroup.POST("/sort", course.SortCourse)
			courseGroup.GET("/find-content", course.FindCourseContent)
		}
		category := admin.Group("/category")
		{
			category.POST("/list", course.ListCategory)
			category.GET("/all", course.AllCategory)
			category.POST("", course.SaveCategory)
			category.PUT("", course.SaveCategory)
			category.DELETE("", course.DeleteCategory)
		}
		chapter := admin.Group("/chapter")
		{
			chapter.POST("/list", course.ListChapter)
			chapter.GET("/all", course.AllChapter)
			chapter.POST("", course.SaveChapter)
			chapter.PUT("", course.SaveChapter)
			chapter.DELETE("", course.DelChapter)
		}
		section := admin.Group("/section")
		{
			section.POST("/list", course.ListSection)
			section.POST("", course.SaveSection)
			section.PUT("", course.SaveSection)
			section.DELETE("", course.DelSection)
		}
		courseFile := admin.Group("/course-file")
		{
			courseFile.GET("/list", course.ListCourseFile)
			courseFile.POST("/save", course.SaveCourseFile)
			courseFile.DELETE("/delete", course.DelCourseFile)
		}
		teacher := admin.Group("/teacher")
		{
			teacher.GET("/all", course.AllTeacher)
			teacher.POST("/list", course.ListTeacher)
			teacher.POST("", course.SaveTeacher)
			teacher.PUT("", course.SaveTeacher)
			teacher.DELETE("", course.DeleteTeacher)
		}
	}
	files := v1.Group("/file")
	files.StaticFS("/store", http.Dir(config.FilePath))
	files.Use(middleware.JWT())
	{
		files.POST("/upload", file.Upload)
		files.POST("/upload_shard", file.UploadShard)
		files.POST("/merge", file.Merge)
		files.GET("/verify_upload", file.VerifyUpload)
		files.GET("/cancel", file.Cancel)
		files.GET("/check", file.Check)
	}
	return ginRouter
}
