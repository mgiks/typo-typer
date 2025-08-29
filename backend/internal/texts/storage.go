package texts

import "context"

type RandomTextGetter interface {
	GetRandomText(ctx context.Context) (string, error)
}
