package main

import (
	"foodV5/common/config"
	"foodV5/common/pkg/logs"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{}
	app.Name = "litworke接口开发"
	app.Usage = "litworke接口开发"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "conf",
			Aliases:  []string{"c"},
			Usage:    "配置文件(.yaml)",
			Required: true,
		},
	}
	app.Action = func(ctx *cli.Context) error {
		err := config.Load(ctx.String("conf"))
		if err != nil {
			return err
		}
		logs.NewLog()

		server, err := wireApp()
		if err != nil {
			return err
		}
		server.Run()
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Printf("%v", err)
	}

}
