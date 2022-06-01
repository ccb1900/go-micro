package nacos

import "github.com/ccb1900/go-micro/registry"

type Option struct {
}

type Registry struct {
}

func (r *Registry) Open() error {
	//TODO implement me
	panic("implement me")
}

func (r *Registry) Close() error {
	//TODO implement me
	panic("implement me")
}

func (r *Registry) Discovery(svcName string, tag string) []registry.DiscoveryServer {
	//TODO implement me
	panic("implement me")
}

func (r *Registry) Register(t registry.RegisterType, svcId, svcName string, svcHost string, svcPort int, tags []string) error {
	//TODO implement me
	panic("implement me")
}

func New() *Registry {
	return &Registry{}
}
