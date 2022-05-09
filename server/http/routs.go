package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"preh/data"
	"preh/server/rabbit"

	"github.com/gin-gonic/gin"
)

type UserCreateReq struct {
	FirstName string
	LastName  string
}

func GetUserById(ctx *gin.Context) {

	id := ctx.Params.ByName("id")
	fmt.Println("passed here", id)
	if id == "" {
		fmt.Println("passed here")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "err.Error()",
		})
	}
	user, err := data.GetUserDetails(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "err.Error()",
		})

	}
	fmt.Println("Ddd", user)
	ctx.JSON(200, user)
	return
}
func CreateUser(ctx *gin.Context) {
	var userCreateReq data.UserCreate
	if err := ctx.ShouldBindJSON(&userCreateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := data.CreateUser(ctx, &userCreateReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(200, user)

}

func GetOrderById(ctx *gin.Context) {
	fmt.Println("passed here")

	id := ctx.Params.ByName("id")
	fmt.Println("passed here", id)
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "err.Error()",
		})
	}
	order, err := data.GetOrderDetails(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "err.Error()",
		})
	}
	ctx.JSON(200, order)

	return
}
func CreateOrder(ctx *gin.Context) {
	var orderCreateReq data.OrderCreate
	if err := ctx.ShouldBindJSON(&orderCreateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	order, err := data.CreateOrder(ctx, &orderCreateReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	pricePair := rabbit.PricePair{
		Price: order.Price,
		Pair:  order.Pair,
	}
	pricePairBytes, err := json.Marshal(&pricePair)
	if err != nil {
		fmt.Println("routes handle failed , TODO :: have to revert transaction")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err = rabbit.Publish("", "price", pricePairBytes)

	if err != nil {
		fmt.Println("routes handle failed , TODO :: have to revert transaction")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(200, order)

}

func UpdateOrderById(ctx *gin.Context) {
	var orderUpdate data.OrderUpdate
	if err := ctx.ShouldBindJSON(&orderUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := data.UpdateOrder(ctx, &orderUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"Message": "updated successfully",
	})
}
