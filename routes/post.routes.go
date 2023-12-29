package routes

import (
	"restfulAPI/controllers"

	"github.com/gin-gonic/gin"
)

// Creating Routes for the CRUD Operations

type PostRouteController struct {
	PostController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{
		PostController: postController,
	}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("/posts")
	router.POST("/", pc.PostController.CreatePost)
	router.GET("/", pc.PostController.FindPosts)
	router.PUT("/:postId", pc.PostController.UpdatePost)
	router.GET("/:postId", pc.PostController.FindPostById)
	router.DELETE("/:postId", pc.PostController.DeletePost)

}
