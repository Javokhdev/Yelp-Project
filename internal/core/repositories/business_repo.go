package repositories

import (
	"database/sql"
	"encoding/json"
	domain "github.com/Javokhdev/Yelp-Project/internal/domain"

	"github.com/google/uuid"
)

type BusinessRepo struct {
	db *sql.DB
}

func NewBusinessRepo(db *sql.DB) *BusinessRepo {
	return &BusinessRepo{
		db: db,
	}
}

func (r *BusinessRepo) CreateBusiness(business *domain.Business) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO businesses (id, name, description, category, address, contact_info, photos, hours)
		VALUES ($1, $2, $3, $4, $5, $6, $7::jsonb, $8::jsonb)
		RETURNING id
	`
	
	// Marshal Photos and Hours into JSON
	photosJSON, err := json.Marshal(business.Photos)
	if err != nil {
		return "", err
	}

	hoursJSON, err := json.Marshal(business.Hours)
	if err != nil {
		return "", err
	}

	var businessID string
	err = r.db.QueryRow(query, id, business.Name, business.Description, business.Category, business.Address, business.ContactInfo, photosJSON, hoursJSON).Scan(&businessID)
	if err != nil {
		return "", err
	}
	return businessID, nil
}

func (r *BusinessRepo) GetBusinessByID(businessID string) (*domain.Business, error) {
	var business domain.Business
	var photosJSON []byte
	var hoursJSON []byte

	query := `SELECT id, name, description, address, category, contact_info, photos, hours FROM businesses WHERE id = $1`
	err := r.db.QueryRow(query, businessID).Scan(&business.ID, &business.Name, &business.Description, &business.Address, &business.Category, &business.ContactInfo, &photosJSON, &hoursJSON)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSONB data into appropriate Go types
	if err := json.Unmarshal(photosJSON, &business.Photos); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(hoursJSON, &business.Hours); err != nil {
		return nil, err
	}

	return &business, nil
}

func (r *BusinessRepo) GetAllBusinesses() ([]*domain.Business, error) {
	query := `SELECT id, name, description, address, category, contact_info, photos, hours FROM businesses`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var businesses []*domain.Business
	for rows.Next() {
		business := &domain.Business{}
		var photosJSON []byte
		var hoursJSON []byte

		err := rows.Scan(&business.ID, &business.Name, &business.Description, &business.Address, &business.Category, &business.ContactInfo, &photosJSON, &hoursJSON)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSONB data into appropriate Go types
		if err := json.Unmarshal(photosJSON, &business.Photos); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(hoursJSON, &business.Hours); err != nil {
			return nil, err
		}

		businesses = append(businesses, business)
	}

	return businesses, nil
}

func (r *BusinessRepo) UpdateBusiness(business *domain.Business) error {
	// Marshal Photos and Hours into JSON
	photosJSON, err := json.Marshal(business.Photos)
	if err != nil {
		return err
	}

	hoursJSON, err := json.Marshal(business.Hours)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		UPDATE businesses
		SET name = $1, description = $2, address = $3, category = $4, contact_info = $5, photos = $6, hours = $7
		WHERE id = $8
	`, business.Name, business.Description, business.Address, business.Category, business.ContactInfo, photosJSON, hoursJSON, business.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *BusinessRepo) DeleteBusiness(businessID string) error {
	_, err := r.db.Exec("DELETE FROM businesses WHERE id = $1", businessID)
	if err != nil {
		return err
	}
	return nil
}
