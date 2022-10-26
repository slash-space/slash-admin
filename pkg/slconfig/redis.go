package slconfig

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisConf struct {
	Host string `json:"Host"`
	Port int    `json:"Port"`
	Pass string `json:"Pass"`
	Type string `json:"Type"`
	TLS  bool   `json:"TLS"`
}

// NewRedis returns a Redis.
func (rc RedisConf) NewRedis() (*redis.Redis, error) {
	var opts []redis.Option

	if rc.Type == redis.ClusterType {
		opts = append(opts, redis.Cluster())
	}
	if len(rc.Pass) > 0 {
		opts = append(opts, redis.WithPass(rc.Pass))
	}
	if rc.TLS {
		opts = append(opts, redis.WithTLS())
	}

	addr := fmt.Sprintf("%s:%d", rc.Host, rc.Port)

	client := redis.New(addr, opts...)

	if client.Ping() {
		return client, nil
	}

	return nil, fmt.Errorf("NewRedis: redis connect failed")
}
