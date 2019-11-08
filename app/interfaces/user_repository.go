package interfaces

import (
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
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	us := make(entities.Users, 0)

	for rows.Next() {
		u := &entities.User{}

		if err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt); err != nil {
			return nil, err
		}

		us = append(us, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return us, nil
}

func (ur *UserRepository) FindByID(id int64) (*entities.User, error) {
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
		WHERE
			id = ?
	`

	rows, err := ur.SQLHandler.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	u := &entities.User{}

	if err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt); err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return u, nil
}
