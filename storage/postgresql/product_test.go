package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateProduct
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateProduct{
				Name:        "Test Product",
				CategoryId:  "795e2770-fce8-4e24-ba90-0e695abdbd1d",
				Description: "Description testing",
				Price:       200,
				Quantity:    10,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := productTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ProductPrimaryKey
		Output  *models.Product
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ProductPrimaryKey{
				Id: "d606a347-8f95-4c75-9da7-8650d79e9824",
			},
			Output: &models.Product{
				Id:          "d606a347-8f95-4c75-9da7-8650d79e9824",
				Name:        "Product Test",
				CategoryId:  "795e2770-fce8-4e24-ba90-0e695abdbd1d",
				Description: "Description",
				Price:       200,
				Quantity:    10,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			product, err := productTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if product.Name != test.Output.Name || product.Id != test.Output.Id {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *product, *test.Output)
				return
			}

		})
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateProduct
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateProduct{
				Id:          "",
				Name:        "Product",
				CategoryId:  "",
				Description: "",
				Price:       200,
				Quantity:    10,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := productTestRepo.Update(context.Background(), test.Input)

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

func TestDeleteProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ProductPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ProductPrimaryKey{
				Id: "",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := productTestRepo.Delete(context.Background(), test.Input)

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
