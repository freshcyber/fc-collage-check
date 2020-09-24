package config

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/freshcyber/fc-collage-check/bean"
)

// MariaDB 数据库相关配置
var MariaDB bean.DBConfig

// Server Server Config
var Server bean.ServerConfig

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	Server.Exec = viper.GetBool("server.exec")
	Server.Host = viper.GetString("server.host")
	Server.Decimal = viper.GetInt64("server.decimal")
	Server.Limit = viper.GetInt("server.limit")
	Server.Spec = viper.GetString("server.spec")
	Server.Urlgo = viper.GetString("server.urlgo")
	Server.APIAppendKey = viper.GetString("server.api.appendkey")
	Server.APIMd5Key = viper.GetString("server.api.md5key")

	MariaDB.Dialect = viper.GetString("database.dialect")
	MariaDB.Database = viper.GetString("database.database")
	MariaDB.User = viper.GetString("database.user")
	MariaDB.Password = viper.GetString("database.password")
	MariaDB.Host = viper.GetString("database.host")
	MariaDB.Port = viper.GetInt("database.port")
	MariaDB.Charset = viper.GetString("database.charset")
	MariaDB.MaxIdleConns = viper.GetInt("database.maxIdleConns")
	MariaDB.MaxOpenConns = viper.GetInt("database.maxOpenConns")
	MariaDB.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		MariaDB.User, MariaDB.Password, MariaDB.Host, MariaDB.Port, MariaDB.Database, MariaDB.Charset)
}
