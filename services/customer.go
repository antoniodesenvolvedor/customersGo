package services

import (
	"customer-crud/database"
	"customer-crud/models"
	"customer-crud/schemas"
	"customer-crud/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func CreateCustomer(newCustomerData *schemas.Customer) (*models.Customer, error) {

	birthDate, err := utils.ConvertDateStringToISODateFormat(newCustomerData.BirthDate)

	if err != nil {
		return nil, err
	}

	parsedBirthDate, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return nil, err
	}

	newCustomer := models.Customer{
		ID: uuid.New(),
		Name: newCustomerData.Name,
		CPF: newCustomerData.CPF,
		BirthDate: parsedBirthDate,
	}

	session := database.GetSession()

	result := session.Create(&newCustomer); if result.Error != nil {
		return nil, result.Error
	}

	return &newCustomer, nil
}

func GetCustomers(limit int, offset int)(*[]models.Customer, error) {
	session := database.GetSession()

	var customer []models.Customer
	result := session.Limit(limit).Offset(offset).Order("id").Find(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func GetCustomer(id string) (*models.Customer, error) {
	session := database.GetSession()
	var customer models.Customer
	result := session.First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}



	return &customer, nil
}

func UpdateCustomer(newCustomerData *schemas.Customer, customerID string) (*models.Customer, error) {
	birthDate, err := utils.ConvertDateStringToISODateFormat(newCustomerData.BirthDate)

	if err != nil {
		return nil, err
	}

	parsedBirthDate, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return nil, err
	}

	newCustomer := models.Customer{
		Name: newCustomerData.Name,
		CPF: newCustomerData.CPF,
		BirthDate: parsedBirthDate,
	}

	session := database.GetSession()
	result := session.Model(&models.Customer{}).Where("id = ?", customerID).
		Updates(newCustomer)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &newCustomer, nil
}

func DeleteCustomer(customerID string) error {
	session := database.GetSession()
	result := session.Where("id = ?", customerID).Delete(&models.Customer{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
