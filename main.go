package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "set-instances",
			Aliases: []string{"si"},
			Usage:   "Sets the number of instances",
			Action:  setInstances,
		},
		{
			Name:    "scale-instances",
			Aliases: []string{"sc"},
			Usage:   "Scale number of instances",
			Action:  setInstances,
		},
	}
	app.Run(os.Args)

}

func setInstances(c *cli.Context) {

	args := c.Args()

	boom := New(args.First())
	size, err := strconv.Atoi(args.Get(2))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = boom.SetInstances(args.Get(1), size)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	boom.Print()
}
