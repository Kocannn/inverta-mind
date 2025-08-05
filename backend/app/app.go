package app

import (
	"github.com/Kocannn/self-dunking-ai/config"
	"github.com/Kocannn/self-dunking-ai/domain"
	"github.com/hammer-code/lms-be/pkg/jwt"
	"gorm.io/driver/postgres"

	"github.com/Kocannn/self-dunking-ai/app/idea"
	"github.com/Kocannn/self-dunking-ai/app/middleware"

	pkgDB "github.com/Kocannn/self-dunking-ai/pkg/database"
)

type App struct {
	IdeaHandler domain.IdeaHandler
	Middleware  domain.Middleware
}

func InitApp(cfg config.Config) App {
	db := config.GetDatabase(postgres.Dialector{
		Config: &postgres.Config{
			DSN: cfg.DB_POSTGRES_DSN,
		}})

	db.AutoMigrate(&domain.SubmitIdeaRequest{})

	jwtInstance := jwt.NewJwt(cfg.JWT_SECRET_KEY)

	middleware := middleware.InitMiddleware(jwtInstance)

	dbTx := pkgDB.NewDBTransaction(db)
	ideaRepo := idea.InitIdeaRepository(dbTx)

	ideaUsecase := idea.InitIdeaUsecase(dbTx, ideaRepo)

	ideaHandler := idea.InitIdeaHandler(ideaUsecase)

	return App{
		IdeaHandler: ideaHandler,
		Middleware:  middleware,
	}
}
