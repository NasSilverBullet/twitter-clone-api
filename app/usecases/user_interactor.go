package usecases

type UserInteractor struct{}

func (ui *UserInteractor) Index() ([]struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}, error) {
	users := []struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
	}{
		{"Luke Skywalker", "male"},
		{"Leia Organa", "female"},
		{"Han Solo", "male"},
		{"Chewbacca", "male"},
	}

	return users, nil
}
