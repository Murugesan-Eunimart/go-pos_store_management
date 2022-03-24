package services

import (
	"context"
	"fmt"
	"go-pos-stores/app/models"
	"go-pos-stores/app/routes/validationSchema"
	"go-pos-stores/app/services/errorcodes"
	"go-pos-stores/app/services/input_schemas"
	"log"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	uuidv4 "github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Create api for creating the pos customers.
func CreateCustomer(c *fiber.Ctx) error {
	var data input_schemas.Data

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":    false,
			"message":   "Error in Input validation",
			"error_obj": errorcodes.GetErrorCodes("INVALID_INPUT"),
		})
	}

	data.Data.Created_By = data.UserInfo.User_id
	data.Data.Customer_id = uuidv4.New().String()

	if data.Data.Customer_type == "" {
		data.Data.Customer_type = "business"
	}

	var customer *models.PosCostomers = &data.Data

	err := mgm.CollectionByName("pos_customer").Create(customer)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":    false,
			"message":   "Customer not created",
			"error_obj": errorcodes.GetErrorCodes("CUSTOMER_NOT_CREATED"),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    customer,
		"message": "customer created",
		"status":  true,
	})
}

//Update customer api
func UpdateCustomer(c *fiber.Ctx) error {
	var data input_schemas.Data

	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	data.Data.Created_By = data.UserInfo.User_id

	if data.Data.Customer_id == "" {
		//return error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "customer_id requied",
		})
	}

	filter := bson.M{
		"account_id":  data.Data.Account_id,
		"customer_id": data.Data.Customer_id,
		"is_deleted":  data.Data.Is_deleted,
	}
	var customer *models.PosCostomers = &data.Data

	res, err := mgm.CollectionByName("pos_customer").UpdateOne(context.Background(), filter, bson.M{"$set": customer})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":    false,
			"message":   "Customer not updated",
			"error_obj": errorcodes.GetErrorCodes("DATABASE_VALIDATION_ERROR"),
		})
	}

	if res.MatchedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":    false,
			"message":   "Customer does not exits",
			"error_obj": errorcodes.GetErrorCodes("NO_DATA_FOUND"),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"data":    customer,
		"message": "customer Updated",
	})
}

//Api for list and get customers. Gets single customer if mobile is provided in the query
//else send list of customers.
func ListCustomer(c *fiber.Ctx) error {

	var perPage, pageNo, next_page, previous_page int
	var customers []models.PosCostomers

	findOptions := options.Find()

	filter := bson.M{
		"account_id": c.Query("account_id"),
	}

	filter["is_deleted"] = false

	if c.Query("mobile") != "" {
		filter["mobile"] = c.Query("mobile")
	} else {
		filter["mobile"] = map[string]bool{
			"$exists": true,
		}
	}

	if c.Query("customer_type") != "" {
		filter["customer_type"] = c.Query("customer_type")
	}

	if c.Query("customer_name") != "" {
		filter["name"] = c.Query("customer_name")
	}

	if c.Query("perPage") != "" {
		perPage, _ = strconv.Atoi(c.Query("perPage"))
	} else {
		perPage = 10
	}

	if c.Query("pageNo") != "" {
		pageNo, _ = strconv.Atoi(c.Query("pageNo"))
	} else {
		pageNo = 1
	}

	findOptions.SetSkip((int64(pageNo) - 1) * int64(perPage))
	findOptions.SetLimit(int64(perPage))

	total, _ := mgm.CollectionByName("pos_customer").CountDocuments(context.Background(), filter)

	err := mgm.CollectionByName("pos_customer").SimpleFind(&customers, filter, findOptions)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Internal error",
			"data":    err.Error(),
		})
	}

	// total_pages := int(total / int64(perPage))
	total_pages := int(math.Round(float64(total) / float64(perPage)))

	if total == 0 || total_pages == pageNo {
		next_page = 0
	} else {
		next_page = pageNo + 1
	}

	if pageNo == 1 {
		previous_page = 0
	} else {
		previous_page = pageNo - 1
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":  true,
		"message": "true",
		"data":    customers,
		"pagination": map[string]interface{}{
			"per_page":      perPage,
			"total_pages":   total_pages,
			"next_page":     next_page,
			"current_page":  previous_page,
			"previous_page": pageNo,
		},
	})

}

//api for delete pos customer
func DeleteCustomer(c *fiber.Ctx) error {
	var data validationSchema.DeleteCustomerSchema

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "error in input validation",
			"data":    err.Error(),
		})
	}

	fmt.Println("body", data)

	filter := bson.M{
		"account_id": data.Data.Account_id,
		"customer_id": bson.M{
			"$in": data.Data.Customer_ids,
		},
	}

	res, err := mgm.CollectionByName("pos_customer").UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"is_deleted": true}})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "internal error",
			"data":    err.Error(),
		})
	}

	if res.MatchedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "no records found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Deleted successfully",
		"data":    data,
	})
}

//Get customer details by account id
func GetCustomerByAccountId(c *fiber.Ctx) error {
	filter := bson.M{
		"account_id": c.Query("account_id"),
		"is_deleted": false,
	}

	if c.Query("customer_type") != "" {
		filter = bson.M{
			"account_id":    c.Query("account_id"),
			"is_deleted":    c.Query("is_deleted"),
			"customer_type": c.Query("customer_type"),
		}
	}

	type details struct {
		Name          string `json:"name"`
		Mobile        string `json:"mobile"`
		Customer_id   string `json:"customer_id"`
		Customer_type string `json:"customer_type"`
		Company_name  string `json:"company_name"`
	}

	var detail []details

	err := mgm.CollectionByName("pos_customer").SimpleFind(&detail, filter)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  true,
			"message": "internal error",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "data fetched",
		"data":    detail,
	})
}
