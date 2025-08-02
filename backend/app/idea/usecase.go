package idea

import (
	"context"
	"github.com/Kocannn/self-dunking-ai/domain"
	pkgDB "github.com/Kocannn/self-dunking-ai/pkg/database"
)

type (
	usecase struct {
		dbTx pkgDB.DatabaseTransaction
		repo domain.IdeaRepository
	}
)

// SubmitIdea implements domain.IdeaUsecase.
func (u *usecase) SubmitIdea(ctx context.Context, idea string) error {
	panic("unimplemented")
}

var (
	uc *usecase
)

func NewIdeaUsecase(dbTx pkgDB.DatabaseTransaction, repo domain.IdeaRepository) domain.IdeaUsecase {
	if uc == nil {
		uc = &usecase{
			dbTx: dbTx,
			repo: repo,
		}
	}
	return uc
}
