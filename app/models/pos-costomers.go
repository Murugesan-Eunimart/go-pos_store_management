package models

import "github.com/kamva/mgm/v3"

type (
	Data struct {
		Data     PosCostomers `json:"data"`
		UserInfo UserInfo     `json:"user_info"`
	}

	PosCostomers struct {
		mgm.DefaultModel
		Name          string  `json:"name"`
		Company_name  string  `json:"company_name"`
		Email         string  `json:"email"`
		Mobile        string  `json:"mobile"`
		Address       Address `json:"address"`
		Orders_Count  string  `json:"orders_count"`
		Account_id    string  `json:"account_id"`
		Created_By    string  `json:"created_by"`
		Customer_type string  `json:"customer_type" default:"business"` //enum ["business", 'customer'], default:'business'
		Customer_id   string  `json:"customer_id" bson:"customer_id"`
		Pan           string  `json:"pan"`
		Gstin         string  `json:"gstin"`
		Notes         string  `json:"notes"`
		Is_deleted    bool    `json:"is_deleted"` // default:"false"
	}

	Address struct {
		Zip     string `json:"zip"`
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
		Address string `json:"address"`
	}

	UserInfo struct {
		User_id string `json:"user_id"`
	}
)
