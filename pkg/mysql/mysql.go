package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

/* Connect Connect to mysql */
func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", viper.GetString("mysql.user"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetString("mysql.port"), viper.GetString("mysql.database"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
