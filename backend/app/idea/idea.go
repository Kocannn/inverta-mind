package idea

import (
	"github.com/Kocannn/self-dunking-ai/domain"
	"github.com/hammer-code/lms-be/pkg/db"
)

func InitIdeaRepository(db db.DatabaseTransaction) domain.IdeaRepository {
	return NewIdeaRepository(db)
}
func InitIdeaUsecase(dbTx db.DatabaseTransaction, repo domain.IdeaRepository) domain.IdeaUsecase {
	return NewIdeaUsecase(dbTx, repo)
}
func InitIdeaHandler(usecase domain.IdeaUsecase) domain.IdeaHandler {
	return NewIdeaHandler(usecase)
}
