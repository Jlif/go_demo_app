package config

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"jiangchen.tech/demo_app/utils/logger"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	driver := viper.GetString("db.driver")
	host := viper.GetString("db.host")
	user := viper.GetString("db.user")
	port := viper.GetString("db.port")
	pass := viper.GetString("db.pass")
	dbname := viper.GetString("db.dbname")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port)
	a, err := gormadapter.NewAdapter(driver, args, dbname)

	logger.PanicError(err, "new adapter", true)
	_, err = casbin.NewEnforcer("casbin.conf", a)
	logger.PanicError(err, "new adapter", true)

}

func GetEnforcer() *casbin.Enforcer {
	return Enforcer
}
