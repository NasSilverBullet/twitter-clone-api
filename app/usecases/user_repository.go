package usecases

import "github.com/NasSilverBullet/twitter-clone-api/app/entities"

type UserRepository interface {
	FindAll() (entities.Users, error)
	FindByID(id int64) (*entities.User, error)
}
