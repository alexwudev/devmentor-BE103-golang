package route

import (
	"devmentor-BE103-golang/handler/general"
	"devmentor-BE103-golang/handler/post"
	"github.com/gin-gonic/gin"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()

	group := server.Group("")

	post.NewPost(group)
	general.NewGeneral(group)

	return server
}
