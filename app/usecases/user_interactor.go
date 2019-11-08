package usecases

import "github.com/NasSilverBullet/twitter-clone-api/app/entities"

type UserInteractor struct {
	UserRepository
}

func (ui *UserInteractor) List() (entities.Users, error) {
	us, err := ui.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return us, nil
}
