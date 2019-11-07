package interfaces

import (
	"time"

	"github.com/NasSilverBullet/twitter-clone-api/app/entities"
)

type UserRepository struct {
	SQLHandler
}

func (ur *UserRepository) FindAll() (entities.Users, error) {
	const query = `
		SELECT
			id,
			name,
			email,
			created_at,
			updated_at,
			deleted_at
		FROM
			users
	`

	rows, err := ur.SQLHandler.Query(query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	us := make(entities.Users, 0)

	for rows.Next() {
		var (
			id        int64
			name      string
			email     string
			createdAt time.Time
			updatedAt time.Time
			deletedAt time.Time
		)

		if err = rows.Scan(&id, &name, &email, &createdAt, &updatedAt, &deletedAt); err != nil {
			return nil, err
		}

		u := &entities.User{
			ID:        id,
			Name:      name,
			Email:     email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}

		us = append(us, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return us, nil
}
