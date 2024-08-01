package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("SLOTH_JOB_NOTIFIER_PROJECT"))
	if err != nil {
		log.Fatalf("failed to create pubsub client: %s", err)
	}

	topic := client.Topic(os.Getenv("SLOTH_JOB_NOTIFIER_TOPIC"))

	if len(os.Args) < 2 {
		log.Fatal("missing job id")
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("failed to parse job id '%s': %s", os.Args[1], err)
	}

	msg, err := json.Marshal(map[string]any{
		"id": id,
	})
	if err != nil {
		log.Fatalf("failed to marshal message: %s", err)
	}

	res := topic.Publish(ctx, &pubsub.Message{
		Attributes: map[string]string{
			"event":     "refresh",
			"namespace": os.Getenv("SLOTH_JOB_NOTIFIER_NAMESPACE"),
		},
		Data: msg,
	})

	_, err = res.Get(ctx)
	if err != nil {
		log.Fatalf("failed to publish message: %s", err)
	}

	topic.Stop()
}
