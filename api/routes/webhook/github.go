package webhook

import (
	"context"
	"encoding/json"

	"github.com/google/go-github/github"
	"github.com/palantir/go-githubapp/githubapp"
)

type PushHandler struct {
	Client githubapp.ClientCreator
}

func (h *PushHandler) Handles() []string {
	return []string{"push"}
}

func (h *PushHandler) Handle(ctx context.Context, eventType, deliveryID string, payload []byte) error {
	// from github.com/google/go-github/github
	var event github.PushEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}
	println(*event.Pusher.Name)
	return nil
}
