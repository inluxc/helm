package main

import (
	"github.com/codegangsta/cli"
	"github.com/kubernetes/deployment-manager/pkg/dm"
	"github.com/kubernetes/deployment-manager/pkg/format"
	"github.com/kubernetes/deployment-manager/pkg/kubectl"
)

func init() {
	addCommands(doctorCmd())
}

func doctorCmd() cli.Command {
	return cli.Command{
		Name:      "doctor",
		Usage:     "Run a series of checks for necessary prerequisites.",
		ArgsUsage: "",
		Action:    func(c *cli.Context) { run(c, doctor) },
	}
}

func doctor(c *cli.Context) error {
	var runner kubectl.Runner
	runner = &kubectl.RealRunner{}
	if dm.IsInstalled(runner) {
		format.Success("You have everything you need. Go forth my friend!")
	} else {
		format.Warning("Looks like you don't have DM installed.\nRun: `helm install`")
	}

	return nil
}