package mysql

import (
	"fmt"

	"github.com/North-al/douyin-mall/app/auth/conf"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	klog.Info("DB: ", DB)

	if err != nil {
		panic(fmt.Errorf("Init mysql failed: %v", err))
	}
}
