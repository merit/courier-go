package examples

import (
	"context"

	"github.com/merit/courier-go/v2"
)

func postBrand() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PostBrand(context.Background(), courier.PostBrandBody{
		Name: "my-brand",
	})
}
