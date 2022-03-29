package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Retorna o cli pronto para ser executado
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI"
	app.Usage = "Busca IP's e nome de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "amazon.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IP's de endereços na internet",
			Flags:  flags,
			Action: findIPS,
		},
		{
			Name:   "server",
			Usage:  "Busca o nome do servidor",
			Flags:  flags,
			Action: findServerName,
		},
	}

	return app
}

func findIPS(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServerName(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
