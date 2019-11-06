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
			id, name, email
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
			id    int
			name  string
			email string
		)

		if err = rows.Scan(&id, &name, &email); err != nil {
			return nil, err
		}

		u := &entities.User{
			ID:    id,
			Name:  name,
			Email: email,
		}

		us = append(us, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return us, nil
}
