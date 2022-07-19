package svc

import (
	"database/sql"
	"github.com/xulei131401/gox/gormx"
	"github.com/xulei131401/holy-go/service/api/admin/internal/config"
	"github.com/xulei131401/holy-go/service/api/admin/internal/model"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
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
	GormDb       *gorm.DB
	ProcessCache *collection.Cache // 进程内缓存

	// model
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	db, _ := conn.RawDB()

	gormDb := gormx.Open(c.Mysql.DataSource)

	redisConn := c.CacheRedis[0].NewRedis()

	processCache, _ := collection.NewCache(time.Second*10, collection.WithLimit(20))

	return &ServiceContext{
		Config:       c,
		SqlDbConn:    db,
		GormDb:       gormDb,
		MysqlConn:    conn,
		RedisConn:    redisConn,
		ProcessCache: processCache,

		UserModel: model.NewUserModel(conn),
	}
}
