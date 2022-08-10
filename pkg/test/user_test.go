package test

import (
	"errors"
	"reflect"
	"testing"
	"weight-tracker-api/pkg/api"
	"weight-tracker-api/pkg/api/request"
)

type mockUserRepo struct{}

func (m mockUserRepo) GetRole(RoleID int) (request.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockUserRepo) CreateUser(request request.NewUserRequest) error {
	if request.Name == "test user already created" {
		return errors.New("repository - user already exists in database")
	}

	return nil
}

func TestCreateNewUser(t *testing.T) {
	mockRepo := mockUserRepo{}
	mockUserService := api.NewUserService(&mockRepo)

	tests := []struct {
		name    string
		request request.NewUserRequest
		want    error
	}{
		{
			name: "should create a new user successfully",
			request: request.NewUserRequest{
				Name:     "test_user",
				Username: "test_user",
				Email:    "test_user@gmail.com",
				Sex:      "male",
				RoleID:   1,
			},
			want: nil,
		}, {
			name: "should return an error because of missing name",
			request: request.NewUserRequest{
				Name:     "",
				Username: "test_user",
				Email:    "test_user@gmail.com",
				Sex:      "male",
				RoleID:   1,
			},
			want: errors.New("user service - name required"),
		}, {
			name: "should return error from database because user already exists",
			request: request.NewUserRequest{
				Name:     "test user already created",
				Username: "test_user",
				Email:    "test_user@gmail.com",
				Sex:      "male",
				RoleID:   1,
			},
			want: errors.New("repository - user already exists in database"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockUserService.New(test.request)

			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}
