package slconfig

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"log"
	"slash-admin/app/admin/ent"
	entcasbin "slash-admin/app/admin/ent/casbin"
)

// CasbinConf casbin config

type CasbinConf struct {
	ModelText string `json:"ModelText"`
}

func (config CasbinConf) NewCasbin(client *ent.Client) (*casbin.SyncedEnforcer, error) {
	var syncedEnforcer *casbin.SyncedEnforcer
	a, err := entcasbin.NewAdapterWithClient(client)

	if err != nil {
		log.Fatal("InitCasbin: init adapter fail!", err)
		return nil, err
	}

	m, err := model.NewModelFromString(config.ModelText)
	if err != nil {
		log.Fatal("InitCasbin: import model fail!", err)
		return nil, err
	}
	syncedEnforcer, err = casbin.NewSyncedEnforcer(m, a)
	if err != nil {
		log.Fatal("InitCasbin: NewSyncedEnforcer fail!", err)
		return nil, err
	}
	err = syncedEnforcer.LoadPolicy()
	if err != nil {
		log.Fatal("InitCasbin: LoadPolicy fail!", err)
		return nil, err
	}
	return syncedEnforcer, nil
}
