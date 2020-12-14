package route

import (
	"course/gateway/handler"
	"course/gateway/handler/course"
	"course/gateway/handler/user"
	"course/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.SaveServices(service), middleware.Cors())
	v1 := ginRouter.Group("/api/v1")
	auth := v1.Group("/oauth")
	{
		auth.GET("/captcha/image-code", handler.GetCaptcha)
		auth.GET("/authorize", handler.Authorize)
		auth.GET("/redirect", handler.Redirect)
		auth.POST("/token", handler.Token)
	}
	admin := v1.Group("/admin")
	admin.Use(middleware.JWT())
	{
		userGroup := admin.Group("/user")
		{
			userGroup.GET("/list", user.GetUserList)
			userGroup.POST("/info", user.UserInfo)
			userGroup.POST("/save-password", user.SavePassword)
			userGroup.POST("/save", user.Save)
			userGroup.POST("/delete", user.DeleteUser)
			userGroup.GET("/logout", user.Logout)
		}
		//admin.GET("/captcha/image-code", handler.GetCaptcha)
		resource := admin.Group("/resource")
		{
			resource.GET("/load-menus", user.LoadMenus)
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

		courseGroup := admin.Group("/course")
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
			teacher.POST("/save", course.SaveTeacher)
			teacher.DELETE("/delete", course.DeleteTeacher)
		}
	}

	return ginRouter
}
