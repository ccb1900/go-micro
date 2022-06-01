package app

import (
	"fmt"
	"github.com/ccb1900/go-micro/log"
	"github.com/ccb1900/go-micro/registry"
	"github.com/spf13/cobra"
)

type App interface {
	RegisterCustomRegistry(name string, registry registry.Registry)
	RegisterCustomLog(name string, logger log.Log)
}

type Application struct {
	rootCmd *cobra.Command
}

func (a *Application) RegisterCustomRegistry(name string, registry registry.Registry) {

}

func (a *Application) RegisterCustomLog(name string, logger log.Log) {
	//TODO implement me
	panic("implement me")
}
func (a *Application) RegisterCustomCmd(cmd *cobra.Command) {
	a.rootCmd.AddCommand(cmd)
}
func (a *Application) Run() {
	fmt.Println("app start...")
	rootCommand := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				return
			}
		},
	}
	rootCommand.AddCommand(&cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				return
			}
		},
	})
	a.rootCmd = rootCommand

	err := rootCommand.Execute()
	if err != nil {
		return
	}
}
func New() *Application {
	cobra.OnInitialize(func() {

	})
	return &Application{}
}

func Run() {
	a := New()
	a.Run()
}
