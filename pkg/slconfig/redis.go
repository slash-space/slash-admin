package slconfig

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisConf struct {
	Host string
	Port int
	Pass string
	Type string
	TLS  bool
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

	client := redis.New(rc.Host, opts...)

	if client.Ping() {
		return client, nil
	}

	return nil, fmt.Errorf("NewRedis: redis connect failed")
}
