package database

import (
	"fmt"

	"github.com/guilhermesants/AppSocialApi/src/config"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB(config *config.Database) {
	connectionString := "sqlserver://" + config.User + ":" +
		config.Password + "@" + config.Host + ":" + config.Port +
		"?database=" + config.DB

	var err error
	DB, err = gorm.Open(config.Driver, connectionString)
	if err != nil {
		fmt.Println("fail to open connetion:", err.Error())
	}

}
