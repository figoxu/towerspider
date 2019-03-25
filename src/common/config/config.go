package config

import (
	"figoxu/towerspider/common/db/model"
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
	RedisServer   = "redis_server"
	RedisPassword = "redis_password"
	RedisDbno     = "redis_dbno"
	RedisSize     = "redis_size"
	MySqlHost     = "mysql_host"
	MySqlPort     = "mysql_port"
	MySqlDbName   = "mysql_db_name"
	MySqlUser     = "mysql_user"
	MySqlPassword = "mysql_password"
	TowerUser     = "tower_user"
	TowerPassword = "tower_password"
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
	viper.SetDefault(TowerUser, "test@qq.com")
	viper.SetDefault(TowerPassword, "123456")
	viper.BindEnv(TowerUser, TowerPassword)
	viper.AutomaticEnv()
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
		db.AutoMigrate(&model.ActionLog{})
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

func TowerInfo() (user, password string) {
	user = viper.GetString(TowerUser)
	password = viper.GetString(TowerPassword)
	return user, password
}
