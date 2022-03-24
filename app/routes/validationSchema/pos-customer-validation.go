package validationSchema

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

//Create customer schema
type (
	CreateCustomerSchema struct {
		Data     CreatePosCostomers `json:"data" validate:"required"`
		UserInfo CreateUserInfo     `json:"user_info" validate:"required"`
	}

	CreatePosCostomers struct {
		Name          string        `json:"name" validate:"required"`
		Company_name  string        `json:"company_name" validate:"required"`
		Email         string        `json:"email" validate:"required"`
		Mobile        string        `json:"mobile" validate:"required"`
		Address       CreateAddress `json:"address" validate:"required"`
		Orders_Count  string        `json:"orders_count"`
		Account_id    string        `json:"account_id" validate:"required"`
		CreatedBy     string        `json:"created_by"`
		Customer_type string        `json:"customer_type" validate:"oneof='business' 'consumer'"` //enum ["business", 'customer'], default:'business'
		Customer_id   string        `json:"customer_id" bson:"customer_id"`
		Pan           string        `json:"pan" validate:"required"`
		Gstin         string        `json:"gstin" validate:"required"`
		Notes         string        `json:"notes" validate:"required"`
		Is_deleted    bool          `json:"is_deleted"` // default:"false"
	}

	CreateAddress struct {
		Zip     string `json:"zip" validate:"required"`
		City    string `json:"city" validate:"required"`
		State   string `json:"state" validate:"required"`
		Country string `json:"country" validate:"required"`
		Address string `json:"address" validate:"required"`
	}

	CreateUserInfo struct {
		User_id string `json:"user_id" validate:"required"`
	}
)

//Update customer schema
type (
	UpdateCustomerSchema struct {
		Data     CreatePosCostomers `json:"data" validate:"required"`
		UserInfo CreateUserInfo     `json:"user_info" validate:"required"`
	}

	UpdatePosCostomers struct {
		Name          string        `json:"name" validate:"required"`
		Company_name  string        `json:"company_name" validate:"required"`
		Email         string        `json:"email"`
		Mobile        string        `json:"mobile"`
		Address       CreateAddress `json:"address" validate:"required"`
		Orders_Count  string        `json:"orders_count"`
		Account_id    string        `json:"account_id" validate:"required"`
		CreatedBy     string        `json:"created_by"`
		Customer_type string        `json:"customer_type" default:"business" validate:"required"` //enum ["business", 'customer'], default:'business'
		Customer_id   string        `json:"customer_id" bson:"customer_id" validate:"required"`
		Pan           string        `json:"pan" validate:"required"`
		Gstin         string        `json:"gstin" validate:"required"`
		Notes         string        `json:"notes" validate:"required"`
		Is_deleted    bool          `json:"is_deleted"` // default:"false"
	}

	UpdateAddress struct {
		Zip     string `json:"zip"`
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
		Address string `json:"address"`
	}

	UpdateUserInfo struct {
		User_id string `json:"user_id" validate:"required"`
	}
)

//Delete customer schema
type (
	DeleteCustomerSchema struct {
		Data DeleteCustomerData `json:"data" validation:"required"`
	}

	DeleteCustomerData struct {
		Account_id   string   `json:"account_id" validation:"required"`
		Customer_ids []string `json:"customer_ids" validation:"required"`
	}
)

var Validate = validator.New()

type validation_errors struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func CreateCustomerValidate(c *fiber.Ctx) error {
	var data CreateCustomerSchema
	var errors []validation_errors

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  false,
			"message": "Error in input validation",
			"data":    err.Error(),
		})
	}

	if err := Validate.Struct(&data); err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			var val_error validation_errors
			val_error.Field = err.Field()
			val_error.Tag = err.Tag()
			val_error.Value = err.Param()
			errors = append(errors, val_error)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Error in input validation",
			"error":   errors,
		})
	}
	return c.Next()
}

func UpdateCustomerValidate(c *fiber.Ctx) error {

	var data UpdateCustomerSchema
	var errors []validation_errors
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  false,
			"message": "Error in input validation",
			"data":    err.Error(),
		})
	}

	if err := Validate.Struct(&data); err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			var val_error validation_errors
			val_error.Field = err.Field()
			val_error.Tag = err.Tag()
			val_error.Value = err.Param()
			errors = append(errors, val_error)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Error in input validation",
			"error":   errors,
		})
	}
	return c.Next()
}

func DeleteCustomerValidate(c *fiber.Ctx) error {
	var data DeleteCustomerSchema
	var errors []validation_errors
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  false,
			"message": "Error in input validation",
			"data":    err.Error(),
		})
	}

	if err := Validate.Struct(&data); err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			var val_error validation_errors
			val_error.Field = err.Field()
			val_error.Tag = err.Tag()
			val_error.Value = err.Param()
			errors = append(errors, val_error)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Error in input validation",
			"error":   errors,
		})
	}
	return c.Next()
}
