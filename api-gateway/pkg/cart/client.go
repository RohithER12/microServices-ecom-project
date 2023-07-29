package cart

import (
	"fmt"

	"github.com/RohithER12/api-gateway/pkg/cart/pb"
	"github.com/RohithER12/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CartServiceClient
}

func InitServiceClient(c *config.Config) pb.CartServiceClient {
	// using WithInsecure() because no SSL running
	fmt.Println("-------66666666666666667777777777777777777777-------\n\n\n", c.CartSvcUrl)

	cc, err := grpc.Dial(c.CartSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCartServiceClient(cc)
}
