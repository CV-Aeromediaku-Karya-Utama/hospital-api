package test

import (
	"errors"
	"inventory-api/pkg/api"
	"inventory-api/pkg/api/request"
	"reflect"
	"testing"
)

type mockRoleRepo struct{}

func (m mockRoleRepo) UpdateRole(RoleID int) (request.UpdateRoleRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockRoleRepo) DeleteRole(RoleID int) (request.DeleteRoleRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockRoleRepo) ListRole() error {
	//TODO implement me
	panic("implement me")
}

func (m mockRoleRepo) CreateRole(request request.NewRoleRequest) error {
	if request.Name == "test" {
		return errors.New("repository - role already exists in database")
	}

	return nil
}

func TestCreateNewRole(t *testing.T) {
	mockRepo := mockRoleRepo{}
	mockRoleService := api.NewRoleService(&mockRepo)

	tests := []struct {
		name    string
		request request.NewRoleRequest
		want    error
	}{
		{
			name: "should create a new role successfully",
			request: request.NewRoleRequest{
				Name: "admin",
			},
			want: nil,
		}, {
			name: "should return an error because of missing name",
			request: request.NewRoleRequest{
				Name: "",
			},
			want: errors.New("role service - name required"),
		}, {
			name: "should return error from database because role already exists",
			request: request.NewRoleRequest{
				Name: "test role already created",
			},
			want: errors.New("repository - role already exists in database"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockRoleService.New(test.request)

			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}
