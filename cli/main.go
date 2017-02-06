package main

import (
	"fmt"
	"os"
	"strconv"

	boomPkg "github.com/jszroberto/boom"
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
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "force, f"},
			},
		},
		{
			Name:    "scale-instances",
			Aliases: []string{"sc"},
			Usage:   "Scale number of instances",
			Action:  scaleInstances,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "force, f"},
			},
		},
	}
	app.Run(os.Args)

}

func setInstances(c *cli.Context) {

	args := c.Args()

	boom := boomPkg.New(args.First(), c.Bool("force"))
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

func scaleInstances(c *cli.Context) {

	args := c.Args()

	boom := boomPkg.New(args.First(), c.Bool("force"))
	factor, err := strconv.ParseFloat(args.Get(2), 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = boom.ScaleInstances(args.Get(1), factor)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	boom.Print()
}
