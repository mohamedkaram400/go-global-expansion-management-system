package ports

import "context"

type MatchNotification struct {
	MatchID   int
	ProjectID int
	VendorID  int
	Score     float64
	// Add recipient emails or project/vendor details if needed
	To      []string
	Subject string
	Body    string
}

type Notifier interface {
	SendMatchNotification(ctx context.Context, payload MatchNotification) error
}
