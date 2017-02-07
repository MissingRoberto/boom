package main

import (
	"os"
	"strconv"

	boomPkg "github.com/jszroberto/boom"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:    "set-instances",
			Aliases: []string{"si"},
			Usage:   "Sets the number of instances",
			Action:  setInstances,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "output, o"},
				cli.BoolFlag{Name: "force, f"},
				cli.BoolFlag{Name: "diff, d"},
			},
		},
		{
			Name:    "scale-instances",
			Aliases: []string{"sc"},
			Usage:   "Scale number of instances",
			Action:  scaleInstances,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "output, o"},
				cli.BoolFlag{Name: "force, f"},
				cli.BoolFlag{Name: "diff, d"},
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
		exitWithError(err)
	}
	err = boom.SetInstances(args.Get(1), size)
	if err != nil {
		exitWithError(err)
	}
	if c.Bool("output") {
		boom.Print()
	} else if c.Bool("diff") {
		diff(boom, args.First())
	} else {
		writeFile(args.First(), boom.String())
	}
}

func scaleInstances(c *cli.Context) {

	args := c.Args()

	boom := boomPkg.New(args.First(), c.Bool("force"))
	factor, err := strconv.ParseFloat(args.Get(2), 64)
	if err != nil {
		exitWithError(err)
	}
	err = boom.ScaleInstances(args.Get(1), factor)
	if err != nil {
		exitWithError(err)
	}
	if c.Bool("output") {
		boom.Print()
	} else if c.Bool("diff") {
		diff(boom, args.First())
	} else {
		writeFile(args.First(), boom.String())
	}
}
