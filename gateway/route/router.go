package route

import (
	"gitee.com/jingshanccc/course/gateway/handler"
	"gitee.com/jingshanccc/course/gateway/handler/course"
	"gitee.com/jingshanccc/course/gateway/handler/file"
	"gitee.com/jingshanccc/course/gateway/handler/user"
	"gitee.com/jingshanccc/course/gateway/middleware"
	"gitee.com/jingshanccc/course/public/config"
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
	// 后台管理系统接口
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
			courseGroup.POST("/sort", course.SortCourse)
			courseGroup.GET("/find-content", course.FindCourseContent)
		}
		category := admin.Group("/category")
		{
			category.POST("/list", course.ListCategory)
			category.GET("/primary", course.PrimaryCategory)
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
			teacher.GET("/search", course.SearchTeacher)
			teacher.POST("/list", course.ListTeacher)
			teacher.POST("", course.SaveTeacher)
			teacher.PUT("", course.SaveTeacher)
			teacher.DELETE("", course.DeleteTeacher)
		}
	}
	// 文件服务接口
	files := v1.Group("/file")
	files.StaticFS("/store", http.Dir(config.Conf.Services["file"].Others["filePath"].(string)))
	files.Use(middleware.JWT())
	{
		files.POST("/upload", file.Upload)
		files.POST("/upload_shard", file.UploadShard)
		files.POST("/merge", file.Merge)
		files.GET("/verify_upload", file.VerifyUpload)
		files.GET("/cancel", file.Cancel)
		files.GET("/check", file.Check)
	}
	// 视频平台接口
	web := v1.Group("/web")
	{
		web.GET("/all-category", course.AllCategory)
		web.GET("/carousel-course", course.CarouselCourse)
		web.GET("/new-publish", course.NewPublishCourse)
		web.GET("/course-list", course.CategoryCourse)
		web.GET("/course", course.GetCourse)
		web.GET("/course-content", course.DownloadCourseContent)
		web.GET("/related-course", course.RelatedCourse)
		// 登陆注册
		web.GET("email-code", user.EmailRegisterCode)
		web.POST("register", user.MemberRegister)
		u := web.Group("/user")
		//u.Use(middleware.JWT())
		{
			// 文件
			u.POST("/upload_shard", file.UploadShard)
			u.POST("/merge", file.Merge)
			u.GET("/verify_upload", file.VerifyUpload)
			u.GET("/cancel", file.Cancel)
			u.GET("/check", file.Check)
			// 我的课程
			u.GET("courses", course.MyCourse)
			// 添加到我的课程
			u.POST("add-course", course.AddToMyCourse)
			// 获取课程学习进度
			u.POST("course-info", course.CourseInfo)
		}
	}
	return ginRouter
}
