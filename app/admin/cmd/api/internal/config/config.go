package config

import (
	"github.com/zeromicro/go-zero/rest"
	"slash-admin/pkg/slconfig"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	//redis config
	Redis slconfig.RedisConf

	// database config
	Database slconfig.DatabaseConf
}
