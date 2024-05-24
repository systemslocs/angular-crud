package controllers

import (
	"goquiz/servicos"

	"github.com/gin-gonic/gin"
)

type PalavrasController struct{}

func (pc PalavrasController) Index(c *gin.Context) {
	servico := servicos.IAServico{}
	c.JSON(200, servico.BuscaPalavras())
}
