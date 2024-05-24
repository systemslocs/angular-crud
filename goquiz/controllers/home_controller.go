package controllers

import (
	"goquiz/models"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (hc HomeController) Index(c *gin.Context) {
	c.JSON(200, models.Mensagem{
		Mensagem: "GO.QUIZ",
	})
}
