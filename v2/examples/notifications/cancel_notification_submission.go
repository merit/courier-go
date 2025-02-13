package examples

import (
	"context"

	"github.com/merit/courier-go/v2"
)

func cancelNotificationSubmission() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.CancelNotificationSubmission(context.Background(), "my-notification", "my-submission")
}
