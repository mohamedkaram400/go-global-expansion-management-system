package ports

import "context"

type MatchNotification struct {
    MatchID   uint
    ProjectID uint
    VendorID  uint
    Score     float64
    // Add recipient emails or project/vendor details if needed
    To []string
    Subject string
    Body    string
}

type Notifier interface {
    SendMatchNotification(ctx context.Context, payload MatchNotification) error
}