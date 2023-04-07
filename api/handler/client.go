package handler

import (
	"app/api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Client godoc
// @ID create_customer
// @Router /client [POST]
// @Summary Create Client
// @Description Create Client
// @Tags Client
// @Accept json
// @Produce json
// @Param client body models.CreateClient true "CreateClientRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateClient(c *gin.Context) {

	var createCustomer models.CreateClient

	err := c.ShouldBindJSON(&createCustomer) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create customer", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Client().Create(context.Background(), &createCustomer)
	if err != nil {
		h.handlerResponse(c, "storage.customer.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Client().GetByID(context.Background(), &models.ClientPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Get By ID Client godoc
// @ID get_by_id_customer
// @Router /client/{id} [GET]
// @Summary Get By ID Client
// @Description Get By ID Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdClient(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storages.Client().GetByID(context.Background(), &models.ClientPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get customer by id", http.StatusCreated, resp)
}

// Get List Client godoc
// @ID get_list_customer
// @Router /client [GET]
// @Summary Get List Client
// @Description Get List Client
// @Tags Client
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListClient(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list customer", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list customer", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Client().GetList(context.Background(), &models.GetListClientRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list customer response", http.StatusOK, resp)
}

// Update Client godoc
// @ID update_customer
// @Router /client/{id} [PUT]
// @Summary Update Client
// @Description Update Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param client body models.UpdateClient true "UpdateClientRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateClient(c *gin.Context) {

	var updateCustomer models.UpdateClient

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateCustomer)
	if err != nil {
		h.handlerResponse(c, "update customer", http.StatusBadRequest, err.Error())
		return
	}

	updateCustomer.Id = id

	rowsAffected, err := h.storages.Client().Update(context.Background(), &updateCustomer)
	if err != nil {
		h.handlerResponse(c, "storage.customer.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.customer.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Client().GetByID(context.Background(), &models.ClientPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update customer", http.StatusAccepted, resp)
}

// DELETE Client godoc
// @ID delete_customer
// @Router /client/{id} [DELETE]
// @Summary Delete Client
// @Description Delete Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param client body models.ClientPrimaryKey true "DeleteClientRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteClient(c *gin.Context) {

	id := c.Param("id")

	rowsAffected, err := h.storages.Client().Delete(context.Background(), &models.ClientPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.customer.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.customer.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete customer", http.StatusNoContent, nil)
}
