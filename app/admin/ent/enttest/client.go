package enttest

import (
	"ariga.io/entcache"
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/migrate"
	"time"
)

func NewTestClient(t TestingT) *ent.Client {
	opts := []Option{
		WithOptions(ent.Debug()),
		WithMigrateOptions(migrate.WithForeignKeys(false)),
	}

	client := Open(t,
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/slash_admin_enhance?charset=utf8mb4&parseTime=True",
		opts...,
	)

	return client
}

func NewTestClientWithCache() *ent.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:16379",
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/slash_admin_enhance?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	drv := entsql.OpenDB("mysql", db)

	cacheDrv := entcache.NewDriver(
		drv,
		entcache.TTL(time.Minute),
		entcache.Levels(
			entcache.NewRedis(rdb),
		),
	)

	client := ent.NewClient(ent.Driver(cacheDrv), ent.Debug())

	err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(false))

	if err != nil {
		panic(err)
	}

	return client
}
