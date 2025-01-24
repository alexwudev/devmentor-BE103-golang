package post

import (
	"github.com/gin-gonic/gin"
)

func (h *Post) updatebyid(c *gin.Context) {
	id := c.Param("id")
	post, err := h.postService.UpdateOne(map[string]interface{}{"id": id}, map[string]interface{}{"title": id})
}
