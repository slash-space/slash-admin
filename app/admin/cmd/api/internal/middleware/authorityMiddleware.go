package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type AuthorityMiddleware struct {
	CasbinSyncedEnforcer *casbin.SyncedEnforcer
	RedisClient          *redis.Redis
}

func NewAuthorityMiddleware(enforcer *casbin.SyncedEnforcer, redisClient *redis.Redis) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		CasbinSyncedEnforcer: enforcer,
		RedisClient:          redisClient,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if you need
		next(w, r)
	}
}
