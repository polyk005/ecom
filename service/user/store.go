package user

import (
	"database/sql"
	"fmt"

	"github.com/sikozonpc/ecom/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// GetUser ByEmail retrieves a user by email using GORM's First method
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User

	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUsers retrieves all users using a raw SQL query and scans them into a slice
func (s *Store) GetUsers() ([]*types.User, error) {
	var users []*types.User

	// Execute a raw SQL query
	rows, err := s.db.Raw("SELECT * FROM users").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows and scan them into User structs
	for rows.Next() {
		user, err := scanRowIntoUser (rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// scanRowIntoUser  scans a single row into a User struct
func scanRowIntoUser (rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
