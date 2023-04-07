package test

import (
	"app/api/models"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

func TestProduct(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			createProduct(t)
			// deleteProduct(t, id)
		}()

		s++
	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createProduct(t *testing.T) string {
	response := &models.Product{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreateProduct{
		Name:        faker.Name(),
		CategoryId:  "c9a98d0b-8007-4698-ae1d-301e4c06c773",
		Description: faker.Paragraph(),
		Price:       float64(rand.Intn(1000000-100) + 100),
		Quantity:    rand.Intn(10-1) + 1,
	}

	resp, err := PerformRequest(http.MethodPost, "/product", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func updateProduct(t *testing.T, id string) string {
	response := &models.Product{}
	request := &models.UpdateProduct{
		Name:        faker.Name(),
		CategoryId:  "c9a98d0b-8007-4698-ae1d-301e4c06c773",
		Description: faker.Paragraph(),
		Price:       float64(rand.Intn(1000000-100) + 100),
		Quantity:    rand.Intn(10-1) + 1,
	}

	resp, err := PerformRequest(http.MethodPut, "/product/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	return response.Id
}

func deleteProduct(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/product/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
