package examples

import (
	"context"

	"github.com/merit/courier-go/v2"
)

func putBrand() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PutBrand(context.Background(), "my-brand", courier.PutBrandBody{
		Name: "my-updated-brand",
	})
}
