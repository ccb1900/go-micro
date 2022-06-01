package app

import (
	"go-micro/log"
	"go-micro/registry"
)

type App interface {
	RegisterCustomRegistry(name string, registry registry.Registry)
	RegisterCustomLog(name string, logger log.Log)
}

type Application struct {
}

func (a *Application) RegisterCustomRegistry(name string, registry registry.Registry) {

}

func (a *Application) RegisterCustomLog(name string, logger log.Log) {
	//TODO implement me
	panic("implement me")
}
