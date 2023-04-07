package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateClient(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateClient
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateClient{
				FirstName:   "Test Name",
				LastName:    "Test Last Name",
				PhoneNumber: "93-379-11-10",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := clientTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if id == "" {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}

func TestGetByIdClient(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ClientPrimaryKey
		Output  *models.Client
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ClientPrimaryKey{
				Id: "326108fa-b97d-4293-931f-d813245537f1",
			},
			Output: &models.Client{
				FirstName:   "Test Name Updated",
				LastName:    "Test Last Name",
				PhoneNumber: "93-379-11-10",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			client, err := clientTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if client.Id != test.Output.Id || client.FirstName != test.Output.FirstName || client.LastName != test.Output.LastName || client.PhoneNumber != test.Output.PhoneNumber {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *client, *test.Output)
				return
			}

		})
	}
}

func TestUpdateClient(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateClient
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateClient{
				Id:          "326108fa-b97d-4293-931f-d813245537f1",
				FirstName:   "Test Name Updated",
				LastName:    "Test Last Name",
				PhoneNumber: "93-379-11-10",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := clientTestRepo.Update(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}

func TestDeleteClient(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ClientPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ClientPrimaryKey{
				Id: "326108fa-b97d-4293-931f-d813245537f1",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := clientTestRepo.Delete(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}
