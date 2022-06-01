package consul

import (
	"fmt"
	"github.com/ccb1900/go-micro/log"
	"github.com/ccb1900/go-micro/registry"
	"github.com/hashicorp/consul/api"
	"strconv"
)

type Registry struct {
	client *api.Client
	option *Option
}
type Option struct {
	Address          string
	Debug            bool
	HealthTimeout    string
	HealthInterval   string
	HealthDeregister string
}

func (o *Option) SetDefault() {
	o.HealthTimeout = "5s"
}
func (r *Registry) Open() error {
	defaultConfig := api.DefaultConfig()
	defaultConfig.Address = r.option.Address
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		log.Default().Debug(err)
		return err
	}
	r.option.SetDefault()
	r.client = client
	return nil
}

func (r *Registry) Close() error {
	return nil
}

func (r *Registry) Register(t registry.RegisterType, svcId, svcName string, svcHost string, svcPort int, tags []string) error {
	// 健康检查地址
	check := api.AgentServiceCheck{
		Interval:                       r.option.HealthInterval,
		Timeout:                        r.option.HealthTimeout,
		Notes:                          "health check",
		DeregisterCriticalServiceAfter: r.option.HealthDeregister,
	}
	if t == registry.GRPC {
		check.GRPC = fmt.Sprintf("%v:%v/%v", svcHost, svcPort, "grpc.service")
	} else {
		check.HTTP = "http://" + svcHost + ":" + strconv.Itoa(svcPort) + "/health"
	}

	// 设置注册信息
	registration := &api.AgentServiceRegistration{
		Name:    svcName,
		ID:      svcId,
		Address: svcHost,
		Port:    svcPort,
		Tags:    tags,
		Check:   &check,
	}
	// 注册
	err := r.client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Default().Error(err)
		return err
	}
	return nil
}
func (r *Registry) Discovery(svcName string, tag string) []registry.DiscoveryServer {
	services, _, err := r.client.Health().Service(svcName, tag, true, nil)
	if err != nil {
		log.Default().Error(err)
		return nil
	}

	ds := make([]registry.DiscoveryServer, len(services))
	for i := 0; i < len(services); i++ {
		ds[i] = registry.DiscoveryServer{
			Ip:   services[i].Node.Address,
			Port: services[i].Service.Port,
		}
	}
	return ds
}
func New(option *Option) *Registry {
	//s := config.Get[string]("a")
	return &Registry{
		option: option,
	}
}
