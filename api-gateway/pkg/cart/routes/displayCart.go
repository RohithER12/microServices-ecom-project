package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/RohithER12/api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

type DisplayCartRequestBody struct {
	UserId int64 `json:"userId"`
}

func DisplayCart(ctx *gin.Context, c pb.CartServiceClient) {

	userIdInterface, _ := ctx.Get("userId")
	userId, ok := userIdInterface.(int64)
	if !ok {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("userId is not of type int64"))
		return
	}
	res, err := c.DisplayCart(context.Background(), &pb.DisplayCartRequest{
		UserId: userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
