package main

import (
	"errors"
	"os"
	"strconv"

	boomPkg "github.com/jszroberto/boom"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = "0.1.0"
	app.Name = "boom"
	app.Usage = "a simple and quick tool for bosh manifest maintenance"
	app.Commands = []cli.Command{
		{
			Name:      "mask",
			Aliases:   []string{"si"},
			Usage:     "creates a mask for a given list and a given key",
			ArgsUsage: "<MANIFEST> <LIST> <KEY>\n\nExample:\n $ boom mask manifest.yml jobs instances",
			Action:    mask,
		},
		{
			Name:      "set-instances",
			Aliases:   []string{"si"},
			Usage:     "Set the number of instances",
			ArgsUsage: "<MANIFEST> <JOB_NAME> <VALUE>\n\nExample:\n $ boom set-instances manifest.yml cell 10",
			Action:    setInstances,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "output, o", Usage: "prints result into stdout"},
				cli.BoolFlag{Name: "diff, d", Usage: "displays differences using wdiff"},
			},
		},
		{
			Name:      "scale-instances",
			Aliases:   []string{"sc"},
			Usage:     "Scale number of instances",
			ArgsUsage: "<MANIFEST> <JOB_NAME> <FACTOR>\n\nExample:\n $ boom scale-instances manifest.yml cell 1.5",
			Action:    scaleInstances,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "output, o", Usage: "prints result into stdout"},
				cli.BoolFlag{Name: "diff, d", Usage: "displays differences using wdiff"},
				cli.BoolFlag{Name: "force, f", Usage: "forces to change the current value at least an unit"},
			},
		},
	}
	app.Run(os.Args)

}

func mask(c *cli.Context) {

	if c.NArg() < 2 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	args := c.Args()
	boom := boomPkg.New(args.First(), false)
	err := boom.Mask(args.Get(1), args.Get(2))
	if err != nil {
		exitWithError(err)
	}
	boom.Print()

}

func setInstances(c *cli.Context) {

	if c.NArg() != 3 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	args := c.Args()

	boom := boomPkg.New(args.First(), false)
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

	if c.NArg() != 3 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

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
