package examples

import (
	"context"

	"github.com/merit/courier-go/v2"
)

func runJob() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.RunJob(context.Background(), "job123")
}
