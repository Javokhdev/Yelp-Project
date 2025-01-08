package repositories

import (
	"context"
	"time"

	"github.com/Javokhdev/Yelp-Project/config"
	"github.com/Javokhdev/Yelp-Project/internal/entity"
	"github.com/Javokhdev/Yelp-Project/pkg/logger"
	"github.com/Javokhdev/Yelp-Project/pkg/postgres"
	"github.com/google/uuid"
)

type CategoryRepo struct {
	pg     *postgres.Postgres
	config *config.Config
	logger *logger.Logger
}

// NewCategoryRepo creates a new CategoryRepo instance.
func NewCategoryRepo(pg *postgres.Postgres, config *config.Config, logger *logger.Logger) *CategoryRepo {
	return &CategoryRepo{
		pg:     pg,
		config: config,
		logger: logger,
	}
}

// CreateCategory inserts a new category into the database.
func (r *CategoryRepo) CreateCategory(ctx context.Context, category *entity.Category) (string, error) {
	category.ID = uuid.NewString()

	query, args, err := r.pg.Builder.Insert("categories").
		Columns("id, name").
		Values(category.ID, category.Name).
		ToSql()
	if err != nil {
		return "", err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return category.ID, nil
}

// GetCategoryByID retrieves a category by its ID.
func (r *CategoryRepo) GetCategoryByID(ctx context.Context, categoryID entity.Id) (*entity.Category, error) {
	var category entity.Category

	query, args, err := r.pg.Builder.
		Select("id, name").
		From("categories").
		Where("id = ?", categoryID).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.pg.Pool.QueryRow(ctx, query, args...).
		Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// GetAllCategories retrieves all categories from the database.
func (r *CategoryRepo) GetAllCategories(ctx context.Context) ([]*entity.Category, error) {
	query, args, err := r.pg.Builder.
		Select("id, name").
		From("categories").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pg.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		category := &entity.Category{}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// UpdateCategory updates an existing category.
func (r *CategoryRepo) UpdateCategory(ctx context.Context, category *entity.Category) error {
	query, args, err := r.pg.Builder.
		Update("categories").
		SetMap(map[string]interface{}{
			"name":       category.Name,
			"updated_at": time.Now(),
		}).
		Where("id = ?", category.ID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	return err
}

// DeleteCategory deletes a category by its ID.
func (r *CategoryRepo) DeleteCategory(ctx context.Context, categoryID entity.Id) error {
	query, args, err := r.pg.Builder.
		Delete("categories").
		Where("id = ?", categoryID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	return err
}
