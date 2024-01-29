package port

import (
	"context"

	"github.com/mkaiho/go-ws-sample/entity"
)

type IDGenerator interface {
	Generate(ctx context.Context) (entity.ID, error)
}
