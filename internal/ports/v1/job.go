package ports

import "context"

type Job interface {
	Name() string
	Schedule() string
	Execute(ctx context.Context) error
}
