package examples

import (
	"context"
	"log"

	"github.com/merit/courier-go/v2"
)

func getBrands() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	data, err := client.GetBrands(context.Background(),
		"", // optional cursor
	)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
