package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/quexer/red"
	"github.com/quexer/utee"
	"github.com/spf13/viper"
	"sync"
)

var (
	once sync.Once
	Ds   *DataSource
)

type DataSource struct {
	Mysql *gorm.DB
	Rp    *redis.Pool
}

const (
	RedisServer   = "redis.server"
	RedisPassword = "redis.password"
	RedisDbno     = "redis.dbno"
	RedisSize     = "redis.size"
	MySqlHost     = "mysql.host"
	MySqlPort     = "mysql.port"
	MySqlDbName   = "mysql.db_name"
	MySqlUser     = "mysql.user"
	MySqlPassword = "mysql.password"
)

func init() {
	viper.SetDefault(RedisServer, "127.0.0.1:6379")
	viper.SetDefault(RedisPassword, "")
	viper.SetDefault(RedisDbno, 2)
	viper.SetDefault(RedisSize, 12)
	viper.SetDefault(MySqlHost, "localhost")
	viper.SetDefault(MySqlPort, 3306)
	viper.SetDefault(MySqlDbName, "tower")
	viper.SetDefault(MySqlUser, "root")
	viper.SetDefault(MySqlPassword, "123456")
}

func mysqlConStr() string {
	user := viper.GetString(MySqlUser)
	password := viper.GetString(MySqlPassword)
	host := viper.GetString(MySqlHost)
	port := viper.GetInt(MySqlPort)
	dbname := viper.GetString(MySqlDbName)
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
}

func InitCfg() {
	once.Do(func() {
		Ds = &DataSource{}
		db, err := gorm.Open("mysql", mysqlConStr())
		utee.Chk(err)
		db.LogMode(true)
		testMysql(db)
		Ds.Mysql = db

		rp := red.CreatePool(viper.GetInt(RedisSize), viper.GetString(RedisServer), viper.GetString(RedisPassword), viper.GetInt(RedisDbno))
		testRedis(rp)
		Ds.Rp = rp
	})
}

func testMysql(db *gorm.DB) {
	utee.Chk(db.Debug().Exec("SELECT sysdate()").Error)
}

func testRedis(rp *redis.Pool) {
	con := rp.Get()
	defer con.Close()
	_, err := redis.String(con.Do("ping"))
	utee.Chk(err)
}
