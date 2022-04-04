package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	unitysceneanalysis "github.com/recolude/unity-scene-analysis"
	"github.com/urfave/cli/v2"
)

const divider string = "================================================================================"

func main() {
	app := &cli.App{
		Authors: []*cli.Author{
			{
				Name:  "Eli C. Davis",
				Email: "eli@recolude.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "examine",
				Aliases: []string{"e"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Usage:    "Unity scene file to examine",
						Required: true,
						Aliases:  []string{"f"},
					},
					&cli.IntFlag{
						Name:     "limit",
						Usage:    "Max number of entires that get listed out",
						Required: false,
						Value:    -1,
					},
				},
				Usage: "lists the contents of a unity file",
				Action: func(c *cli.Context) error {
					f, err := os.Open(c.String("file"))
					if err != nil {
						return err
					}

					scene, err := unitysceneanalysis.ReadScene(f)
					if err != nil {
						return err
					}

					fmt.Fprintf(c.App.Writer, "%s\n| %d Categories Found\n%s\n", divider, len(scene.ObjectSummaries), divider)

					for _, summary := range scene.ObjectSummaries {
						fmt.Fprintf(
							c.App.Writer,
							"%d %s %s (%s)\n",
							summary.Count,
							summary.Name,
							sizeDisplay(summary.Memory),
							sizePercentDisplay(summary.Memory, scene.TotalMemory),
						)
					}

					limit := c.Int("limit")
					entriesFoundLimit := ""
					if limit > 0 {
						entriesFoundLimit = fmt.Sprintf(" (%d shown)", limit)
					}

					fmt.Fprintf(c.App.Writer, "%s\n| %d Entries Found%s\n%s\n", divider, len(scene.Sections), entriesFoundLimit, divider)

					for i, section := range scene.Sections {
						if limit > 0 && i >= limit {
							break
						}
						fmt.Fprintf(
							c.App.Writer,
							"%s: %d lines [%d-%d] %s (%s)\n",
							section.Name,
							section.Lines,
							section.StartingLine,
							section.EndingLine(),
							sizeDisplay(section.Size),
							sizePercentDisplay(section.Size, scene.TotalMemory),
						)
						if section.Info != "" {
							fmt.Fprintf(
								c.App.Writer,
								"\t%s\n\n",
								section.Info,
							)
						}
					}

					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func sizePercentDisplay(x uint64, y uint64) string {
	percent := (float64(x) / float64(y)) * 100.0
	return fmt.Sprintf("%.2f%%", percent)
}

func sizeDisplay(size uint64) string {
	if size < 1024 {
		return fmt.Sprintf("%d b", size)
	}

	if size < 1024*1024 {
		return fmt.Sprintf("%d kb", size/1024.0)
	}

	return fmt.Sprintf("%d mb", size/(1024.0*1024.0))
}
