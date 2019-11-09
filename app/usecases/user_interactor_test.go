package usecases_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/twitter-clone-api/app/entities"
	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
)

type MockUserRepository struct{}

func (m MockUserRepository) FindAll() (entities.Users, error) {
	return entities.Users{&entities.User{}, &entities.User{}, &entities.User{}}, nil
}

func TestUserInteractor_List(t *testing.T) {
	ui := &usecases.UserInteractor{&MockUserRepository{}}
	got, err := ui.List()
	if err != nil {
		t.Errorf("Unexpected Error: UserInteractor.List() >> %v", err)
	}

	want := entities.Users{&entities.User{}, &entities.User{}, &entities.User{}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UserInteractor.List() = %v, want %v", got, want)
	}
}

func (m MockUserRepository) FindByID(id int64) (*entities.User, error) {
	return &entities.User{}, nil
}

func TestUserInteractor_Get(t *testing.T) {
	ui := &usecases.UserInteractor{&MockUserRepository{}}
	arg := int64(1)
	got, err := ui.Get(arg)
	if err != nil {
		t.Errorf("Unexpected Error: UserInteractor.Get() >> %v", err)
	}

	want := &entities.User{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UserInteractor.Get() = %v, want %v", got, want)
	}
}

func (m MockUserRepository) Save(u *entities.User) (int64, error) {
	return 1, nil
}

func TestUserInteractor_Create(t *testing.T) {
	ui := &usecases.UserInteractor{&MockUserRepository{}}
	arg := &entities.User{}
	got, err := ui.Create(arg)
	if err != nil {
		t.Errorf("Unexpected Error: UserInteractor.Create() >> %v", err)
	}

	want := int64(1)

	if got != want {
		t.Errorf("UserInteractor.Create() = %v, want %v", got, want)
	}
}
