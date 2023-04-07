package postgresql

import (
	"app/api/models"
	"context"
	"math/rand"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateOrder
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateOrder{
				ClientId: "eeb13e6e-2312-43e6-a926-dc7b0ac6ff45",
				Price:    float64(rand.Intn(1000000-100) + 100),
				Status:   "new",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := orderTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdOrder(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.OrderPrimaryKey
		Output  *models.Order
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.OrderPrimaryKey{
				Id: "83d30858-c9e2-49cc-8fa5-23e49a72a793",
			},
			Output: &models.Order{
				Id: "83d30858-c9e2-49cc-8fa5-23e49a72a793",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			order, err := orderTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if order.Id != test.Output.Id {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *order, *test.Output)
				return
			}

		})
	}
}

func TestUpdateOrder(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateOrder
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateOrder{
				Id:       "83d30858-c9e2-49cc-8fa5-23e49a72a793",
				ClientId: "eeb13e6e-2312-43e6-a926-dc7b0ac6ff45",
				Price:    200000,
				Status:   "in_proccess",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := orderTestRepo.Update(context.Background(), test.Input)

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

func TestDeleteOrder(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.OrderPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.OrderPrimaryKey{
				Id: "83d30858-c9e2-49cc-8fa5-23e49a72a793",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := orderTestRepo.Delete(context.Background(), test.Input)

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

func TestCreateOrderProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateOrderItem
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateOrderItem{
				OrderId:   "05102f47-8dbe-4c80-b8db-0a00d0ad2c28",
				ProductId: "62d5cb0b-9798-4eeb-8fb0-156734306e68",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := orderTestRepo.AddOrderProduct(context.Background(), test.Input)

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

func TestDeleteOrderProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.OrderProductPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.OrderProductPrimaryKey{
				Id: "dd781ac6-8146-4450-8e9e-8622bdf5dbfd",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := orderTestRepo.RemoveOrderItem(context.Background(), test.Input)

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
