package post

import (
	"devmentor-BE103-golang/model/database"
	"devmentor-BE103-golang/model/datatransfer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Post) create(c *gin.Context) {

	f := datatransfer.PostCreate{}
	err := c.ShouldBindJSON(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := database.Post{
		Title:   f.Title,
		Content: f.Content,
	}
	err = h.postService.Create(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
