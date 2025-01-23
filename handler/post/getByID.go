package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Post) getbyid(c *gin.Context) {
	id := c.Param("id")
	post, err := h.postService.FindOne(map[string]interface{}{"id": id})
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": post.Id, "title": post.Title, "content": post.Content})
}
