package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

type AddCartRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func AddCart(ctx *gin.Context, c pb.CartServiceClient) {
	body := AddCartRequestBody{}

	userIdInterface, _ := ctx.Get("userId")
	userId, ok := userIdInterface.(int64)
	if !ok {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("userId is not of type int64"))
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println("====================================\n\n\n\n\n", body, "\n\n\n\n\n", userId)
	res, err := c.AddCart(context.Background(), &pb.AddCartRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId,
	})
	fmt.Println("\n\n\n\n\nerooooooooooooooooooooooooor\n\n\n\n\n", err, "\n\n\n\n\n", "-----------==========")
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
