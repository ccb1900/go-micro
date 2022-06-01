package registry

import "github.com/ccb1900/go-micro/registry/consul"

type Registry interface {
	Open() error
	Close() error
	Discovery(svcName string, tag string) string
	HttpRegister(svcId, svcName string, svcHost string, svcPort int, tags []string) error
}

type Option struct {
	Driver string
	Consul consul.Option
}

var container map[string]Registry

func Init() {
	container = make(map[string]Registry)
}

func Get(driver string) Registry {
	return container[driver]
}

func Set(driver string, registry Registry) {
	container[driver] = registry
}
