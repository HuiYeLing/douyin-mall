package dal

import (
	"github.com/North-al/douyin-mall/app/checkout/biz/dal/mysql"
	"github.com/North-al/douyin-mall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
