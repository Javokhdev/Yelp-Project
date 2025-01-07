package repositories

import (
	"database/sql"
	domain "yalp/internal/domain"

	"github.com/google/uuid"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) CreateCategory(category *domain.Category) (string, error) {
	id := uuid.New().String()
	_, err := r.db.Exec("INSERT INTO categories (id, name) VALUES ($1, $2)", id, category.Name)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *CategoryRepo) GetCategoryByID(categoryID string) (*domain.Category, error) {
	var category domain.Category
	err := r.db.QueryRow("SELECT id, name FROM categories WHERE id = $1", categoryID).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepo) GetAllCategories() ([]*domain.Category, error) {	
	query := "SELECT id, name FROM categories"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var categories []*domain.Category
	for rows.Next() {
		category := &domain.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryRepo) UpdateCategory(category *domain.Category) error {
	_, err := r.db.Exec("UPDATE categories SET name = $1 WHERE id = $2", category.Name, category.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepo) DeleteCategory(categoryID string) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id = $1", categoryID)
	if err != nil {
		return err
	}
	return nil
}
