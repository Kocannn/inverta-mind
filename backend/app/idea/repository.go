package idea

import (
	"context"

	"github.com/Kocannn/self-dunking-ai/domain"
	pkgDB "github.com/Kocannn/self-dunking-ai/pkg/database"
	"github.com/sirupsen/logrus"
)

type (
	repository struct {
		db pkgDB.DatabaseTransaction
	}
)

// GetIdea implements domain.IdeaRepository.
func (r *repository) GetIdea(ctx context.Context, id int) (domain.SubmitIdeaRequest, error) {
	data := domain.SubmitIdeaRequest{}
	err := r.db.DB(ctx).First(&data, id).Error
	if err != nil {
		return domain.SubmitIdeaRequest{}, err
	}
	return data, nil
}

// SubmitIdeaStream implements domain.IdeaRepository.
func (r *repository) SubmitIdeaStream(ctx context.Context, idea domain.SubmitIdeaRequest) (domain.SubmitIdeaRequest, error) {
	err := r.db.DB(ctx).Create(&idea).Error
	if err != nil {
		logrus.Error("repository.SubmitIdeaStream: failed to save idea")
		return domain.SubmitIdeaRequest{}, err
	}
	return idea, nil
}

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
