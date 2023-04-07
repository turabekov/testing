package test

import (
	"app/api/models"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/test-go/testify/assert"
)

// func TestOrder(t *testing.T) {
// 	s = 0
// 	wg := &sync.WaitGroup{}

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			orderId := createOrder(t)

// 			wg2 := &sync.WaitGroup{}
// 			for i := 0; i < 10; i++ {
// 				wg2.Add(1)
// 				go func() {
// 					defer wg2.Done()
// 					createOrderProduct(t, orderId)
// 					// deleteOrderProduct(t, id)
// 				}()
// 			}
// 			wg2.Wait()

// 			// deleteOrder(t, orderId)
// 		}()

// 		s++
// 	}

// 	wg.Wait()

// 	fmt.Println("s: ", s)
// }

func TestOrder(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			orderId := createOrder(t)

			for i := 0; i < 10; i++ {
				id := createOrderProduct(t, orderId)

				deleteOrderProduct(t, id)
			}
			deleteOrder(t, orderId)
		}()

		s++
	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createOrder(t *testing.T) string {
	response := &models.Order{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreateOrder{
		ClientId: "eeb13e6e-2312-43e6-a926-dc7b0ac6ff45",
		Price:    float64(rand.Intn(1000000-100) + 100),
		Status:   "new",
	}
	resp, err := PerformRequest(http.MethodPost, "/order", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func updateOrder(t *testing.T, id string) string {
	response := &models.Order{}
	request := &models.UpdateOrder{
		ClientId: "eeb13e6e-2312-43e6-a926-dc7b0ac6ff45",
		Price:    float64(rand.Intn(1000000-100) + 100),
		Status:   "in_proccess",
	}

	resp, err := PerformRequest(http.MethodPut, "/order/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	return response.Id
}

func deleteOrder(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/order/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

func createOrderProduct(t *testing.T, orderId string) string {
	response := &models.OrderProductPrimaryKey{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreateOrderItem{
		OrderId:   orderId,
		ProductId: "30d0bcf6-460c-4213-8e28-8a3b4c7d2f68",
	}
	resp, err := PerformRequest(http.MethodPost, "/order_item", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	return response.Id
}

func deleteOrderProduct(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/order_item/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
