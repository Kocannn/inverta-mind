package app

import (
	"github.com/Kocannn/self-dunking-ai/config"
	"github.com/Kocannn/self-dunking-ai/domain"
	"gorm.io/driver/postgres"

	"github.com/Kocannn/self-dunking-ai/app/idea"

	pkgDB "github.com/Kocannn/self-dunking-ai/pkg/database"
)

type App struct {
	IdeaHandler domain.IdeaHandler
}

func InitApp(cfg config.Config) App {
	db := config.GetDatabase(postgres.Dialector{
		Config: &postgres.Config{
			DSN: cfg.DB_POSTGRES_DSN,
		}})

	dbTx := pkgDB.NewDBTransaction(db)
	ideaRepo := idea.InitIdeaRepository(dbTx)

	ideaUsecase := idea.InitIdeaUsecase(dbTx, ideaRepo)

	ideaHandler := idea.InitIdeaHandler(ideaUsecase)

	return App{
		IdeaHandler: ideaHandler,
	}
}
