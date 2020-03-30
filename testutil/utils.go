package testutil

//UT 相关的变量初始化
import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
	"testing"
	"tgin/conf"
	"tgin/utils"
)

var testOnce sync.Once
var TestCfg conf.BITConfig
var TestMysql *xorm.Engine
var TestRedis *redis.Client

func init() {
	testing.Init()
	testOnce.Do(func() {
		TestCfg = conf.DefaultConfig

        //初始化一下mysql
		utils.InitMySQL(TestCfg.DBMysql,true)
		TestMysql = utils.GetMysqlClient()

		//初始化一下redis
		utils.InitRedisClients(TestCfg.DBRedis)
		TestRedis = utils.GetRedisClient()

	})
}
