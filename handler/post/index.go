package post

import (
	"devmentor-BE103-golang/service"
	"github.com/gin-gonic/gin"
)

type Post struct {
	postService service.PostServiceInterface
}

func NewPost(
	r *gin.RouterGroup,

) *Post {
	h := &Post{
		postService: service.NewPostService(nil),
	}

	newRoute(h, r)

	return h
}

func newRoute(h *Post, r *gin.RouterGroup) {
	Group := r.Group("posts")

	Group.GET("", h.get)
	Group.POST("", h.create)
}
