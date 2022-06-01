package cmd

import (
	"github.com/ccb1900/go-micro/registry"
	"github.com/spf13/cobra"
)

func Server() *cobra.Command {
	c := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				return
			}
		},
	}
	c.AddCommand(&cobra.Command{
		Use: "run",
		Run: func(cmd *cobra.Command, args []string) {
			// 日志初始化
			registry.Get("xxxx")

		},
	})
	return c
}
