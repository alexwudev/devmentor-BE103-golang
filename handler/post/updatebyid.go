package post

import (
	"devmentor-BE103-golang/model/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Post) updatebyid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error("Invalid ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := database.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		logrus.Error("JSON binding failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	post.Id = id
	posts, err := h.postService.UpdateOne(strconv.Itoa(id), &post)
	if err != nil {
		logrus.Error("Update error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}
