package idea

import (
	"context"
	"github.com/Kocannn/self-dunking-ai/domain"
	pkgDB "github.com/Kocannn/self-dunking-ai/pkg/database"
)

type (
	repository struct {
		db pkgDB.DatabaseTransaction
	}
)

// SubmitIdea implements domain.IdeaRepository.
func (r *repository) SubmitIdea(ctx context.Context, idea string) error {
	panic("unimplemented")
}

var (
	repo *repository
)

func NewIdeaRepository(db pkgDB.DatabaseTransaction) domain.IdeaRepository {
	if repo == nil {
		repo = &repository{
			db,
		}

	}
	return repo
}
