package mocks

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
)

type MockNotifier struct {
	SendMatchNotificationFunc func(ctx context.Context, payload ports.MatchNotification) error
}

func (m *MockNotifier) SendMatchNotification(ctx context.Context, payload ports.MatchNotification) error {
	return m.SendMatchNotificationFunc(ctx, payload)
}
