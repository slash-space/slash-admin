package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type MenuMeta struct {
	Title              string `json:"title"`              // menu name | 菜单显示标题
	Icon               string `json:"icon"`               // menu icon | 菜单图标
	HideMenu           bool   `json:"hideMenu"`           // hide menu | 是否隐藏菜单
	HideBreadcrumb     bool   `json:"hideBreadcrumb"`     // hide the breadcrumb | 隐藏面包屑
	CurrentActiveMenu  string `json:"currentActiveMenu"`  // set the active menu | 激活菜单
	IgnoreKeepAlive    bool   `json:"ignoreKeepAlive"`    // do not keep alive the tab | 取消页面缓存
	HideTab            bool   `json:"hideTab"`            // hide the tab header | 隐藏页头
	FrameSrc           string `json:"frameSrc"`           // show iframe | 内嵌 iframe
	CarryParam         bool   `json:"carryParam"`         // the route carries parameters or not | 携带参数
	HideChildrenInMenu bool   `json:"hideChildrenInMenu"` // hide children menu or not | 隐藏所有子菜单
	Affix              bool   `json:"affix"`              // affix tab | Tab 固定
	DynamicLevel       uint32 `json:"dynamicLevel"`       // the maximum number of pages the router can open | 能打开的子TAB数
	RealPath           string `json:"realPath"`           // the real path of the route without dynamic part | 菜单路由不包含参数部分
}

func (m MenuMeta) Value() (driver.Value, error) {
	val, err := json.Marshal(m)
	if err != nil {
		fmt.Println("MenuMeta.Value() error:", err)
		return nil, err
	}
	return string(val), nil
}

func (m *MenuMeta) Scan(v any) error {
	s2 := asString(v)
	err := json.Unmarshal([]byte(s2), m)
	if err != nil {
		fmt.Println("MenuMeta.Scan() error:", err)
		return err
	}
	return nil
}
