package examples

import (
	"context"

	"github.com/merit/courier-go/v2"
)

func putList() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.PutList(context.Background(),
		"my-list",
		courier.PutListBody{
			Name: "Besties",
		},
	)
}
