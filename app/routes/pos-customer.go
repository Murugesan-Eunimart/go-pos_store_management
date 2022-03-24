package routes

import (
	"go-pos-stores/app/routes/validationSchema"
	"go-pos-stores/app/services"

	"github.com/gofiber/fiber/v2"
)

func PosCustomersRoutes(app fiber.Router) {

	//Create a new pos customer
	app.Post("/pos_customer/create", validationSchema.CreateCustomerValidate, services.CreateCustomer)

	//Update a pos customer
	app.Post("/pos_customer/update", validationSchema.UpdateCustomerValidate, services.UpdateCustomer)

	//Delete a pos customer
	app.Post("pos_customer/delete", validationSchema.DeleteCustomerValidate, services.DeleteCustomer)

	//Get pos_customer list by account id
	app.Get("/pos_customer/by_account_id/list", services.GetCustomerByAccountId)

	//List pos customer
	app.Get("/pos_customer/list", services.ListCustomer)

	//Get single pos customer
	app.Get("/pos_customer/get", services.ListCustomer)

	app.Post("/pos_customer/quick_add", validationSchema.UpdateCustomerValidate, services.CreateCustomer)
}
