package routes

import (
	"bookLine-Backend/api"
	"bookLine-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	OutAuth := router.Group("/api/")
	{
		OutAuth.GET("/captive", api.UserTest)                //测试通信
		OutAuth.POST("/user/login", api.UserLogin)           //login   登录系统
		OutAuth.POST("/user/register", api.UserRegister)     //register  注册账号
		OutAuth.GET("/novel/all", api.GetNovelPassed)        //获取全部审核通过的小说
		OutAuth.GET("/novel/category", api.GetNovelCategory) //获取分类小说
	}

	Auth := router.Group("/api/auth/")
	{
		Auth.Use(middleware.JWTAuth())
		//获取收藏小说列表
		//点赞小说

		//用户操作
		Auth.GET("/checkToken", api.CheckToken)    //检查token是否正确
		Auth.POST("/novel/upload", api.UserUpload) //上传小说
		Auth.GET("/novel/search", api.SearchNovel) //小说搜索

		Auth.POST("/novel/fav/:id", api.FavNovel)    //收藏小说
		Auth.DELETE("/novel/fav/:id", api.DeleteFav) //删除收藏
		Auth.GET("/novel/fav/", api.GetFav)          //获取收藏列表

		//管理员操作
		Auth.GET("/novel/review", api.GetNovelReview)     //待审
		Auth.GET("/novel/all", api.GetNovelAll)           //全部
		Auth.DELETE("/novel/:id", api.DeleteNovel)        //删除
		Auth.PATCH("/novel/:id/:status", api.UpdateNovel) //更新审核状态
	}

	return router
}
