package gologsnag

import (
	"context"
	"os"
	"testing"
)

func TestLogSnagLog(t *testing.T) {
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

func TestLogSnagInsight(t *testing.T) {
    logsnag := NewLogSnag(os.Getenv("LOGSNAG_TOKEN"), os.Getenv("LOGSNAG_PROJECT"))

    err := logsnag.Insight(
        context.Background(),
        &InsightOptions{
            Title: "Test Insight",
            Value: 1,
            Icon: "🤭",
        },
    )

    if err != nil {
        t.Error(err)
    }

    t.Log("LogSnag insight published")
}
