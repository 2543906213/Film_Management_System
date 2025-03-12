package routes

import (
	"backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

// 注册路由
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// 操作照片卡的接口
		api.GET("/photos", controllers.GetPhotos)
		api.GET("/photos/:id", controllers.GetPhotoByID)
		api.POST("/photos", controllers.CreatePhoto)
		api.PUT("/photos/:id", controllers.UpdatePhoto)
		api.DELETE("/photos/:id", controllers.DeletePhoto)

		// 操作评论的接口
		api.GET("/comments", controllers.GetComments)
		api.POST("/comments", controllers.CreateComment)
		api.DELETE("/comments/:id", controllers.DeleteComment)
	}

}
