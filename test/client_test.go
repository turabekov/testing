package test

import (
	"app/api/models"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

func TestClient(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			createClient(t)
			// deleteClient(t, id)
		}()

		s++
	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createClient(t *testing.T) string {
	response := &models.Client{}

	request := &models.Client{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
	}

	resp, err := PerformRequest(http.MethodPost, "/client", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func updateClient(t *testing.T, id string) string {
	response := &models.Client{}
	request := &models.UpdateClient{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
	}

	resp, err := PerformRequest(http.MethodPut, "/client/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.Id
}

func deleteClient(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/client/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
