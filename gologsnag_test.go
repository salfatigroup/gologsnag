package logsnag

import (
	"context"
	"os"
	"testing"
)

func TestLogSnag(t *testing.T) {
    logsnag := NewLogSnag(os.Getenv("LOGSNAG_TOKEN"), os.Getenv("LOGSNAG_PROJECT"))

    err := logsnag.Publish(
        context.Background(),
        &PublishOptions{
            Channel: "test",
            Event: "User Joined",
            Description: "Email: john@doe.com",
            Icon: "🎉",
            Tags: &Tags{
                "email": "john@doe.com",
                "user-id": "uid-12",
            },
            Notify: true,
        },
    )

    if err != nil {
        t.Error(err)
    }

    t.Log("LogSnag event published")
}
