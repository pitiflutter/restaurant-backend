package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func GetInvoices()gin.HandlerFunc{
	return func (c *gin.Context)  {
		ctx ,canecl := context.WithTimeout(context.Background(),100*time.Second)
		
	}
}
func GetInvoice()gin.HandlerFunc{
	return func (ctx *gin.Context)  {
		
	}
}
func CreateInvoice()gin.HandlerFunc{
	return func (ctx *gin.Context)  {
		
	}
}
func UpdateInvoice()gin.HandlerFunc{
	return func (ctx *gin.Context)  {
		
	}
}