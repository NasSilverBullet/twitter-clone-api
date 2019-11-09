package interfaces_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/twitter-clone-api/app/entities"
	"github.com/NasSilverBullet/twitter-clone-api/app/interfaces"
)

type MockSQLHandler struct {
	interfaces.SQLHandler
}

var (
	allUsers = entities.Users{
		&entities.User{
			ID:    1,
			Name:  "1",
			Email: "1@example.com",
		},
		&entities.User{
			ID:    2,
			Name:  "2",
			Email: "2@example.com",
		},
		&entities.User{
			ID:    3,
			Name:  "3",
			Email: "3@example.com",
		},
	}
)

func (m *MockSQLHandler) Query(string, ...interface{}) (interfaces.Row, error) {
	return &MockRow{count: 0, users: allUsers}, nil
}

type MockRow struct {
	interfaces.Row
	count int
	users entities.Users
}

func (mRow *MockRow) Close() error { return nil }

func (mRow *MockRow) Next() bool {
	mRow.count++
	if mRow.count > 3 {
		return false
	}
	return true
}

func (mRow *MockRow) Scan(value ...interface{}) error {
	v0 := value[0].(*int64)
	*v0 = mRow.users[mRow.count-1].ID

	v1 := value[1].(*string)
	*v1 = mRow.users[mRow.count-1].Name

	v2 := value[2].(*string)
	*v2 = mRow.users[mRow.count-1].Email
	return nil
}

func (mRow *MockRow) Err() error { return nil }

func (m *MockSQLHandler) Begin() (interfaces.Tx, error) {
	return &MockTx{}, nil
}

type MockTx struct {
	interfaces.Tx
}

func (mTx *MockTx) Commit() error {
	return nil
}

func (mTx *MockTx) Rollback() error {
	return nil
}

func (mTx *MockTx) Exec(string, ...interface{}) (interfaces.Result, error) {
	return &MockResult{lastInsertID: 4}, nil
}

type MockResult struct {
	interfaces.Result
	lastInsertID int64
}

func (mResult *MockResult) LastInsertId() (int64, error) {
	return mResult.lastInsertID, nil
}

func TestUserRepository_FindAll(t *testing.T) {
	ur := &interfaces.UserRepository{&MockSQLHandler{}}
	got, err := ur.FindAll()
	if err != nil {
		t.Errorf("Unexpected Error: UserRepository.FindAll() >> %v", err)
	}

	want := entities.Users{
		&entities.User{
			ID:    1,
			Name:  "1",
			Email: "1@example.com",
		},
		&entities.User{
			ID:    2,
			Name:  "2",
			Email: "2@example.com",
		},
		&entities.User{
			ID:    3,
			Name:  "3",
			Email: "3@example.com",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UserRepository.FindAll() = %v, want %v", got, want)
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	ur := &interfaces.UserRepository{&MockSQLHandler{}}
	arg := int64(1)
	got, err := ur.FindByID(arg)
	if err != nil {
		t.Errorf("Unexpected Error: UserRepository.FindAll() >> %v", err)
	}

	want := &entities.User{
		ID:    1,
		Name:  "1",
		Email: "1@example.com",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UserRepository.FindAll() = %v, want %v", got, want)
	}
}

func TestUserRepository_Save(t *testing.T) {
	ur := &interfaces.UserRepository{&MockSQLHandler{}}
	arg := &entities.User{}
	got, err := ur.Save(arg)
	if err != nil {
		t.Errorf("Unexpected Error: UserRepository.FindAll() >> %v", err)
	}

	want := int64(4)

	if got != want {
		t.Errorf("UserRepository.Save(%v) = %v, want %v", arg, got, want)
	}
}
