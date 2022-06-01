package consul

import (
	"github.com/hashicorp/consul/api"
	"go-micro/log"
	"math/rand"
	"strconv"
	"time"
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

func (r *Registry) HttpRegister(svcId, svcName string, svcHost string, svcPort int, tags []string) error {
	// 健康检查地址
	check := api.AgentServiceCheck{
		HTTP:                           "http://" + svcHost + ":" + strconv.Itoa(svcPort) + "/health",
		Interval:                       r.option.HealthInterval,
		Timeout:                        r.option.HealthTimeout,
		Notes:                          "http health check",
		DeregisterCriticalServiceAfter: r.option.HealthDeregister,
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
	}
	return nil
}
func (r *Registry) Discovery(svcName string, tag string) string {
	services, _, err := r.client.Health().Service(svcName, tag, true, nil)
	if err != nil {
		log.Default().Error(err)
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	return services[rand.Intn(len(services))].Service.Address
}
func New(option *Option) *Registry {
	//s := config.Get[string]("a")
	return &Registry{
		option: option,
	}
}
