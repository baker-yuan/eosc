package main

import (
	"github.com/eolinker/eosc"
)

//DriverConfig http路由驱动配置
type DriverConfig struct {
	Listen  int                      `json:"listen" yaml:"listen" title:"port" description:"使用端口" default:"80" label:"端口号" maximum:"65535"`
	Method  []string                 `json:"method" yaml:"method" enum:"GET,POST,PUT,DELETE,PATH,HEAD,OPTIONS" label:"请求方式"`
	Host    []string                 `json:"host" yaml:"host" label:"域名"`
	Rules   []DriverRule             `json:"rules" yaml:"rules" label:"路由规则"`
	Target  eosc.RequireId           `json:"target" yaml:"target" skill:"github.com/eolinker/apinto/service.service.IService" required:"true" label:"目标服务"`
	Disable bool                     `json:"disable" yaml:"disable" label:"禁用路由"`
	Plugins map[string]*PluginConfig `json:"plugins" yaml:"plugins" label:"插件配置"`
}

//DriverRule http路由驱动配置Rule结构体
type DriverRule struct {
	Location string            `json:"location" yaml:"location"`
	Header   map[string]string `json:"header" yaml:"header" label:"请求头部（key:value类型）"`
	Query    map[string]string `json:"query" yaml:"query" label:"query参数（key:value类型）"`
}
