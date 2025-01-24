package post

import (
	"github.com/gin-gonic/gin"
)

func (h *Post) updatebyid(c *gin.Context) {
	id := c.Param("id")
	f := datatransfer.PostCreate{}
	err := c.ShouldBindJSON(&f)
	post, err := h.postService.UpdateOne({"id": id}, {"title": "title"})
}
