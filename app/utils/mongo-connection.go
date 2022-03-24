package utils

import (
	"fmt"
	"go-pos-stores/config"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {

	err := mgm.SetDefaultConfig(nil, "pos_stores_management", options.Client().ApplyURI(config.GetEnvValues("MONGO_URI")))

	if err != nil {

		fmt.Println(err.Error())

		log.Fatal(err)
	}

	fmt.Println("connected to database...")
}
