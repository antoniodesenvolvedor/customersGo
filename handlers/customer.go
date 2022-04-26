package handlers

import (
	"customer-crud/schemas"
	"customer-crud/services"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)


// CreateCustomer godoc

// @tags customer
// @param name body schemas.Customer true "nome"
// @Summary Show a account
// @Description create customer
// @Accept  json
// @Produce  json
// @Success 200 {array} schemas.Customer "meu corpo de respota"
// @failure 404 {string} erro "Resposta de erro"
// @Router /customers/ [post]
func CreateCustomer(c *gin.Context) {

	var NewCustomerData schemas.Customer

	if err := c.BindJSON(&NewCustomerData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	newCustomer, err := services.CreateCustomer(&NewCustomerData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError	, err)
		return
	}


	c.IndentedJSON(http.StatusCreated, newCustomer)

}

func GetCustomers(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	limit, foundLimit := queryParams["limit"]
	if !foundLimit {
		c.IndentedJSON(http.StatusBadRequest, "Query parameter limit não informado")
		return
	}

	offset, foundOffset := queryParams["offset"]
	if !foundOffset{
		c.IndentedJSON(http.StatusBadRequest, "Query parameter offset não informado")
		return
	}

	parsedLimit, errLimit := strconv.Atoi(limit[0])
	parsedOffset, errOffset :=  strconv.Atoi(offset[0])
	if errLimit != nil || errOffset != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("Query parameters inválidos "))
		return
	}

	customer, err := services.GetCustomers(parsedLimit, parsedOffset)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, "ID não encontrado")
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, customer)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")

	customer, err := services.GetCustomer(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, "ID não encontrado")
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")

	var NewCustomerData schemas.Customer

	if err := c.BindJSON(&NewCustomerData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := services.UpdateCustomer(&NewCustomerData, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, "ID não encontrado")
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "dados atualizados com sucesso")
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteCustomer(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, "ID não encontrado")
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, fmt.Sprintf("recurso id %s deletado com sucesso", id))

}