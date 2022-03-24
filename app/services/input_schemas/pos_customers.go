package input_schemas

import "go-pos-stores/app/models"

type (
	Data struct {
		Data     models.PosCostomers `json:"data"`
		UserInfo UserInfo            `json:"user_info"`
	}

	UserInfo struct {
		User_id string `json:"user_id"`
	}
)
