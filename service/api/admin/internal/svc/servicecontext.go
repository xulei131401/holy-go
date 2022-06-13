package svc

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"holy-go/service/api/admin/internal/config"
	"holy-go/service/api/admin/internal/model"
	"time"
)

type ServiceContext struct {
	Config config.Config

	// 中间件
	Auth rest.Middleware

	// 外部依赖
	RedisConn    *redis.Redis
	MysqlConn    sqlx.SqlConn
	SqlDbConn    *sql.DB
	ProcessCache *collection.Cache // 进程内缓存

	// model
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	db, _ := conn.RawDB()

	redisConn := c.CacheRedis[0].NewRedis()

	processCache, _ := collection.NewCache(time.Second*10, collection.WithLimit(20))

	return &ServiceContext{
		Config:       c,
		SqlDbConn:    db,
		MysqlConn:    conn,
		RedisConn:    redisConn,
		ProcessCache: processCache,

		UserModel: model.NewUserModel(conn),
	}
}
