package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Javokhdev/Yelp-Project/config"
	"github.com/Javokhdev/Yelp-Project/internal/entity"
	"github.com/Javokhdev/Yelp-Project/pkg/logger"
	"github.com/Javokhdev/Yelp-Project/pkg/postgres"
	"github.com/google/uuid"
)

type BusinessRepo struct {
	pg     *postgres.Postgres
	config *config.Config
	logger *logger.Logger
}

// NewBusinessRepo creates a new BusinessRepo instance.
func NewBusinessRepo(pg *postgres.Postgres, config *config.Config, logger *logger.Logger) *BusinessRepo {
	return &BusinessRepo{
		pg:     pg,
		config: config,
		logger: logger,
	}
}

// CreateBusiness inserts a new business into the database.
func (r *BusinessRepo) CreateBusiness(ctx context.Context, business *entity.Business) (string, error) {
	business.ID = uuid.NewString()

	photosJSON, err := json.Marshal(business.Photos)
	if err != nil {
		return "", err
	}

	hoursJSON, err := json.Marshal(business.Hours)
	if err != nil {
		return "", err
	}

	query, args, err := r.pg.Builder.Insert("businesses").
		Columns("id, name, description, category, address, contact_info, photos, hours").
		Values(business.ID, business.Name, business.Description, business.Category, business.Address, business.ContactInfo, photosJSON, hoursJSON).
		ToSql()
	if err != nil {
		return "", err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return business.ID, nil
}

// GetBusinessByID retrieves a business by its ID.
func (r *BusinessRepo) GetBusinessByID(ctx context.Context, businessID string) (*entity.Business, error) {
	var (
		business  entity.Business
		photosJSON, hoursJSON []byte
	)

	query, args, err := r.pg.Builder.
		Select("id, name, description, address, category, contact_info, photos, hours").
		From("businesses").
		Where("id = ?", businessID).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.pg.Pool.QueryRow(ctx, query, args...).
		Scan(&business.ID, &business.Name, &business.Description, &business.Address, &business.Category, &business.ContactInfo, &photosJSON, &hoursJSON)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(photosJSON, &business.Photos); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(hoursJSON, &business.Hours); err != nil {
		return nil, err
	}

	return &business, nil
}

// GetAllBusinesses retrieves all businesses from the database.
func (r *BusinessRepo) GetAllBusinesses(ctx context.Context) ([]*entity.Business, error) {
	query, args, err := r.pg.Builder.
		Select("id, name, description, address, category, contact_info, photos, hours").
		From("businesses").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pg.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var businesses []*entity.Business
	for rows.Next() {
		var (
			business  entity.Business
			photosJSON, hoursJSON []byte
		)

		err := rows.Scan(&business.ID, &business.Name, &business.Description, &business.Address, &business.Category, &business.ContactInfo, &photosJSON, &hoursJSON)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(photosJSON, &business.Photos); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(hoursJSON, &business.Hours); err != nil {
			return nil, err
		}

		businesses = append(businesses, &business)
	}

	return businesses, nil
}

// UpdateBusiness updates an existing business.
func (r *BusinessRepo) UpdateBusiness(ctx context.Context, business *entity.Business) error {
	photosJSON, err := json.Marshal(business.Photos)
	if err != nil {
		return err
	}

	hoursJSON, err := json.Marshal(business.Hours)
	if err != nil {
		return err
	}

	query, args, err := r.pg.Builder.
		Update("businesses").
		SetMap(map[string]interface{}{
			"name":        business.Name,
			"description": business.Description,
			"address":     business.Address,
			"category":    business.Category,
			"contact_info": business.ContactInfo,
			"photos":      photosJSON,
			"hours":       hoursJSON,
			"updated_at":  time.Now(),
		}).
		Where("id = ?", business.ID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	return err
}

// DeleteBusiness deletes a business by its ID.
func (r *BusinessRepo) DeleteBusiness(ctx context.Context, businessID string) error {
	query, args, err := r.pg.Builder.Delete("businesses").Where("id = ?", businessID).ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	return err
}
