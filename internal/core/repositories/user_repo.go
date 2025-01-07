package repositories

import (
	"database/sql"
	domain "yalp/internal/domain"

	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *domain.User) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO users (id, name, email, password_hash, profile_picture, bio)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	var userID string
	err := r.db.QueryRow(query, id, user.Name, user.Email, user.Password, user.ProfilePic, user.Bio).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (r *UserRepo) GetUserByID(userID string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRow("SELECT id, name, email, password_hash, profile_picture, bio FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.ProfilePic, &user.Bio)	
	if err != nil {
		return nil, err
	}
	return &user, nil		
}

func (r *UserRepo) GetAllUsers() ([]*domain.User, error) {
	query := `SELECT id, name, email, profile_picture, bio FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.ProfilePic, &user.Bio)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

func (r *UserRepo) UpdateUser(user *domain.User) error {
	_, err := r.db.Exec(`
		UPDATE users
		SET name = $1, email = $2, password_hash = $3, profile_picture = $4, bio = $5
		WHERE id = $6
	`, user.Name, user.Email, user.Password, user.ProfilePic, user.Bio, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteUser(userID string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}