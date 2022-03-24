package errorcodes

import (
	"github.com/tidwall/gjson"
)

const ErrorCodes = `{
	"INVALID_ID": {
        "error_code" : "SUB_ID",
        "description": "Object with following id doesn't exist."
    },

    "INVALID_INPUT": {
        "error_code" : "POS_CUSTOMER_0001",
        "description": "Error in input validation"
    },

    "DATABASE_VALIDATION_ERROR": {
        "error_code":"POS_STORE_0001",
        "description": "Validation error(Duplicate/Required filed empty) in Database"
    } , 

    "ERROR_WHILE_DATA_UPDATE": {
        "error_code":"POS_CUSTOMER_0001",
        "description": "Customer Data not found"
    },

    "STORE_NOT_CREATED" : {
        "error_code":"POS_STORE_0001",
        "description": "Pos store not created"
    },
    "STORE_NOT_UPDATED":{
        "error_code":"POS_STORE_0001",
        "description": "Pos store not updated"
    },
    "CUSTOMER_NOT_CREATED":{
        "error_code":"POS_CUSTOMER_0001",
        "description": "pos customer not created"
    },
    "NO_DATA_FOUND" : {
        "error_code":"POS_CUSTOMER_0001",
        "description": "Customer Data not found"
    },
    "NO_SOTRE_DATA_FOUND" : {
        "error_code":"POS_STORE_0001",
        "description": "Store Data not found"
    },
    "DATABASE_ERROR":{
        "error_code":"POS_DATABASE_0001",
        "description": "Invalid data"
    }
}
`

func GetErrorCodes(code string) map[string]string {
	value := gjson.Get(ErrorCodes, code)

	res := value.Map()

	result := map[string]string{
		"error_code":  res["error_code"].Str,
		"description": res["description"].Str,
	}

	return result
}
