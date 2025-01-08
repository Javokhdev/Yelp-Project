package usecase

import (
	"github.com/Javokhdev/Yelp-Project/config"
	"github.com/Javokhdev/Yelp-Project/internal/usecase/repositories"
	"github.com/Javokhdev/Yelp-Project/pkg/logger"
	"github.com/Javokhdev/Yelp-Project/pkg/postgres"
)

// UseCase -.
type UseCase struct {
	UserRepo    UserRepoI
	SessionRepo SessionRepoI
	BusinessRepo BusinessI
	ReviewRepo ReviewI
	CategoryRepo CategoryRepoI
}

// New -.
func New(pg *postgres.Postgres, config *config.Config, logger *logger.Logger) *UseCase {
	return &UseCase{
		UserRepo:    repositories.NewUserRepo(pg, config, logger),
		SessionRepo: repositories.NewSessionRepo(pg, config, logger),
		BusinessRepo: repositories.NewBusinessRepo(pg, config, logger),
		ReviewRepo: repositories.NewReviewRepo(pg, config, logger),
		CategoryRepo: repositories.NewCategoryRepo(pg, config, logger),
	}
}
